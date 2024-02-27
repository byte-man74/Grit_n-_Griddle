// authentication_controller_test.go
package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount_ValidInput(t *testing.T) {
	// Set up a sample valid input payload
	router := gin.Default()
	validPayload := []byte(`{
		"phone_number": "12345678901",
		"password": "strongpassword",
		"first_name": "John",
		"last_name": "Doe"
	}`)

	// Create a request with the valid payload
	req, err := http.NewRequest("POST", "/create-account", bytes.NewBuffer(validPayload))
	assert.NoError(t, err)

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Call the endpoint handler function
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, res.Code)

	// Decode the response body for further assertions
	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

	// Perform assertions on the response data
	assert.Contains(t, response, "data")
	// Add more assertions based on your expected response format
}

func TestCreateAccount_InvalidInput(t *testing.T) {
	router := gin.Default()
	// Set up a sample invalid input payload
	invalidPayload := []byte(`{
		"phone_number": "123",  // Invalid phone number
		"password": "weak",     // Weak password
		"first_name": "John",
		"last_name": "Doe"
	}`)

	// Create a request with the invalid payload
	req, err := http.NewRequest("POST", "/create-account", bytes.NewBuffer(invalidPayload))
	assert.NoError(t, err)

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Call the endpoint handler function
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, res.Code)

	// Decode the response body for further assertions
	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

	// Perform assertions on the error response data
	assert.Contains(t, response, "error")
	// Add more assertions based on your expected error response format
}
