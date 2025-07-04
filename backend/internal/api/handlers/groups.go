package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	db *sql.DB
}

func NewGroupHandler(db *sql.DB) *GroupHandler {
	return &GroupHandler{db: db}
}

type Group struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenant_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateGroupRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *GroupHandler) GetGroups(c *gin.Context) {
	tenantID, _ := c.Get("tenant_id")

	rows, err := h.db.Query(`
		SELECT id, tenant_id, name, description, created_at 
		FROM groups 
		WHERE tenant_id = $1 
		ORDER BY created_at DESC
	`, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var group Group
		if err := rows.Scan(&group.ID, &group.TenantID, &group.Name, &group.Description, &group.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan group"})
			return
		}
		groups = append(groups, group)
	}

	c.JSON(http.StatusOK, groups)
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var req CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := c.Get("tenant_id")

	var group Group
	err := h.db.QueryRow(`
		INSERT INTO groups (tenant_id, name, description) 
		VALUES ($1, $2, $3) 
		RETURNING id, tenant_id, name, description, created_at
	`, tenantID, req.Name, req.Description).Scan(
		&group.ID, &group.TenantID, &group.Name, &group.Description, &group.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusCreated, group)
}

func (h *GroupHandler) GetGroup(c *gin.Context) {
	groupID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	var group Group
	err := h.db.QueryRow(`
		SELECT id, tenant_id, name, description, created_at 
		FROM groups 
		WHERE id = $1 AND tenant_id = $2
	`, groupID, tenantID).Scan(&group.ID, &group.TenantID, &group.Name, &group.Description, &group.CreatedAt)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	groupID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	var req CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.db.Exec(`
		UPDATE groups 
		SET name = $1, description = $2 
		WHERE id = $3 AND tenant_id = $4
	`, req.Name, req.Description, groupID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group updated successfully"})
}

func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	groupID := c.Param("id")
	tenantID, _ := c.Get("tenant_id")

	result, err := h.db.Exec(`
		DELETE FROM groups 
		WHERE id = $1 AND tenant_id = $2
	`, groupID, tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete group"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
}