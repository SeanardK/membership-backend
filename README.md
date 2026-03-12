# Go Backend API template

A robust RESTful API backend service for managing portfolio projects, built with Go and designed for modern web applications. Features secure authentication, image upload capabilities, and full CRUD operations for portfolio management.

## Quick Start

```bash
# Clone the repository
git clone https://github.com/SeanardK/go-backend-template
cd go-backend-template

# Copy environment file and configure
cp .env.example .env
# Edit .env with your database and Keycloak settings

# Install dependencies and run
go mod download
go run cmd/main.go
```

Or using Docker:
```bash
docker-compose up -d
```

## Features

- **RESTful API Design** - Clean and intuitive endpoints following REST principles
- **OIDC Authentication** - Secure authentication using Keycloak with JWT tokens
- **Image Upload** - Support for portfolio project image uploads
- **PostgreSQL Database** - Reliable data persistence with GORM ORM
- **Docker Support** - Containerized deployment with Docker Compose
- **CORS Enabled** - Cross-Origin Resource Sharing configuration
- **Structured Logging** - Comprehensive logging with Logrus
- **Auto Migration** - Automatic database schema migration on startup
- **File Management** - Automatic cleanup of associated files on deletion
- **Timezone Support** - Configured for Asia/Jakarta timezone (UTC storage)

## Table of Contents

