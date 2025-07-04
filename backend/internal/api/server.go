package api

import (
	"database/sql"

	"github.com/ForIAM/ForIAM/backend/internal/api/handlers"
	"github.com/ForIAM/ForIAM/backend/internal/api/middleware"
	"github.com/ForIAM/ForIAM/backend/internal/config"
	"github.com/gin-gonic/gin"
)

func NewServer(db *sql.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORS())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db, cfg)
	userHandler := handlers.NewUserHandler(db)
	roleHandler := handlers.NewRoleHandler(db)
	groupHandler := handlers.NewGroupHandler(db)
	auditHandler := handlers.NewAuditHandler(db)

	// Auth routes (no middleware)
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}

	// Protected routes
	api := r.Group("/")
	api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Auth profile
		api.GET("/auth/profile", authHandler.GetProfile)

		// Users
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUser)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)

		// Roles
		api.GET("/roles", roleHandler.GetRoles)
		api.POST("/roles", roleHandler.CreateRole)
		api.GET("/roles/:id", roleHandler.GetRole)
		api.PUT("/roles/:id", roleHandler.UpdateRole)
		api.DELETE("/roles/:id", roleHandler.DeleteRole)

		// Groups
		api.GET("/groups", groupHandler.GetGroups)
		api.POST("/groups", groupHandler.CreateGroup)
		api.GET("/groups/:id", groupHandler.GetGroup)
		api.PUT("/groups/:id", groupHandler.UpdateGroup)
		api.DELETE("/groups/:id", groupHandler.DeleteGroup)

		// Audit
		api.GET("/audit", auditHandler.GetAuditLogs)
	}

	return r
}