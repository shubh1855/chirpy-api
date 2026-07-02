<p align="center">
  A Twitter-like social media backend built with Go and PostgreSQL.
</p>

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/shubh1855/chirpy-api" alt="Go Version">
  <img src="https://img.shields.io/github/v/release/shubh1855/chirpy-api" alt="Release">
  <img src="https://img.shields.io/github/actions/workflow/status/shubh1855/chirpy-api/release.yml?branch=release" alt="CI">
  <img src="https://img.shields.io/github/license/shubh1855/chirpy-api" alt="License">
  <img src="https://img.shields.io/github/last-commit/shubh1855/chirpy-api" alt="Last Commit">
</p>

---

## Overview

Chirpy API is a production-style REST backend inspired by Twitter. The project was built as part of the Boot.dev Backend Path and expanded into a complete social media API featuring authentication, authorization, refresh tokens, webhooks, and PostgreSQL persistence.

The project demonstrates modern backend engineering practices in Go, including:

- JWT authentication and authorization
- Refresh token sessions
- Argon2 password hashing
- Type-safe database queries with SQLC
- Database migrations with Goose
- Webhook integrations
- Automated releases with GitHub Actions RESTful API design

---

## Features

### Authentication & Security

- User registration
- User login
- Argon2 password hashing
- JWT access tokens
- Refresh token authentication
- Refresh token revocation
- Protected endpoints
- Resource ownership authorization

### Chirps

- Create chirps
- Retrieve all chirps
- Retrieve a chirp by ID
- Delete chirps
- Filter chirps by author
- Sort chirps by creation date
- Profanity filtering

### User Management

- Register users
- Update email and password
- Chirpy Red memberships

### Admin

- Health check endpoint
- Metrics endpoint
- Development reset endpoint

### Integrations

- Polka webhook integration
- API key authentication for webhooks

---

## Tech Stack

| Category            | Technology     |
| ------------------- | -------------- |
| Language            | Go             |
| Database            | PostgreSQL     |
| Authentication      | JWT            |
| Password Hashing    | Argon2id       |
| Query Generation    | SQLC           |
| Database Migrations | Goose          |
| CI/CD               | GitHub Actions |
| API Style           | REST           |

---

## Architecture

```text
Client
   │
   ▼
REST API (Go)
   │
   ├── Authentication (JWT + Refresh Tokens)
   ├── Business Logic
   ├── Webhooks
   └── PostgreSQL
```

---

## Project Structure

```text
.
├── assets
│   └── logo.png
├── internal
│   ├── auth
│   │   ├── apikey.go
│   │   ├── bearer.go
│   │   ├── jwt.go
│   │   ├── passwords.go
│   │   └── refresh.go
│   └── database
│       ├── chirps.sql.go
│       ├── refresh_tokens.sql.go
│       ├── users.sql.go
│       └── models.go
├── sql
│   ├── queries
│   └── schema
├── main.go
├── go.mod
├── sqlc.yaml
└── README.md
```

---

## Requirements

- Go 1.26+
- PostgreSQL
- SQLC
- Goose

---

## Environment Variables

Create a `.env` file:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable

JWT_SECRET=your-super-secret-jwt-key

POLKA_KEY=f271c81ff7084ee5b99a5091b42d486e

PLATFORM=dev
```

> [!IMPORTANT]
> Never commit your real `.env` file. Create a `.env.example` file and commit that instead.

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/shubh1855/chirpy-api.git
cd chirpy-api
```

### Install Dependencies

```bash
go mod download
```

### Create the Database

```sql
CREATE DATABASE chirpy;
```

### Run Migrations

```bash
goose postgres "$DB_URL" up
```

### Generate SQLC Code

```bash
sqlc generate
```

### Start the Server

```bash
go run .
```

The server will be available at:

```text
http://localhost:8080
```

---

## Running Tests

```bash
go test ./...
```

---

# API Summary

| Method | Endpoint                | Description               |
| ------ | ----------------------- | ------------------------- |
| POST   | `/api/users`            | Register a user           |
| POST   | `/api/login`            | Login                     |
| POST   | `/api/refresh`          | Refresh access token      |
| POST   | `/api/revoke`           | Revoke refresh token      |
| PUT    | `/api/users`            | Update user               |
| POST   | `/api/chirps`           | Create chirp              |
| GET    | `/api/chirps`           | Get all chirps            |
| GET    | `/api/chirps/{chirpID}` | Get chirp by ID           |
| DELETE | `/api/chirps/{chirpID}` | Delete chirp              |
| POST   | `/api/polka/webhooks`   | Process webhook           |
| GET    | `/api/healthz`          | Health check              |
| GET    | `/admin/metrics`        | Metrics                   |
| POST   | `/admin/reset`          | Reset database (dev only) |

