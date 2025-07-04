# 🛡️ Branch Protection Setup Instructions

## Setting Up Status Checks for Branch Protection

### Step 1: Go to Branch Protection Settings
1. Navigate to your GitHub repository
2. Go to `Settings` > `Branches`
3. Click `Add rule` or edit existing rule for `main` branch

### Step 2: Configure Required Status Checks
In the branch protection rule, enable these options:

✅ **Require a pull request before merging**
- Require approvals: 1
- Dismiss stale PR approvals when new commits are pushed

✅ **Require status checks to pass before merging**
- Require branches to be up to date before merging

**Add these EXACT status check names:**
```
Backend Tests
Frontend Tests
```

### Step 3: Additional Protection Rules (Recommended)
✅ **Restrict pushes that create files larger than 100MB**
✅ **Require linear history** (optional, keeps git history clean)
✅ **Include administrators** (applies rules to admins too)

### Step 4: Verify Setup
After setting up, the branch protection should look like this:

```
Branch protection rule for main:
├── Require a pull request before merging
│   ├── Require approvals: 1
│   └── Dismiss stale PR approvals when new commits are pushed
├── Require status checks to pass before merging
│   ├── Require branches to be up to date before merging
│   ├── Backend Tests ✅
│   └── Frontend Tests ✅
├── Restrict pushes that create files larger than 100MB
└── Include administrators
```

## Testing the Setup

### 1. Create a Test PR
```bash
git checkout -b test-branch-protection
echo "# Test change" >> README.md
git add README.md
git commit -m "test: verify branch protection"
git push origin test-branch-protection
```

### 2. Create Pull Request
- Go to GitHub and create a PR from `test-branch-protection` to `main`
- You should see status checks running
- The merge button should be disabled until checks pass

### 3. Verify Status Checks
The PR should show:
- ⏳ Backend Tests (running/pending)
- ⏳ Frontend Tests (running/pending)
- ✅ Backend Tests (after completion)
- ✅ Frontend Tests (after completion)

## Troubleshooting

### Status Checks Not Appearing?
1. **Check workflow names**: Must match exactly "Backend Tests" and "Frontend Tests"
2. **Trigger a workflow**: Push a commit to see if workflows run
3. **Check workflow files**: Ensure `.github/workflows/ci.yml` has correct job names

### Still Can't Find Status Checks?
1. Go to your repo's Actions tab
2. Run a workflow manually
3. After it completes, go back to branch protection settings
4. The status check names should now appear in the dropdown

### Common Issues:
- **Case sensitivity**: "Backend Tests" ≠ "backend tests"
- **Spaces matter**: "Backend Tests" ≠ "BackendTests"
- **Workflow must run first**: Status checks only appear after running at least once

## Alternative: Use Job IDs Instead of Names
If you prefer to use job IDs, update the branch protection to require:
```
test-backend
test-frontend
```

And remove the `name:` fields from the workflow jobs.

---

✅ **After following these steps, your branch protection will prevent merging until all tests pass!**