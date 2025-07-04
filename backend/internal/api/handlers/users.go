package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserRequest struct {
	Email    string `json:"email,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"`
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	tenantID, _ := c.Get("tenant_id")

	rows, err := h.db.Query(`
		SELECT id, tenant_id, email, is_active, created_at 
		FROM users 
		WHERE tenant_id = $1 
		ORDER BY created_at DESC
	`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.TenantID, &user.Email, &user.IsActive, &user.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := c.Get("tenant_id")

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	var user User
	err = h.db.QueryRow(`
		INSERT INTO users (tenant_id, email, password_hash) 
		VALUES ($1, $2, $3) 
		RETURNING id, tenant_id, email, is_active, created_at
	`, tenantID, req.Email, string(hashedPassword)).Scan(
		&user.ID, &user.TenantID, &user.Email, &user.IsActive, &user.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	var user User
	err := h.db.QueryRow(`
		SELECT id, tenant_id, email, is_active, created_at 
		FROM users 
		WHERE id = $1 AND tenant_id = $2
	`, userID, tenantID).Scan(&user.ID, &user.TenantID, &user.Email, &user.IsActive, &user.CreatedAt)

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

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build dynamic update query
	query := "UPDATE users SET "
	args := []interface{}{}
	argCount := 0

	if req.Email != "" {
		argCount++
		query += "email = $" + string(rune(argCount+'0')) + ", "
		args = append(args, req.Email)
	}

	if req.IsActive != nil {
		argCount++
		query += "is_active = $" + string(rune(argCount+'0')) + ", "
		args = append(args, *req.IsActive)
	}

	if argCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No fields to update"})
		return
	}

	// Remove trailing comma and add WHERE clause
	query = query[:len(query)-2] + " WHERE id = $" + string(rune(argCount+1+'0')) + " AND tenant_id = $" + string(rune(argCount+2+'0'))
	args = append(args, userID, tenantID)

	_, err := h.db.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	result, err := h.db.Exec(`
		DELETE FROM users 
		WHERE id = $1 AND tenant_id = $2
	`, userID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}