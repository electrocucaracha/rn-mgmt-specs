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

func TestAuthRegisterContract(t *testing.T) {
	// This test MUST fail initially - no implementation exists yet
	app := setupTestApp(t)

	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
		expectedFields []string
	}{
		{
			name: "successful registration",
			payload: map[string]interface{}{
				"email":      "test@example.com",
				"password":   "testpass123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: 201,
			expectedFields: []string{"id", "email", "first_name", "last_name"},
		},
		{
			name: "missing email",
			payload: map[string]interface{}{
				"password":   "testpass123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "invalid email format",
			payload: map[string]interface{}{
				"email":      "invalid-email",
				"password":   "testpass123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "password too short",
			payload: map[string]interface{}{
				"email":      "test@example.com",
				"password":   "short",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: 400,
			expectedFields: []string{"error"},
		},
		{
			name: "duplicate email",
			payload: map[string]interface{}{
				"email":      "duplicate@example.com",
				"password":   "testpass123",
				"first_name": "Test",
				"last_name":  "User",
			},
			expectedStatus: 409,
			expectedFields: []string{"error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request
			jsonPayload, err := json.Marshal(tt.payload)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(jsonPayload))
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

			// Additional validation for successful registration
			if tt.expectedStatus == 201 {
				assert.Equal(t, tt.payload["email"], response["email"])
				assert.Equal(t, tt.payload["first_name"], response["first_name"])
				assert.Equal(t, tt.payload["last_name"], response["last_name"])
				assert.NotEmpty(t, response["id"])
				assert.NotContains(t, response, "password")
				assert.NotContains(t, response, "password_hash")
			}
		})
	}
}
