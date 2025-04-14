# Identity & Access Management Platform - API Specification

## Overview
This document defines the public API surface for the IAM platform. It adheres to RESTful design and returns JSON responses. Authentication is handled via JWTs passed in the `Authorization: Bearer` header.

---

## Authentication

### POST /auth/login
Authenticate a user and issue tokens.

**Body:**
```json
{
  "email": "user@example.com",
  "password": "examplePassword"
}
```

**Response:**
```json
{
  "access_token": "...",
  "refresh_token": "...",
  "expires_in": 3600,
  "token_type": "Bearer"
}
```

### POST /auth/logout
Revokes current session token.

### POST /auth/token/refresh
Exchange a refresh token for a new access token.

---

## Users

### GET /users
List all users in current tenant.

### POST /users
Create a new user.

### GET /users/{id}
Get user details.

### PUT /users/{id}
Update user.

### DELETE /users/{id}
Delete user.

---

## Groups

### GET /groups
List all groups.

### POST /groups
Create a group.

### POST /groups/{id}/users/{user_id}
Add user to group.

### DELETE /groups/{id}/users/{user_id}
Remove user from group.

---

## Roles

### GET /roles
List available roles.

### POST /roles
Create new role.

### POST /roles/{id}/permissions
Assign permissions to role.

### GET /roles/{id}/users
List users with this role.

---

## Permissions

### GET /permissions
List all permissions (global or tenant-specific).

---

## Audit Logs

### GET /audit
Query audit logs (paginated, filterable).

**Query Parameters:**
- `action` (e.g. "login", "user.create")
- `user_id`
- `date_from`, `date_to`

---

## Tokens

### POST /tokens
Issue API token for a user or service.

### DELETE /tokens/{id}
Revoke a token.

---

## Health

### GET /health
Returns 200 OK if the service is healthy.

**Response:**
```json
{
  "status": "ok",
  "version": "1.0.0"
}
```

---

> All requests must include tenant context, either from the JWT or explicitly via headers if needed (e.g., `X-Tenant-ID`).
