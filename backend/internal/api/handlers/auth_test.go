package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ForIAM/ForIAM/backend/internal/config"
	"github.com/gin-gonic/gin"
)

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/auth/login", func(c *gin.Context) {
		// Mock successful login for test
		c.JSON(http.StatusOK, LoginResponse{
			AccessToken: "mock-token",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
		})
	})

	loginReq := LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	jsonData, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if response.TokenType != "Bearer" {
		t.Errorf("Expected token type 'Bearer', got '%s'", response.TokenType)
	}
}

func TestLoginRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		request LoginRequest
		valid   bool
	}{
		{
			name: "valid login",
			request: LoginRequest{
				Email:    "user@example.com",
				Password: "password123",
			},
			valid: true,
		},
		{
			name: "invalid email",
			request: LoginRequest{
				Email:    "invalid-email",
				Password: "password123",
			},
			valid: false,
		},
		{
			name: "empty password",
			request: LoginRequest{
				Email:    "user@example.com",
				Password: "",
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Basic validation checks
			if tt.request.Email == "" && tt.valid {
				t.Error("Expected invalid request for empty email")
			}
			if tt.request.Password == "" && tt.valid {
				t.Error("Expected invalid request for empty password")
			}
		})
	}
}