-- +migrate Up

-- Create all tables
-- (Include everything from SCHEMA.sql here, typically in smaller parts per migration file)

-- +migrate Down

-- Drop all tables (in reverse order to avoid FK issues)
DROP TABLE IF EXISTS api_tokens;
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS user_groups;
DROP TABLE IF EXISTS group_roles;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tenants;