---

# API Reference

## Authentication

### Register User

```http
POST /api/users
```

Request:

```json
{
  "email": "user@example.com",
  "password": "super-secret-password"
}
```

Response:

```json
{
  "id": "uuid",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

---

### Login

```http
POST /api/login
```

Request:

```json
{
  "email": "user@example.com",
  "password": "super-secret-password"
}
```

Response:

```json
{
  "id": "uuid",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z",
  "email": "user@example.com",
  "is_chirpy_red": false,
  "token": "<jwt>",
  "refresh_token": "<refresh-token>"
}
```

---

### Refresh Access Token

```http
POST /api/refresh
```

Headers:

```http
Authorization: Bearer <refresh-token>
```

Response:

```json
{
  "token": "<new-jwt>"
}
```

---

### Revoke Refresh Token

```http
POST /api/revoke
```

Headers:

```http
Authorization: Bearer <refresh-token>
```

Response:

```http
204 No Content
```

---

## Users

### Update User

```http
PUT /api/users
```

Headers:

```http
Authorization: Bearer <jwt>
```

Request:

```json
{
  "email": "new@example.com",
  "password": "new-password"
}
```

Response:

```json
{
  "id": "uuid",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-02T00:00:00Z",
  "email": "new@example.com",
  "is_chirpy_red": false
}
```

---

## Chirps

### Create Chirp

```http
POST /api/chirps
```

Headers:

```http
Authorization: Bearer <jwt>
```

Request:

```json
{
  "body": "Hello Chirpy!"
}
```

Response:

```json
{
  "id": "uuid",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z",
  "body": "Hello Chirpy!",
  "user_id": "uuid"
}
```

---

### Get All Chirps

```http
GET /api/chirps
```

Optional query parameters:

```text
author_id=<user-id>
sort=asc
sort=desc
```

Examples:

```http
GET /api/chirps
GET /api/chirps?sort=desc
GET /api/chirps?author_id=<user-id>
GET /api/chirps?author_id=<user-id>&sort=desc
```

---

### Get Chirp By ID

```http
GET /api/chirps/{chirpID}
```

---

### Delete Chirp

```http
DELETE /api/chirps/{chirpID}
```

Headers:

```http
Authorization: Bearer <jwt>
```

Responses:

```http
204 No Content
403 Forbidden
404 Not Found
```

---

## Admin Endpoints

### Health Check

```http
GET /api/healthz
```

Response:

```text
OK
```

---

### Metrics

```http
GET /admin/metrics
```

Returns an HTML page showing application metrics.

---

### Reset

```http
POST /admin/reset
```

Development mode only.

---

## Polka Webhooks

### Upgrade User to Chirpy Red

```http
POST /api/polka/webhooks
```

Headers:

```http
Authorization: ApiKey <polka-api-key>
```

Request:

```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "uuid"
  }
}
```

Response:

```http
204 No Content
```

---

# Authentication Schemes

## JWT Access Tokens

```http
Authorization: Bearer <jwt>
```

Used for:

- Creating chirps
- Updating users
- Deleting chirps

---

## Refresh Tokens

```http
Authorization: Bearer <refresh-token>
```

Used for:

- Refreshing access tokens
- Revoking refresh tokens

---

## Polka API Keys

```http
Authorization: ApiKey <polka-key>
```

Used for:

- Webhook authentication

---

## CI/CD

The project uses GitHub Actions for automated releases.

Every merge into the `release` branch:

- Runs tests
- Builds the application
- Generates the next semantic version
- Creates a GitHub Release
- Publishes release binaries

---

## Roadmap

- [ ] Chirp editing
- [ ] Pagination
- [ ] Rate limiting
- [ ] Docker support
- [ ] OpenAPI/Swagger documentation
- [ ] Structured logging
- [ ] Role-based authorization

---

## License

This project was built as part of the Boot.dev Backend Development curriculum and expanded into a complete REST API backend in Go.
