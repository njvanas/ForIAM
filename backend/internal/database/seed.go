package database

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) error {
	// Check if system tenant already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tenants WHERE name = 'system'").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing tenants: %w", err)
	}

	if count > 0 {
		return nil // Already seeded
	}

	// Create system tenant
	var tenantID string
	err = db.QueryRow(`
		INSERT INTO tenants (name) 
		VALUES ('system') 
		RETURNING id
	`).Scan(&tenantID)
	if err != nil {
		return fmt.Errorf("failed to create system tenant: %w", err)
	}

	// Hash admin password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create admin user
	var userID string
	err = db.QueryRow(`
		INSERT INTO users (tenant_id, email, password_hash) 
		VALUES ($1, 'admin@system.local', $2) 
		RETURNING id
	`, tenantID, string(hashedPassword)).Scan(&userID)
	if err != nil {
		return fmt.Errorf("failed to create admin user: %w", err)
	}

	// Create admin role
	var roleID string
	err = db.QueryRow(`
		INSERT INTO roles (tenant_id, name, description) 
		VALUES ($1, 'admin', 'System Administrator Role') 
		RETURNING id
	`, tenantID).Scan(&roleID)
	if err != nil {
		return fmt.Errorf("failed to create admin role: %w", err)
	}

	// Create permissions
	permissions := []struct {
		name        string
		description string
	}{
		{"user.read", "Read user data"},
		{"user.write", "Write user data"},
		{"user.delete", "Delete user data"},
		{"group.read", "Read group data"},
		{"group.write", "Write group data"},
		{"group.delete", "Delete group data"},
		{"role.read", "Read role data"},
		{"role.write", "Write role data"},
		{"role.delete", "Delete role data"},
		{"audit.read", "View audit logs"},
		{"system.admin", "System administration"},
	}

	var permissionIDs []string
	for _, perm := range permissions {
		var permID string
		err = db.QueryRow(`
			INSERT INTO permissions (name, description) 
			VALUES ($1, $2) 
			RETURNING id
		`, perm.name, perm.description).Scan(&permID)
		if err != nil {
			return fmt.Errorf("failed to create permission %s: %w", perm.name, err)
		}
		permissionIDs = append(permissionIDs, permID)
	}

	// Assign all permissions to admin role
	for _, permID := range permissionIDs {
		_, err = db.Exec(`
			INSERT INTO role_permissions (role_id, permission_id) 
			VALUES ($1, $2)
		`, roleID, permID)
		if err != nil {
			return fmt.Errorf("failed to assign permission to admin role: %w", err)
		}
	}

	// Assign admin role to admin user
	_, err = db.Exec(`
		INSERT INTO user_roles (user_id, role_id) 
		VALUES ($1, $2)
	`, userID, roleID)
	if err != nil {
		return fmt.Errorf("failed to assign admin role to user: %w", err)
	}

	return nil
}