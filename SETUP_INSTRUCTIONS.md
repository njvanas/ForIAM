# 🚀 Complete Setup Instructions for ForIAM CI/CD

## Required GitHub Repository Settings

### 1. **Secrets Configuration**
Add these secrets in GitHub: `Settings > Secrets and variables > Actions`

```bash
# Docker Hub credentials (REQUIRED for Docker builds)
DOCKER_USERNAME=your-dockerhub-username
DOCKER_PASSWORD=your-dockerhub-password-or-token

# Optional: Code coverage
CODECOV_TOKEN=your-codecov-token  # Get from codecov.io
```

### 2. **Environment Configuration**
Create environments in GitHub: `Settings > Environments`

**Create two environments:**
- `staging` 
- `production`

**For each environment, configure:**
- **Protection rules**: Require reviewers for production
- **Environment secrets** (if different from global):
  - `KUBECONFIG` (base64 encoded kubeconfig for deployment)
  - `DEPLOY_URL` (deployment target URL)

### 3. **Branch Protection Rules**
Go to `Settings > Branches` and protect the `main` branch:

```yaml
Branch protection rules for 'main':
✅ Require a pull request before merging
✅ Require status checks to pass before merging
  - Required status checks: "test"
✅ Require branches to be up to date before merging
✅ Restrict pushes that create files larger than 100MB
```

## Docker Hub Setup

### 1. **Create Docker Hub Account**
- Go to https://hub.docker.com
- Create account or login
- Create repositories:
  - `foriam/backend`
  - `foriam/frontend`

### 2. **Generate Access Token**
- Go to Docker Hub > Account Settings > Security
- Create new access token
- Use this as `DOCKER_PASSWORD` secret (more secure than password)

## Workflow Behavior After Setup

### **Automatic Triggers:**

1. **Push to any branch** → CI tests run
2. **Push to main** → CI tests + Docker build + Deploy to staging
3. **Pull Request** → CI tests + Docker build (no push)

### **Manual Triggers:**

1. **Manual Deploy**: 
   ```
   Actions > Deploy + Coverage + Docker > Run workflow
   Choose: staging or production
   ```

2. **Create Release**:
   ```
   Actions > Release > Run workflow
   Version: v1.0.0
   Release notes: "Initial release"
   ```

## Optional Enhancements

### 1. **Kubernetes Deployment** (if using K8s)
Add to environment secrets:
```bash
KUBECONFIG=<base64-encoded-kubeconfig>
KUBE_NAMESPACE=foriam-staging  # or foriam-production
```

Update deploy step in `.github/workflows/deploy.yml`:
```yaml
- name: Deploy to Kubernetes
  run: |
    echo "${{ secrets.KUBECONFIG }}" | base64 -d > kubeconfig
    export KUBECONFIG=kubeconfig
    kubectl set image deployment/backend backend=foriam/backend:${{ needs.test-and-build.outputs.version }} -n ${{ secrets.KUBE_NAMESPACE }}
    kubectl set image deployment/frontend frontend=foriam/frontend:${{ needs.test-and-build.outputs.version }} -n ${{ secrets.KUBE_NAMESPACE }}
    kubectl rollout status deployment/backend -n ${{ secrets.KUBE_NAMESPACE }}
    kubectl rollout status deployment/frontend -n ${{ secrets.KUBE_NAMESPACE }}
```

### 2. **Slack/Discord Notifications**
Add webhook URLs as secrets and notification steps to workflows.

### 3. **Security Scanning**
The workflows already include dependency scanning. For additional security:
- Enable GitHub Advanced Security
- Add Snyk or other security scanning tools

## Testing the Setup

### 1. **Test CI Pipeline**
```bash
# Create a test branch
git checkout -b test-ci
echo "# Test change" >> README.md
git add README.md
git commit -m "test: trigger CI pipeline"
git push origin test-ci

# Create PR to main - should trigger CI
```

### 2. **Test Docker Build**
```bash
# Push to main branch - should trigger full pipeline
git checkout main
git merge test-ci
git push origin main
```

### 3. **Test Release Process**
```bash
# Go to GitHub Actions > Release > Run workflow
# Enter version: v0.1.0
# Should create tag, release, and trigger production deployment
```

## Troubleshooting

### Common Issues:

1. **Docker push fails**: Check DOCKER_USERNAME and DOCKER_PASSWORD secrets
2. **Tests fail**: Ensure all dependencies are properly specified in package.json/go.mod
3. **Deploy fails**: Check environment configuration and secrets
4. **Permission denied**: Ensure GitHub token has proper permissions

### Debug Commands:
```bash
# Check workflow logs in GitHub Actions tab
# For local testing:
docker build -t test-backend ./backend
docker build -t test-frontend ./frontend
```

## Next Steps After Setup

1. **Monitor first few runs** to ensure everything works
2. **Set up monitoring** for deployed applications
3. **Configure alerts** for failed deployments
4. **Document deployment procedures** for your team
5. **Set up staging environment** that mirrors production

---

✅ **Once you complete these steps, your CI/CD pipeline will be fully functional!**