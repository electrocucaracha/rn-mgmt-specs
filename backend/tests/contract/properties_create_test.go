package contract

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropertiesCreateContract(t *testing.T) {
	// This test MUST fail initially - no implementation exists yet
	app := setupTestApp(t)

	// Create test user and get auth token
	testUser := map[string]interface{}{
		"email":      "property@example.com",
		"password":   "testpass123",
		"first_name": "Property",
		"last_name":  "Owner",
	}
	createTestUser(t, app, testUser)
	token := getAuthToken(t, app, "property@example.com", "testpass123")

	tests := []struct {
		name           string
		payload        map[string]interface{}
		useAuth        bool
		expectedStatus int
		expectedFields []string
	}{
		{
			name: "successful property creation",
			payload: map[string]interface{}{
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
			},
			useAuth:        true,
			expectedStatus: 201,
			expectedFields: []string{"id", "address", "purchase_price", "intended_rent", "user_id", "created_at"},
		},
		{
			name: "missing required address",
			payload: map[string]interface{}{
				"purchase_price": 250000,
				"intended_rent":  2100,
			},
			useAuth:        true,
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "missing required purchase_price",
			payload: map[string]interface{}{
				"address":       "123 Main St, Anytown, ST 12345",
				"intended_rent": 2100,
			},
			useAuth:        true,
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "negative purchase_price",
			payload: map[string]interface{}{
				"address":        "123 Main St, Anytown, ST 12345",
				"purchase_price": -100000,
				"intended_rent":  2100,
			},
			useAuth:        true,
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "unauthorized request",
			payload: map[string]interface{}{
				"address":        "123 Main St, Anytown, ST 12345",
				"purchase_price": 250000,
				"intended_rent":  2100,
			},
			useAuth:        false,
			expectedStatus: 401,
			expectedFields: []string{"error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request
			jsonPayload, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/properties", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")

			if tt.useAuth {
				req.Header.Set("Authorization", "Bearer "+token)
			}

			// Execute request
			resp, err := app.Test(req)
			require.NoError(t, err)

			// Verify status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Verify response structure
			var response map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			require.NoError(t, err)

			// Check expected fields exist
			for _, field := range tt.expectedFields {
				assert.Contains(t, response, field, "Response should contain field: %s", field)
			}

			// Additional validation for successful creation
			if tt.expectedStatus == 201 {
				assert.Equal(t, tt.payload["address"], response["address"])
				assert.Equal(t, tt.payload["purchase_price"], response["purchase_price"])
				assert.Equal(t, tt.payload["intended_rent"], response["intended_rent"])
				assert.NotEmpty(t, response["id"])
				assert.NotEmpty(t, response["user_id"])
				assert.NotEmpty(t, response["created_at"])
			}
		})
	}
}
