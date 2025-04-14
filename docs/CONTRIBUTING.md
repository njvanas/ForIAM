# 🧑‍💻 Contributing to the IAM Platform

We welcome contributions to help us build a secure, modern IAM system.

---

## 🚀 Getting Started

1. Fork this repo
2. Clone your fork
3. Install Docker, Node.js (for frontend), Go (for backend)
4. Run the full stack:
```bash
make up
```

---

## 🛠 Development Workflow

- Work from feature branches: `feat/feature-name`, `fix/bug-name`
- Keep PRs small and focused
- Use [Conventional Commits](https://www.conventionalcommits.org)

---

## ✅ Pre-Commit Checklist

- [ ] `make lint` passes (backend & frontend)
- [ ] `make test` passes
- [ ] No secrets or credentials committed
- [ ] Updated relevant docs if needed
- [ ] Tests written for new features or fixes

---

## 🧪 Testing

- Unit + integration tests
- Backend: Go test suite with coverage
- Frontend: React + Vite test runner

---

## 🧱 Code Standards

### Backend (Go)
- Use idiomatic Go practices
- Prefer interfaces for modularity
- One feature per package

### Frontend (React)
- TSX + TypeScript types required
- Folder-per-feature structure
- Minimal UI logic in components

---

## 🔐 Security Guidelines

- Never log passwords or tokens
- Never commit `.env` files or secrets
- Use secure headers and access controls

---

## 🙏 Thanks

We appreciate all PRs, issue reports, feedback, and contributions!
