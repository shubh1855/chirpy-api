# Chirpy API

A Twitter-like social media backend written in Go.

Chirpy provides user authentication, JWT-based authorization, refresh tokens, user profile management, chirp creation and deletion, Chirpy Red memberships, and webhook integrations.

---

## Features

### Authentication & Security

- User registration
- User login
- Argon2 password hashing
- JWT access tokens
- Refresh token authentication
- Token revocation
- Protected API endpoints
- Authorization checks for resource ownership

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
- Chirpy Red membership support

### Admin

- Metrics endpoint
- Reset endpoint (development mode)

### Integrations

- Polka webhook integration
- API key validation for webhook security

---

## Tech Stack

- Go
- PostgreSQL
- SQLC
- Goose
- JWT
- Argon2id

---

## Project Structure

```text
.
├── internal
│   ├── auth
│   │   ├── bearer.go
│   │   ├── apikey.go
│   │   ├── jwt.go
│   │   ├── passwords.go
│   │   └── refresh.go
│   │
│   └── database
│
├── sql
│   ├── schema
│   └── queries
│
├── main.go
├── go.mod
└── README.md
```

---

## Requirements

- Go 1.24+
- PostgreSQL
- SQLC
- Goose

---

## Environment Variables

Create a `.env` file in the project root:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable

JWT_SECRET=your-secret-key

POLKA_KEY=f271c81ff7084ee5b99a5091b42d486e

PLATFORM=dev
```

---

## Installation

Clone the repository:

```bash
git clone https://github.com/<username>/chirpy-api.git

cd chirpy-api
```

Install dependencies:

```bash
go mod download
```

---

## Database Setup

Create the database:

```sql
CREATE DATABASE chirpy;
```

Run migrations:

```bash
goose postgres "$DB_URL" up
```

Generate SQLC code:

```bash
sqlc generate
```

---

## Running the Server

```bash
go run .
```

The server starts on:

```text
http://localhost:8080
```

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

Response:

```json
[
  {
    "id": "uuid",
    "body": "hello world",
    "user_id": "uuid"
  }
]
```

---

### Filter by Author

```http
GET /api/chirps?author_id=<user-id>
```

---

### Sort Ascending

```http
GET /api/chirps?sort=asc
```

---

### Sort Descending

```http
GET /api/chirps?sort=desc
```

---

### Combined Query

```http
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
```

```http
403 Forbidden
```

```http
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

Returns a simple HTML page containing application metrics.

---

### Reset

```http
POST /admin/reset
```

Development mode only.

---

## Polka Webhooks

### Upgrade User To Chirpy Red

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

## Authentication Schemes

### JWT Access Tokens

```http
Authorization: Bearer <jwt>
```

Used for:

- Creating chirps
- Updating users
- Deleting chirps

---

### Refresh Tokens

```http
Authorization: Bearer <refresh-token>
```

Used for:

- Refreshing access tokens
- Revoking refresh tokens

---

### Polka API Keys

```http
Authorization: ApiKey <polka-key>
```

Used for:

- Webhook authentication

---

## Future Improvements

- Chirp editing
- Pagination
- Rate limiting
- OpenAPI/Swagger documentation
- Docker support
- CI/CD pipeline
- Structured logging
- Role-based authorization

---

## License

This project was built as part of the Boot.dev Backend Development curriculum and expanded into a complete REST API backend in Go.
