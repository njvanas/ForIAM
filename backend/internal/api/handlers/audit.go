package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AuditHandler struct {
	db *sql.DB
}

func NewAuditHandler(db *sql.DB) *AuditHandler {
	return &AuditHandler{db: db}
}

type AuditLog struct {
	ID         string    `json:"id"`
	TenantID   string    `json:"tenant_id"`
	UserID     *string   `json:"user_id"`
	Action     string    `json:"action"`
	Resource   *string   `json:"resource"`
	ResourceID *string   `json:"resource_id"`
	IPAddress  *string   `json:"ip_address"`
	UserAgent  *string   `json:"user_agent"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type AuditResponse struct {
	Logs  []AuditLog `json:"logs"`
	Total int        `json:"total"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
}

func (h *AuditHandler) GetAuditLogs(c *gin.Context) {
	tenantID, _ := c.Get("tenant_id")

	// Parse query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	action := c.Query("action")
	userID := c.Query("user_id")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 50
	}

	offset := (page - 1) * limit

	// Build query
	query := `
		SELECT id, tenant_id, user_id, action, resource, resource_id, 
		       ip_address, user_agent, status, created_at 
		FROM audit_logs 
		WHERE tenant_id = $1
	`
	args := []interface{}{tenantID}
	argCount := 1

	if action != "" {
		argCount++
		query += " AND action = $" + strconv.Itoa(argCount)
		args = append(args, action)
	}

	if userID != "" {
		argCount++
		query += " AND user_id = $" + strconv.Itoa(argCount)
		args = append(args, userID)
	}

	query += " ORDER BY created_at DESC"

	// Count total records
	countQuery := "SELECT COUNT(*) FROM audit_logs WHERE tenant_id = $1"
	countArgs := []interface{}{tenantID}
	if action != "" {
		countQuery += " AND action = $2"
		countArgs = append(countArgs, action)
	}
	if userID != "" {
		if action != "" {
			countQuery += " AND user_id = $3"
		} else {
			countQuery += " AND user_id = $2"
		}
		countArgs = append(countArgs, userID)
	}

	var total int
	err := h.db.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count audit logs"})
		return
	}

	// Add pagination
	query += " LIMIT $" + strconv.Itoa(argCount+1) + " OFFSET $" + strconv.Itoa(argCount+2)
	args = append(args, limit, offset)

	rows, err := h.db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var logs []AuditLog
	for rows.Next() {
		var log AuditLog
		if err := rows.Scan(
			&log.ID, &log.TenantID, &log.UserID, &log.Action,
			&log.Resource, &log.ResourceID, &log.IPAddress,
			&log.UserAgent, &log.Status, &log.CreatedAt,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan audit log"})
			return
		}
		logs = append(logs, log)
	}

	response := AuditResponse{
		Logs:  logs,
		Total: total,
		Page:  page,
		Limit: limit,
	}

	c.JSON(http.StatusOK, response)
}