- [Requirements](#requirements)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Development](#development)
- [Docker Deployment](#docker-deployment)
- [Authentication](#authentication)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Requirements

- Go 1.25.4 or higher
- PostgreSQL 13+
- Keycloak Server (for authentication)
- Docker & Docker Compose (optional, for containerized deployment)

## Tech Stack

- **[Go](https://golang.org/)** - Programming language
- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[GORM](https://gorm.io/)** - ORM library for Go
- **[PostgreSQL](https://www.postgresql.org/)** - Relational database
- **[Keycloak](https://www.keycloak.org/)** - Identity and access management
- **[go-oidc](https://github.com/coreos/go-oidc)** - OpenID Connect client
- **[Logrus](https://github.com/sirupsen/logrus)** - Structured logger
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable loader
- **[Docker](https://www.docker.com/)** - Containerization platform

## Installation

### Local Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/SeanardK/go-backend-template
   cd go-backend-template
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your configuration (see [Configuration](#configuration))

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
# Server Configuration
PORT=3001

# Keycloak Configuration
KEYCLOAK_BASE_URL=http://localhost:8080
KEYCLOAK_REALM=your-realm
CLIENT_ID=your-client-id

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-password
DB_NAME=membership_db
DB_SSLMODE=disable

# CORS Configuration (optional)
# Note: Currently set to allow all origins in development
# ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

### Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `PORT` | Server port number | No | 3001 |
| `KEYCLOAK_BASE_URL` | Keycloak server base URL | Yes | - |
| `KEYCLOAK_REALM` | Keycloak realm name | Yes | - |
| `CLIENT_ID` | Keycloak client ID | Yes | - |
| `DB_HOST` | PostgreSQL host | Yes | - |
| `DB_PORT` | PostgreSQL port | Yes | - |
| `DB_USER` | Database username | Yes | - |
| `DB_PASSWORD` | Database password | Yes | - |
| `DB_NAME` | Database name | Yes | - |
| `DB_SSLMODE` | PostgreSQL SSL mode | No | disable |

## Running the Application

### Development Mode

```bash
go run cmd/main.go
```

### Production Build

```bash
go build -o portfolio-api cmd/main.go
./portfolio-api
```

The server will start on the configured port (default: 3001).

## API Endpoints

Base URL: `http://localhost:3001`

### Response Format

All API responses follow a consistent JSON structure:

**Success Response:**
```json
{
  "message": "Operation description",
  "data": { /* result data */ }
}
```

**Error Response:**
```json
{
  "message": "Error description",
  "error": "Detailed error message (optional)"
}
```

### Static Files

Images are served from:
```
GET /public/portfolio/images/<filename>
```

## 📁 Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── pkg/
│   ├── config/
│   │   └── connection.go       # Database connection configuration
│   ├── controller/
│   │   └── portfolio.go        # Portfolio controller (handlers)
│   ├── database/
│   │   └── main.go            # Database initialization & migration
│   ├── middleware/
│   │   └── auth.go            # Authentication middleware (OIDC)
│   ├── model/
│   │   └── portfolio.go       # Portfolio data model
│   ├── routes/
│   │   ├── index.go           # Route registration
│   │   └── portfolio.go       # Portfolio routes
│   └── utils/
│       ├── env.go             # Environment variable utilities
│       └── file.go            # File upload utilities
├── public/
│   └── portfolio/
│       └── images/            # Uploaded images storage
├── basic/                     # Basic/example files
├── docker-compose.yml         # Docker Compose configuration
├── dockerfile                 # Dockerfile for container build
├── go.mod                     # Go module dependencies
├── go.sum                     # Dependency checksums
└── README.md                  # This file
```

## Development

### Database Migration

The application automatically runs database migrations on startup. The schema is defined in the model files.

### Adding New Endpoints

1. Create a model in `pkg/model/`
2. Create a controller in `pkg/controller/`
3. Define routes in `pkg/routes/`
4. Register routes in `pkg/routes/index.go`

### File Upload Limits

The application has a maximum multipart memory size of **8MB** for file uploads. This is configured in the main application file and can be adjusted as needed.

## Docker Deployment

### Using Docker Compose

1. **Build and run the container**
   ```bash
   docker-compose up -d
   ```

2. **View logs**
   ```bash
   docker-compose logs -f backend
   ```

3. **Stop the container**
   ```bash
   docker-compose down
   ```

### Using Docker

1. **Build the image**
   ```bash
   docker build -t portfolio-backend .
   ```

2. **Run the container**
   ```bash
   docker run -d \
     -p 3001:3001 \
     --env-file .env \
     -v $(pwd)/public:/root/public \
     --name portfolio-backend \
     portfolio-backend
   ```

### Production Deployment

For production deployment, consider:

- **Security**: Change `AllowAllOrigins` to `false` and specify allowed origins in CORS configuration
- Using a reverse proxy (Nginx, Traefik)
- Setting up SSL/TLS certificates
- Configuring proper CORS origins for your frontend domain
- Using managed PostgreSQL service
- Implementing rate limiting
- Setting up monitoring and logging (e.g., Prometheus, Grafana)
- Using environment-specific configuration
- Implementing health check endpoints
- Securing file upload directory with proper permissions
- Regular database backups

## Authentication

This API uses Keycloak for authentication with OpenID Connect (OIDC). Protected endpoints require a valid JWT token.

### Obtaining a Token

1. Configure a Keycloak client for your application
2. Use the Keycloak authentication endpoints to obtain a token
3. Include the token in the `Authorization` header:
   ```
   Authorization: Bearer <your-jwt-token>
   ```

## Troubleshooting

### Database Connection Issues

If you encounter database connection errors:
1. Verify PostgreSQL is running: `psql -U postgres -l`
2. Check your `.env` file has correct credentials
3. Ensure the database exists: `CREATE DATABASE membership_db;`
4. Verify SSL mode is set correctly (use `disable` for local development)

### Keycloak Authentication Issues

If authentication fails:
1. Verify Keycloak server is running and accessible
2. Check `KEYCLOAK_BASE_URL`, `KEYCLOAK_REALM`, and `CLIENT_ID` are correct
3. Ensure your Keycloak client is configured for OIDC
4. Verify the JWT token is valid and not expired

### Image Upload Issues

If image uploads fail:
1. Check the `public/portfolio/images/` directory exists and is writable
2. Verify the file size is under 8MB
3. Ensure the content type is set to `multipart/form-data`
4. Check server logs for detailed error messages

### Port Already in Use

If port 3001 is already in use:
1. Change the `PORT` in your `.env` file
2. Or kill the process using the port: `netstat -ano | findstr :3001` (Windows)

## Author
- [**Seanard K**](https://github.com/SeanardK)