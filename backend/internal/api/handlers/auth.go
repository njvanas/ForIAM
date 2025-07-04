package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ForIAM/ForIAM/backend/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db  *sql.DB
	cfg *config.Config
}

func NewAuthHandler(db *sql.DB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{db: db, cfg: cfg}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type User struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user from database
	var user User
	var passwordHash string
	err := h.db.QueryRow(`
		SELECT id, tenant_id, email, password_hash, is_active, created_at 
		FROM users 
		WHERE email = $1 AND is_active = true
	`, req.Email).Scan(&user.ID, &user.TenantID, &user.Email, &passwordHash, &user.IsActive, &user.CreatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"tenant_id": user.TenantID,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"iat":       time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Log successful login
	h.logAudit(user.TenantID, user.ID, "auth.login", "", "success", c.ClientIP(), c.GetHeader("User-Agent"))

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   86400, // 24 hours
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// In a real implementation, you might want to blacklist the token
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var user User
	err := h.db.QueryRow(`
		SELECT id, tenant_id, email, is_active, created_at 
		FROM users 
		WHERE id = $1
	`, userID).Scan(&user.ID, &user.TenantID, &user.Email, &user.IsActive, &user.CreatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) logAudit(tenantID, userID, action, resource, status, ip, userAgent string) {
	_, err := h.db.Exec(`
		INSERT INTO audit_logs (tenant_id, user_id, action, resource, status, ip_address, user_agent)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, tenantID, userID, action, resource, status, ip, userAgent)
	if err != nil {
		// Log error but don't fail the request
		println("Failed to log audit:", err.Error())
	}
}