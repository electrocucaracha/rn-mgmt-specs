package integration

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropertyAnalysisIntegration(t *testing.T) {
	// This integration test covers the full property creation and metric calculation workflow
	// from quickstart scenario 2: Property Creation and Basic Analysis

	app := setupIntegrationApp(t)
	defer cleanupIntegrationApp(t, app)

	// Step 1: Register and login user
	testUser := map[string]interface{}{
		"email":      "investor@example.com",
		"password":   "testpass123",
		"first_name": "John",
		"last_name":  "Investor",
	}

	user := createTestUser(t, app, testUser)
	token := getAuthToken(t, app, "investor@example.com", "testpass123")

	// Step 2: Create property with complete financial data
	propertyData := map[string]interface{}{
		"address":            "123 Main St, Anytown, ST 12345",
		"year_built":         2000,
		"land_area_sqft":     6000,
		"building_area_sqft": 1500,
		"purchase_price":     250000,
		"intended_rent":      2100,
		"operating_expenses": map[string]interface{}{
			"insurance":      1200,
			"property_taxes": 3600,
			"hoa":            0,
		},
		"financing_terms": map[string]interface{}{
			"interest_rate":        7.5,
			"loan_term":            30,
			"down_payment_percent": 20,
			"closing_costs":        5000,
		},
		"operating_assumptions": map[string]interface{}{
			"vacancy_rate":    0.05,
			"maintenance_pct": 0.10,
			"management_pct":  0.08,
			"utilities":       0,
		},
	}

	// Create property
	jsonPayload, err := json.Marshal(propertyData)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/api/v1/properties", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var property map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&property)
	require.NoError(t, err)

	propertyID := property["id"].(string)
	assert.NotEmpty(t, propertyID)
	assert.Equal(t, user["id"], property["user_id"])

	// Step 3: Verify financial metrics are calculated automatically
	req = httptest.NewRequest("GET", "/api/v1/properties/"+propertyID+"/metrics", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var metrics map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&metrics)
	require.NoError(t, err)

	// Verify expected metrics calculations (from quickstart.md expected results)
	// Monthly Mortgage Payment: ~$1,398
	monthlyPayment := metrics["monthly_mortgage_payment"].(float64)
	assert.InDelta(t, 1398, monthlyPayment, 50, "Monthly mortgage payment should be around $1,398")

	// Net Operating Income: ~$19,392
	noi := metrics["net_operating_income"].(float64)
	assert.InDelta(t, 19392, noi, 500, "NOI should be around $19,392")

	// Cap Rate: ~7.76%
	capRate := metrics["cap_rate"].(float64)
	assert.InDelta(t, 7.76, capRate, 0.5, "Cap rate should be around 7.76%")

	// Cash-on-Cash Return: ~15.2%
	cocReturn := metrics["cash_on_cash_return"].(float64)
	assert.InDelta(t, 15.2, cocReturn, 1.0, "Cash-on-Cash return should be around 15.2%")

	// Cash to Close: $55,000
	cashToClose := metrics["cash_to_close"].(float64)
	assert.Equal(t, 55000.0, cashToClose, "Cash to close should be exactly $55,000")

	// RTV: 10.08%
	rtv := metrics["rent_to_value_ratio"].(float64)
	assert.InDelta(t, 10.08, rtv, 0.1, "RTV should be around 10.08%")

	// GRM: 9.92
	grm := metrics["gross_rent_multiplier"].(float64)
	assert.InDelta(t, 9.92, grm, 0.1, "GRM should be around 9.92")

	// Step 4: Verify property details include calculated metrics
	req = httptest.NewRequest("GET", "/api/v1/properties/"+propertyID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var propertyDetail map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&propertyDetail)
	require.NoError(t, err)

	// Verify property includes metrics
	assert.Contains(t, propertyDetail, "financial_metrics")
	propertyMetrics := propertyDetail["financial_metrics"].(map[string]interface{})
	assert.Equal(t, metrics["cap_rate"], propertyMetrics["cap_rate"])
	assert.Equal(t, metrics["cash_on_cash_return"], propertyMetrics["cash_on_cash_return"])

	// Step 5: Test metric recalculation when property data changes
	updateData := map[string]interface{}{
		"intended_rent": 2300, // Increase from 2100 to 2300
	}

	jsonPayload, err = json.Marshal(updateData)
	require.NoError(t, err)

	req = httptest.NewRequest("PUT", "/api/v1/properties/"+propertyID, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Trigger metric recalculation
	req = httptest.NewRequest("POST", "/api/v1/properties/"+propertyID+"/metrics", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify updated metrics
	req = httptest.NewRequest("GET", "/api/v1/properties/"+propertyID+"/metrics", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var updatedMetrics map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&updatedMetrics)
	require.NoError(t, err)

	// With rent increased from 2100 to 2300, Cap Rate should increase to ~8.44%
	newCapRate := updatedMetrics["cap_rate"].(float64)
	assert.Greater(t, newCapRate, capRate, "Cap rate should increase with higher rent")
	assert.InDelta(t, 8.44, newCapRate, 0.5, "New cap rate should be around 8.44%")

	// Cash-on-Cash Return should increase to ~18.1%
	newCocReturn := updatedMetrics["cash_on_cash_return"].(float64)
	assert.Greater(t, newCocReturn, cocReturn, "CoC return should increase with higher rent")
	assert.InDelta(t, 18.1, newCocReturn, 1.0, "New CoC return should be around 18.1%")
}
