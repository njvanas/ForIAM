package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	db *sql.DB
}

func NewRoleHandler(db *sql.DB) *RoleHandler {
	return &RoleHandler{db: db}
}

type Role struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *RoleHandler) GetRoles(c *gin.Context) {
	tenantID, _ := c.Get("tenant_id")

	rows, err := h.db.Query(`
		SELECT id, tenant_id, name, description, created_at 
		FROM roles 
		WHERE tenant_id = $1 
		ORDER BY created_at DESC
	`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.ID, &role.TenantID, &role.Name, &role.Description, &role.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan role"})
			return
		}
		roles = append(roles, role)
	}

	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := c.Get("tenant_id")

	var role Role
	err := h.db.QueryRow(`
		INSERT INTO roles (tenant_id, name, description) 
		VALUES ($1, $2, $3) 
		RETURNING id, tenant_id, name, description, created_at
	`, tenantID, req.Name, req.Description).Scan(
		&role.ID, &role.TenantID, &role.Name, &role.Description, &role.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	roleID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	var role Role
	err := h.db.QueryRow(`
		SELECT id, tenant_id, name, description, created_at 
		FROM roles 
		WHERE id = $1 AND tenant_id = $2
	`, roleID, tenantID).Scan(&role.ID, &role.TenantID, &role.Name, &role.Description, &role.CreatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, role)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	roleID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.db.Exec(`
		UPDATE roles 
		SET name = $1, description = $2 
		WHERE id = $3 AND tenant_id = $4
	`, req.Name, req.Description, roleID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully"})
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	result, err := h.db.Exec(`
		DELETE FROM roles 
		WHERE id = $1 AND tenant_id = $2
	`, roleID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete role"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}