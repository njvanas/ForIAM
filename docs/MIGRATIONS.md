# IAM Platform - Migrations

This document describes the structure and strategy for applying database schema migrations using tools like `golang-migrate`, including versioning, rollback, and environment safety tips.

- Use timestamped SQL files: `0001_init_schema.up.sql`, `0001_init_schema.down.sql`
- Always test migrations locally before pushing
- Track applied migrations in CI logs
- Store migration files in `/migrations` directory
