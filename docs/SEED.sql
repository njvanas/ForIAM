-- Seed data for IAM Platform

-- Create system tenant
INSERT INTO tenants (id, name) VALUES (
    '00000000-0000-0000-0000-000000000001', 
    'system'
);

-- Create initial admin user
INSERT INTO users (id, tenant_id, email, password_hash) VALUES (
    '00000000-0000-0000-0000-000000000100',
    '00000000-0000-0000-0000-000000000001',
    'admin@system.local',
    '$2y$12$examplehashedpassword'  -- Replace with real bcrypt hash
);

-- Create admin role
INSERT INTO roles (id, tenant_id, name, description) VALUES (
    '00000000-0000-0000-0000-000000000200',
    '00000000-0000-0000-0000-000000000001',
    'admin',
    'System Administrator Role'
);

-- Assign admin role to admin user
INSERT INTO user_roles (user_id, role_id) VALUES (
    '00000000-0000-0000-0000-000000000100',
    '00000000-0000-0000-0000-000000000200'
);

-- Common permissions
INSERT INTO permissions (id, name, description) VALUES
    (gen_random_uuid(), 'user.read', 'Read user data'),
    (gen_random_uuid(), 'user.write', 'Write user data'),
    (gen_random_uuid(), 'group.manage', 'Manage user groups'),
    (gen_random_uuid(), 'role.manage', 'Manage roles and permissions'),
    (gen_random_uuid(), 'audit.view', 'View audit logs');

-- Assign all permissions to admin role
INSERT INTO role_permissions (role_id, permission_id)
SELECT 
    '00000000-0000-0000-0000-000000000200',
    id
FROM permissions;
