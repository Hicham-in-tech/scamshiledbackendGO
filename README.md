# ScamShield Backend v2 (Go)

A high-performance backend service for URL analysis and email security review using Go, Gin, and GORM.

## 📋 Features

- **URL Risk Analysis** - Analyze URLs for phishing and malware risks
- **Email Security Review** - Review emails for phishing and scam indicators
- **User Management** - User registration and authentication
- **RESTful API** - Clean and documented API endpoints
- **Database Persistence** - MySQL integration with GORM ORM
- **CORS Support** - Ready for frontend integration

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: MySQL with GORM ORM
- **Authentication**: Header-based token authentication
- **Environment**: godotenv for configuration

## 📁 Project Structure

```
backendv2/
├── app/
│   ├── handlers/       # HTTP request handlers
│   ├── models/         # Data models (User, Scan, EmailReview)
│   ├── services/       # Business logic services
│   └── middleware/     # Authentication and CORS middleware
├── config/            # Database configuration
├── database/          # Database migrations and seeders
├── .env               # Environment variables (local)
├── .env.example       # Environment variables template
├── main.go            # Application entry point
└── go.mod            # Go module dependencies
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- MySQL 5.7+
- Git

### Installation

1. Clone and navigate to the directory:
```bash
cd nexjsproject/backendv2
```

2. Install dependencies:
```bash
go mod download
go mod tidy
```

3. Setup environment variables:
```bash
cp .env.example .env
# Edit .env with your database credentials
```

4. Create database:
```bash
# Make sure MySQL is running, then create the database:
mysql -u root -p < ../scamshield_db.sql
```

5. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8000`

## 📚 API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login user

### Scans (Protected)

- `POST /api/scans` - Create a new URL scan
- `GET /api/scans` - Get all scans for user
- `GET /api/scans/:id` - Get specific scan

### Email Reviews (Protected)

- `POST /api/emails/review` - Review an email for security
- `GET /api/emails` - Get all email reviews
- `GET /api/emails/:id` - Get specific review

### Health Check

- `GET /health` - Server health status

## 🔐 Authentication

Protected endpoints require:
- `Authorization` header with token
- `X-User-ID` header with user ID

Example:
```bash
curl -H "Authorization: Bearer token" \
     -H "X-User-ID: 1" \
     http://localhost:8000/api/scans
```

## 📦 Building for Production

```bash
# Build binary
go build -o scamshield-backend main.go

# Run binary
./scamshield-backend
```

## 📝 Environment Variables

See `.env.example` for all available configuration options.

## 🤝 Contributing

To extend the backend:
1. Add new models in `app/models/`
2. Create handlers in `app/handlers/`
3. Add business logic in `app/services/`
4. Update routes in `main.go`

## 📄 License

MIT License
