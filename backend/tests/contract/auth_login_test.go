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

func TestAuthLoginContract(t *testing.T) {
	// This test MUST fail initially - no implementation exists yet
	app := setupTestApp(t)

	// First create a test user (this will also fail until registration is implemented)
	testUser := map[string]interface{}{
		"email":      "login@example.com",
		"password":   "testpass123",
		"first_name": "Login",
		"last_name":  "User",
	}
	createTestUser(t, app, testUser)

	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
		expectedFields []string
	}{
		{
			name: "successful login",
			payload: map[string]interface{}{
				"email":    "login@example.com",
				"password": "testpass123",
			},
			expectedStatus: 200,
			expectedFields: []string{"token", "user"},
		},
		{
			name: "invalid email",
			payload: map[string]interface{}{
				"email":    "nonexistent@example.com",
				"password": "testpass123",
			},
			expectedStatus: 401,
			expectedFields: []string{"error"},
		},
		{
			name: "invalid password",
			payload: map[string]interface{}{
				"email":    "login@example.com",
				"password": "wrongpassword",
			},
			expectedStatus: 401,
			expectedFields: []string{"error"},
		},
		{
			name: "missing email",
			payload: map[string]interface{}{
				"password": "testpass123",
			},
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "missing password",
			payload: map[string]interface{}{
				"email": "login@example.com",
			},
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request
			jsonPayload, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")

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

			// Additional validation for successful login
			if tt.expectedStatus == 200 {
				assert.NotEmpty(t, response["token"])
				userObj := response["user"].(map[string]interface{})
				assert.Equal(t, tt.payload["email"], userObj["email"])
				assert.NotContains(t, userObj, "password")
				assert.NotContains(t, userObj, "password_hash")
			}
		})
	}
}
