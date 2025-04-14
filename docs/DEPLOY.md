# IAM Platform - Deployment Guide

This guide outlines how to deploy the IAM system in local development, staging, and production environments. The architecture supports consistent Docker-based workflows and Kubernetes orchestration.

---

## 1. Environments Overview

| Environment | Description               | Stack         | DB         | Notes                            |
|-------------|---------------------------|---------------|------------|----------------------------------|
| `dev`       | Local development         | Docker Compose| SQLite/Postgres | Hot reload, debugging        |
| `staging`   | Pre-production test       | Kubernetes    | PostgreSQL | Replica of production config     |
| `prod`      | Live customer-facing      | Kubernetes    | PostgreSQL (Managed) | Secure, scaled, HA    |

---

## 2. Docker Compose (Local Dev)

### Prerequisites:
- Docker + Docker Compose
- Make (optional but useful)

### Quick Start:
```bash
make up
```

Will run:
- Go backend (localhost:8080)
- React frontend (localhost:3000)
- Postgres + Redis

### Tear Down:
```bash
make down
```

---

## 3. Kubernetes (Staging / Production)

### Option 1: Manual Deployment

- Build Docker images:
```bash
docker build -t registry/iam-backend ./backend
docker build -t registry/iam-frontend ./frontend
```

- Push images:
```bash
docker push registry/iam-backend
docker push registry/iam-frontend
```

- Apply manifests:
```bash
kubectl apply -f k8s/
```

### Option 2: Helm Charts (Recommended)

> Coming soon: `charts/iam/`

Helm will manage:
- Service deployment
- ConfigMap injection
- Secrets management
- Rolling updates

---

## 4. Configuration

Environment variables:

| Name            | Description                        |
|-----------------|------------------------------------|
| `DB_URL`        | Postgres connection string         |
| `REDIS_URL`     | Redis connection string            |
| `JWT_SECRET`    | Secret for HMAC signing (or public key) |
| `ENV`           | `development` / `production`       |
| `SMTP_HOST`     | Optional email server config       |

Secrets should be injected via:
- Docker `.env` file (dev)
- K8s Secrets or Vault (prod)

---

## 5. Persistent Storage

- **Dev**: Postgres volume via Docker
- **Prod**: Managed PostgreSQL (AWS RDS, GCP CloudSQL)
- Use Redis persistence if needed (`appendonly yes`)

---

## 6. Logging & Monitoring

- All services log JSON to stdout
- Recommended stack: EFK (Elasticsearch + Fluentd + Kibana)
- Use Prometheus/Grafana for metrics

---

## 7. CI/CD Pipeline

**Using GitHub Actions:**
- On PR: Lint, test, build
- On merge to `main`: Build + push Docker images
- On tag: Deploy to staging/prod

**Pipeline Stages:**
1. Checkout & cache deps
2. Lint backend/frontend
3. Run tests
4. Build images
5. Push to container registry
6. Trigger deploy (Helm or `kubectl apply`)

---

## 8. Future Cloud Provider Setup

- GCP: GKE + Cloud SQL + Workload Identity
- AWS: EKS + RDS + KMS
- Azure: AKS + Postgres + Azure Vault

---

## 9. TLS/HTTPS

- Use Cloudflare or ingress controller (e.g., NGINX + Cert-Manager)
- All traffic should terminate TLS at ingress
- Internal services can use mTLS for secure service-to-service auth

---

## 10. Backup & Disaster Recovery

- Postgres backups (daily snapshots + PITR)
- Redis optional persistence
- Docker volume backup in dev (for local state)

---

## 11. Maintenance Tasks

| Task             | Frequency | Tooling           |
|------------------|-----------|-------------------|
| DB Migrations    | On deploy | Go-Migrate / Flyway |
| Token Cleanup    | Daily     | Scheduled Job     |
| Log Rotation     | Daily     | Fluentd/Logrotate |
| Health Check     | Every 5s  | K8s liveness/readiness probes |

---

## 12. Access Control (Ops)

- Use GitHub Teams + branch protections
- Limit K8s access via `rbac.authorization.k8s.io`
- Use CI/CD tokens with minimal scope
- Secrets access via Vault with audit trail

---

> Always run infrastructure changes through the same CI/CD system for consistency and traceability.
