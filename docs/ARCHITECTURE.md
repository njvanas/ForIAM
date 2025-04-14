# Identity & Access Management Platform - Architecture Overview

## Purpose
This document outlines the full architecture, design principles, feature set, and system components of our IAM platform. It is a living document, evolving as we develop and grow the platform.

## Vision
Build a lightweight, secure, modular, and enterprise-grade IAM system that is easy to extend and scale. Inspired by Okta and OpenIAM, but cleaner, faster, and more developer-friendly.

## Guiding Principles
- Consistency across all environments (dev, staging, prod)
- Security-first approach
- Minimal, modular services
- Full multitenancy support
- Clean REST API and pluggable modules
- Modern tooling and practices

---

## High-Level Architecture

### Services:
- Frontend (React + TypeScript)
  - Admin & end-user portals
  - SPA consuming REST API
  - Auth, user & role management, audit log UI

- Backend (Go)
  - Auth Service: login, MFA, session handling, JWT & OAuth2
  - Identity Service: user, group, org, and role management
  - RBAC Engine: Casbin-based permission logic
  - Audit Service: logs actions, user events
  - Token Service: issues and verifies JWTs, refresh tokens, API keys

- Storage
  - PostgreSQL: primary data store
  - Redis: caching & temporary auth/session data
  - SQLite: embedded for edge/single-tenant mode

- DevOps & Tools
  - Docker/Docker Compose: for local/dev
  - Kubernetes (K8s): for staging/prod
  - GitHub Actions CI/CD
  - MailDev for email flows
  - Swagger UI for API docs
  - Adminer/pgAdmin for DB inspection

---

## Multitenancy Strategy
- Shared schema, tenant isolation by tenant_id
- Optional Row-Level Security in Postgres
- JWT includes tenant_id claim
- Per-tenant roles, groups, and config
- Super-admin tenant for platform-level management

---

## API Design
- RESTful with standard verbs
- JWT-based auth (Bearer token)
- OAuth2 support for app integrations
- API versioning via headers or URL
- All endpoints scoped by tenant

---

## Database Schema (Simplified)

Tables:
- tenants
- users (FK: tenant_id)
- groups (FK: tenant_id)
- roles (FK: tenant_id)
- permissions
- user_roles
- group_roles
- user_groups
- audit_logs
- api_tokens

Optional:
- mfa_devices
- login_attempts
- custom_attributes

---

## Authentication
- Email/password (bcrypt or argon2)
- JWT for API & session auth
- OAuth2 / OIDC for external integrations
- TOTP-based MFA (Google Authenticator compatible)
- Session storage (Redis)

---

## Authorization
- Role-Based Access Control (RBAC)
- Powered by Casbin (supporting role hierarchies)
- Matrix UI in frontend for role vs permissions
- Option to extend to ABAC or ReBAC in future

---

## Security Practices
- Secure HTTP-only, SameSite cookies
- Bcrypt/Argon2 password hashing
- Rate limiting + brute force protection
- Role & resource-based authZ
- Audit logging for sensitive actions
- CI scans + dependency vulnerability checks

---

## CI/CD & DevOps
- GitHub Actions pipeline:
  - Lint, test, build for backend + frontend
  - Docker image build + push
  - Deploy to dev/stage environments
- Future: Helm charts for Kubernetes

---

## Dev Environments
- Use docker-compose to spin up:
  - Go backend
  - React frontend
  - Postgres DB
  - Redis
- .env and Makefile for fast bootstrapping
- Optional SQLite mode for isolated dev/testing

---

## Roadmap (Q2–Q4)
- ✅ Containerized local dev
- ✅ Backend + frontend scaffolding
- 🔲 Database schema + migrations
- 🔲 Auth flow (login, token issue, refresh)
- 🔲 RBAC enforcement
- 🔲 Admin UI MVP

---

## Contribution Workflow
- PRs must include tests
- Follow Go and React style guides
- Run make lint && make test before pushing

---

## Documentation Assets
- docs/ARCHITECTURE.md (this file)
- docs/API.md – endpoint specs
- docs/SECURITY.md – threat model & hardening
- docs/DEPLOY.md – deployment & hosting guides
- docs/SCHEMA.sql – schema reference
- docs/README.md – project overview + quick start

> This architecture is evolving. Treat this doc as the single source of truth, and update it with every system change.
