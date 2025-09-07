package integration

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// setupIntegrationApp creates a test app with real database for integration tests
// This will fail until we implement the actual app and database setup
func setupIntegrationApp(t *testing.T) *fiber.App {
	// TODO: This will be implemented in the main implementation phase
	// For now, this will panic to ensure tests fail during RED phase
	panic("setupIntegrationApp not implemented - this ensures tests fail in RED phase")
}

// cleanupIntegrationApp cleans up test data and closes connections
func cleanupIntegrationApp(t *testing.T, app *fiber.App) {
	// TODO: Cleanup database connections and test data
	// For now, this is a no-op since setupIntegrationApp panics
}

// createTestUser creates a user for integration testing
func createTestUser(t *testing.T, app *fiber.App, userData map[string]interface{}) map[string]interface{} {
	jsonPayload, err := json.Marshal(userData)
	if err != nil {
		t.Fatalf("Failed to marshal user data: %v", err)
	}

	req := httptest.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode user creation response: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("Failed to create test user, status: %d, response: %v", resp.StatusCode, response)
	}

	return response
}

// getAuthToken gets a JWT token for authenticated requests
func getAuthToken(t *testing.T, app *fiber.App, email, password string) string {
	loginData := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	jsonPayload, err := json.Marshal(loginData)
	if err != nil {
		t.Fatalf("Failed to marshal login data: %v", err)
	}

	req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode login response: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Failed to login, status: %d, response: %v", resp.StatusCode, response)
	}

	token, ok := response["token"].(string)
	if !ok {
		t.Fatalf("No token in login response")
	}

	return token
}
