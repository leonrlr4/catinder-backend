// internal/handler/user_test.go
package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", RegisterUser)

	// Create a valid registration request
	regInfo := RegisterInfo{
		Username: "testuser",
		Password: "testpassword",
	}

	// Marshal the registration info to JSON
	regInfoJSON, _ := json.Marshal(regInfo)
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(regInfoJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	// Record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, resp.Code)

	// Check the response body is what we expect
	expectedBody := gin.H{"message": "Registration successful"}
	expectedBodyJSON, _ := json.Marshal(expectedBody)
	assert.JSONEq(t, string(expectedBodyJSON), resp.Body.String())
}
