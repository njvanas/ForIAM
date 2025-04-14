# 🔐 ForIAM

A modern, modular, scalable Identity and Access Management (IAM) system — designed to be clean, extensible, and future-proof.

Inspired by OpenIAM and Okta, but with our own minimalistic, developer-focused twist.

---

## 📌 Overview

This platform provides:
- ✅ Secure user authentication (JWT, sessions, OAuth2)
- ✅ Role-based access control (RBAC + groups)
- ✅ Multitenant architecture from the ground up
- ✅ Full audit logging & session tracking
- ✅ REST API + Admin UI (React-based)
- ✅ Lightweight, containerized, cloud-ready design

> For a full feature breakdown, see [`ARCHITECTURE.md`](./docs/ARCHITECTURE.md)

---

## 🚀 Quick Start (Developers)

### Prerequisites:
- Docker & Docker Compose
- `make` (optional)

### Steps:
```bash
git clone https://github.com/ForIAM/ForIAM.git
cd ForIAM
make up
```

- React UI: http://localhost:3000  
- Backend API: http://localhost:8080  
- Database: Postgres (or SQLite in dev)  
- Cache: Redis  

Use `make down` to stop and clean up.

---

## 🛠️ Tech Stack

| Layer        | Tech                  |
|--------------|------------------------|
| Frontend     | React + TypeScript     |
| Backend API  | Go (Golang)            |
| DB           | PostgreSQL / SQLite    |
| Cache        | Redis                  |
| AuthN        | OAuth2, JWT, Sessions  |
| AuthZ        | Casbin (RBAC)          |
| DevOps       | Docker, Kubernetes     |

---

## 🌐 API Overview

See full spec in [`API.md`](./docs/API.md)

Examples:
```http
POST /auth/login
GET /users
POST /groups/:id/users/:user_id
GET /audit
```

All requests require Bearer JWT and tenant context.

---

## 🔒 Security Highlights

- Passwords hashed with Argon2 or Bcrypt
- Multi-factor auth (TOTP)
- RBAC with per-tenant role scoping
- Full audit logging on critical actions
- Secure cookies, SameSite, CSRF protection
- See [`SECURITY.md`](./docs/SECURITY.md)

---

## 🚚 Deployment

Use:
- Docker Compose for local dev
- Kubernetes (with Helm) for staging/prod
- GitHub Actions for CI/CD

See [`DEPLOY.md`](./docs/DEPLOY.md) for full deployment instructions.

---

## 📆 Roadmap

For planned features and milestones, see [`ROADMAP.md`](./docs/ROADMAP.md)

Coming soon:
- 🔜 SAML 2.0 / SCIM support
- 🔜 WebAuthn (passwordless)
- 🔜 Policy engine (ABAC, ReBAC)
- 🔜 Admin approval workflows
- 🔜 CIAM use-case extensions

---

## 🧑‍💻 Contributing

1. Fork the repo
2. Create a feature branch
3. Run:
```bash
make lint
make test
```
4. Submit a PR with a clear description

See [`CONTRIBUTING.md`](./docs/CONTRIBUTING.md)

---

## 📁 Documentation Index

| File                 | Description                        |
|----------------------|------------------------------------|
| `ARCHITECTURE.md`    | High-level system design           |
| `API.md`             | Full API reference                 |
| `SECURITY.md`        | Auth, threats, encryption, etc     |
| `DEPLOY.md`          | Setup & infrastructure guide       |
| `SCHEMA.sql`         | Database schema                    |
| `ROADMAP.md`         | Feature plans                      |
| `CONTRIBUTING.md`    | Dev workflow & PR process          |

---

## 🧠 License & Ownership

This project is open for internal and external contributors. Reach out to the core team if you want to get involved in long-term development.

---

> Identity is infrastructure. We're building it right, from the start.

