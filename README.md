# Link Shortener

A simple and robust URL shortening service built with Go.

## Features
- **User Authentication**: Register, Login, Logout, and OTP verification.
- **JWT Authorization**: Secure endpoints using Access and Refresh tokens.
- **Link Shortening**: Create and manage shortened links.
- **PostgreSQL Database**: Data persistence using `sqlc` for type-safe queries.
- **Go Chi Router**: Fast and lightweight HTTP routing.

## Tech Stack
- Go 1.20+
- [Chi](https://github.com/go-chi/chi) Router
- PostgreSQL
- [sqlc](https://sqlc.dev/)
- [golang-jwt](https://github.com/golang-jwt/jwt)
- bcrypt for password hashing

## Prerequisites
- Go installed on your machine
- PostgreSQL running locally or remotely

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository_url>
   cd link-shortner
   ```

2. **Environment Variables:**
   Create a `.env` file in the root directory with the following variables:
   ```env
   SERVER_PORT=8080
   GOOSE_DBSTRING="postgres://user:password@localhost:5432/link_shortener?sslmode=disable"
   JWT_SECRET="your_super_secret_key"
   ```

3. **Database Setup:**
   Make sure you have Goose installed for database migrations, or run the raw SQL schema inside the `internal/database` folder.
   
4. **Run the Server:**
   ```bash
   go run cmd/server/main.go
   ```

## API Endpoints

### Authentication & Users
- `POST /v1/users/` - Register a new user
- `POST /v1/users/login` - Login and receive `access_token` and `refresh_token`
- `GET /v1/users/` - List users
- `GET /v1/users/{id}` - Get user by ID
- `DELETE /v1/users/{id}` - Delete user
- `POST /v1/refresh-token` - Issue a new access token using a refresh token

### Protected Routes (Requires Bearer Token)
- `POST /v1/link` - Create a new short link
- `POST /v1/logout` - Logout (invalidates refresh token)

## License
MIT
