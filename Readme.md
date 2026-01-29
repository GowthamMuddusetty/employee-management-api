# Employee Management API

A production-ready backend API built using Go, Gin, and PostgreSQL.

This project is a take-home assignment focused on designing and implementing
a clean, maintainable, API-only backend with authentication, employee
management, salary calculation, and salary metrics.

---

## Tech Stack

- Go
- Gin (HTTP framework)
- PostgreSQL
- JWT Authentication
- Docker & Docker Compose

---

## API Versioning

All APIs are exposed under a versioned base path.

Current version:
- **v1** → `/api/v1`

This allows future iterations without breaking existing clients.

---

## Running the Application

### Prerequisites
- Go 1.21+
- Docker & Docker Compose

---

## Database & Application Setup (Docker)

Both PostgreSQL and the API can be run using Docker Compose.

### Start services

```bash
docker compose up --build -d
```

Docker Compose starts both the database and API services together.

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
docker run --rm `
  --network employee-management-api_default `
  -v ${PWD}\internal\migrations:/migrations `
  migrate/migrate `
  -path /migrations `
  -database "postgres://user:secret@employee_management_postgres:5432/emp_management?sslmode=disable" `
  up
```

This runs migrations inside a temporary container and connects to PostgreSQL via the Docker network.

### Running the Server

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=emp_management
DB_USER=user
DB_PASSWORD=secret

JWT_SECRET=emp_secret_key
JWT_EXPIRY=1h
```

Start the server:
```bash
go run ./cmd/server
```

Server runs at:
`http://localhost:8080`

## Authentication
JWT-based authentication is used.

## Auth APIs
- POST /api/v1/auth/register
- POST /api/v1/auth/login
On successful login, a JWT token is issued.

All non-auth endpoints require:
`Authorization: Bearer <token>`

## Employee Management APIs (Protected)
- POST /api/v1/employees
- GET /api/v1/employees
- GET /api/v1/employees/{id}
- PUT /api/v1/employees/{id}
- DELETE /api/v1/employees/{id}
These APIs manage employee records including name, job title, country, and salary.

### Salary APIs (Protected)

## Salary Calculation
`GET /api/v1/employees/{id}/salary`

Salary deduction rules:
  India → 10% deduction
  United States → 12% deduction
  Other countries → No deduction

Response includes:
  Gross salary
  Deduction amount
  Net salary

### Salary Metrics APIs (Protected)
  `GET /api/v1/metrics/salary/country/{country}`
    Returns min, max, and average salary for a country

  `GET /api/v1/metrics/salary/job-title/{jobTitle}`
    Returns average salary for a job title
  
## Postman Collection
A Postman collection is included to test all APIs end-to-end.

## ERD
- Users are used for authentication
- Employees represent business domain data

```mermaid
erDiagram
  USERS {
    uuid id PK
    string email
    string password_hash
    timestamp created_at
  }

  EMPLOYEES {
    uuid id PK
    string full_name
    string job_title
    string country
    float salary
    timestamp created_at
    timestamp updated_at
  }
