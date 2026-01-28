# Employee Management API

Backend API built using Go and PostgreSQL.

This project is a take-home assignment focused on building a production-ready,
API-only backend with authentication, employee management, and salary metrics.

## Tech Stack
- Go
- PostgreSQL
- JWT Authentication
- Docker

## API Versioning

All APIs are exposed under a versioned base path.

Current version:
- v1 → `/api/v1`

This allows future iterations without breaking existing clients.

## Database Setup

PostgreSQL is run locally using Docker Compose.

Start PostgreSQL:

```bash
docker compose up -d
```

Ensure the database container is running before starting the application.

## Database Migrations

This project uses versioned SQL migrations managed with `golang-migrate`.
Migrations are run manually to keep schema changes explicit and controlled.

### Run migrations locally (CLI)

If you have the `migrate` CLI installed, run:

```bash
migrate -path internal/migrations \
  -database "postgres://user:secret@localhost:5432/emp_management?sslmode=disable" \
  up
```

### Run migrations using Docker (recommended)

```bash
docker run --rm \
  --network employee-management-api_default \
  -v $(pwd)/internal/migrations:/migrations \
  migrate/migrate \
  -path /migrations \
  -database "postgres://user:secret@employee_management_postgres:5432/emp_management?sslmode=disable" \
  up
```

This runs migrations inside a temporary container and connects to PostgreSQL via the Docker network.