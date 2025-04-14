# IAM Platform - Security Overview

This document outlines the security model, protections, and best practices embedded in the platform's design. It will evolve alongside the product.

---

## 🧱 Security Foundations

| Principle         | Description                                                   |
|------------------|---------------------------------------------------------------|
| Least Privilege  | Each user/service has the minimum required permissions        |
| Defense in Depth | Multiple layers of controls: authN, authZ, encryption, logging |
| Secure by Default| Secure settings as baseline: HTTPS, JWT, hashed passwords      |

---

## 🔐 Authentication

- Email + password login (bcrypt / argon2 hashed)
- Session-based login (HTTP-only, SameSite cookies)
- OAuth2/OIDC token-based login (JWT)
- Refresh token flow for session renewal
- Multi-Factor Authentication (TOTP apps)
- Password policies: length, complexity, expiration

---

## 🔒 Authorization

- Role-Based Access Control (RBAC) via Casbin
- Roles can be assigned to users or groups
- Tenant-aware: all access is scoped to tenant
- Permissions mapped to actions/resources
- Matrix UI to visualize who can do what

---

## 🔑 API Security

- Bearer JWT tokens for all REST APIs
- HMAC or RSA signing (configurable)
- Token expiration (default: 15 mins)
- Token revocation support (logout, compromised keys)
- API token support for service-to-service auth
- Rate limiting for sensitive endpoints (e.g., login)

---

## 🗂️ Data Protection

- HTTPS enforced across all environments
- PostgreSQL with encrypted storage (where supported)
- Optional row-level encryption for sensitive fields
- No plaintext secrets checked into code or logs
- Encrypted JWT payloads (if required by policy)

---

## 🔍 Auditing & Monitoring

- All security-critical events logged:
  - Logins, logouts, permission changes, group edits, token use
- Fields include: actor, action, timestamp, IP, result
- Logs stored per tenant and protected from tampering
- Option to export logs to SIEM or ELK stack
- Alerts for suspicious patterns (future feature)

---

## 🛡️ Platform Hardening

- All Go services run as non-root
- Docker containers use distroless or `scratch` base
- Kubernetes limits (CPU, memory, autoscaling)
- Secrets injected via Vault or K8s Secrets (no env baked in)
- CI/CD pipelines scan for CVEs on build

---

## ✅ Common Threat Protections

| Threat                  | Mitigation                                               |
|-------------------------|----------------------------------------------------------|
| SQL Injection           | Parameterized queries, ORM enforcement                   |
| XSS                     | Auto-escaped React frontend, CSP headers                 |
| CSRF                    | SameSite cookies, token-based APIs                       |
| Brute Force / Enumeration | Rate limiting, login lockout after failures            |
| Session Hijacking       | Secure, HttpOnly, short-lived tokens                     |
| IDOR (Broken Access Control) | Tenant + permission check on every resource access   |

---

## 🔄 Secrets Management

- `.env` files for local development only (excluded from Git)
- Vault or AWS KMS for production secrets
- Secrets include:
  - DB credentials
  - JWT signing keys
  - SMTP creds
  - 3rd-party integrations

---

## 🧪 Secure Dev Practices

- Static analysis in CI
- Code review required on all PRs
- Security-specific integration tests
- Manual penetration testing for major releases
- Automated test coverage on authN/authZ

---

## 🗺️ Future Enhancements

- WebAuthn (FIDO2) support for passwordless login
- SCIM for secure automated user provisioning
- Tenant-specific audit log encryption
- Fine-grained token scopes
- Security awareness dashboard for admins

---

> Security is a process, not a feature. This document will evolve with every feature shipped.
