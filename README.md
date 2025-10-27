# Go Echo REST API Template

A clean, scalable REST API template built with Go and Echo framework, following Clean Architecture principles.

## 🚀 Features

- ✅ **Clean Architecture** - Separation of concerns with clear layer boundaries
- ✅ **Echo Framework v4** - High performance, minimalist web framework
- ✅ **JWT Authentication** - Secure token-based authentication
- ✅ **MySQL/SQLite Support** - Flexible database options
- ✅ **Request Validation** - Built-in validation using go-playground/validator
- ✅ **CORS Support** - Configurable cross-origin resource sharing
- ✅ **Graceful Shutdown** - Proper cleanup on server termination
- ✅ **Structured Logging** - Request logging with custom formats
- ✅ **Error Handling** - Consistent error responses
- ✅ **Environment Config** - Environment-based configuration

## 📋 Prerequisites

- Go 1.24.3 or higher
- MySQL 8.0+ or SQLite 3
- Git

## 🛠️ Tech Stack

| Category | Technology |
|----------|-----------|
| **Language** | Go 1.24.3 |
| **Web Framework** | Echo v4 |
| **Database** | MySQL / SQLite |
| **ORM/Query Builder** | sqlx |
| **Authentication** | JWT (golang-jwt/jwt) |
| **Validation** | go-playground/validator |
| **Password Hashing** | bcrypt (golang.org/x/crypto) |
| **Environment** | godotenv |
| **ID Generation** | ksuid |

## 📁 Project Structure

```
be-go-echo/
├── internal/           # Private application code
│   ├── entity/        # Domain models
│   ├── handlers/      # HTTP handlers
│   ├── middlewares/   # HTTP middlewares
│   ├── repositories/  # Data access layer
│   ├── request/       # Request/Response DTOs
│   ├── route/         # API routes
│   └── services/      # Business logic
├── pkg/               # Public packages
│   ├── common/        # Shared utilities
│   ├── configs/       # Configuration
│   ├── database/      # Database connections
│   └── util/          # Utilities
├── scripts/           # SQL scripts
├── server/            # Server setup
└── main.go            # Entry point
```

See [FOLDER_STRUCTURE.md](FOLDER_STRUCTURE.md) for detailed structure.

## 🚦 Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd be-go-echo
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Setup Environment Variables

Copy the example environment file and configure it:

```bash
copy _env_example .env
```

Edit `.env` with your configuration:

```env
# Server
PORT_API=1234
ALLOW_ORIGINS=*
ALLOW_METHODS=GET,POST,PUT,DELETE,OPTIONS
ALLOW_HEADERS=Origin,Content-Type,Accept,Authorization,X-Www-Form-Urlencoded

# Logger
LOG_TIMEZONE=Asia/Bangkok
LOG_TIME_FORMAT=2006-01-02 15:04:05
LOG_FORMAT=${ip} - [${time}] "${method} ${url} ${protocol}" ${status} ${latency}

# Database (MySQL)
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=your_database
DB_CONN_MAX_IDLE_TIME=300
DB_CONNECTION_MAX_LIFE_TIME=300
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=20

# Authentication
SECRET_KEY=your-secret-key-here
TOKEN_EXPIRE=168
```

### 4. Setup Database

#### Using MySQL

```bash
# Create database
mysql -u root -p
CREATE DATABASE your_database;

# Run migration script
mysql -u root -p your_database < scripts/users.sql
```

#### Using SQLite

The SQLite database will be created automatically on first run.

### 5. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:1234` (or your configured port).

## 📡 API Endpoints

### Authentication

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/v1/register` | Register new user | ❌ |
| POST | `/api/v1/login` | User login | ❌ |

### Users

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/v1/users/filter` | Get filtered users | ✅ |

## 🔐 Authentication

This API uses JWT (JSON Web Tokens) for authentication.

### Register a User

```bash
curl -X POST http://localhost:1234/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "username": "johndoe",
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

### Login

```bash
curl -X POST http://localhost:1234/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "securepassword"
  }'
```

Response:
```json
{
  "status": "success",
  "message": "login success",
  "data": {
    "user_id": "...",
    "username": "johndoe",
    "email": "john@example.com"
  },
  "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Using the Token

Include the token in the Authorization header:

```bash
curl -X POST http://localhost:1234/api/v1/users/filter \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "page": 1,
    "limit": 10
  }'
```

## 🏗️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

```
┌─────────────┐
│   Handler   │ ← HTTP Layer (Request/Response)
└──────┬──────┘
       │
┌──────▼──────┐
│   Service   │ ← Business Logic Layer
└──────┬──────┘
       │
┌──────▼──────┐
│ Repository  │ ← Data Access Layer
└──────┬──────┘
       │
┌──────▼──────┐
│  Database   │ ← Data Storage
└─────────────┘
```

### Layer Responsibilities

- **Handler**: Handles HTTP requests, validates input, calls services
- **Service**: Implements business logic, orchestrates operations
- **Repository**: Manages data access, database queries
- **Entity**: Defines domain models

See [RULES.md](RULES.md) for detailed guidelines.

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## 📦 Building for Production

### Build Binary

```bash
# Build for current platform
go build -o api.exe main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o api main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o api.exe main.go
```

### Run Binary

```bash
# Windows
.\api.exe

# Linux/Mac
./api
```

## 🔧 Development

### Adding a New Feature

Follow these steps to add a new resource (e.g., "Product"):

1. **Create Entity**: `internal/entity/product_entity.go`
2. **Create Repository**: `internal/repositories/product_repository.go`
3. **Create Service**: `internal/services/product_service.go`
4. **Create Handler**: `internal/handlers/product_handler.go`
5. **Create Request DTOs**: `internal/request/product_request.go`
6. **Create Routes**: `internal/route/product_route.go`
7. **Register Routes**: Update `internal/route/route.go`

See [FOLDER_STRUCTURE.md](FOLDER_STRUCTURE.md) for more details.

### Code Style

- Follow Go conventions and best practices
- Use `gofmt` for code formatting
- Use meaningful variable and function names
- Write comments for exported functions
- Keep functions small and focused

```bash
# Format code
go fmt ./...

# Run linter (if golangci-lint is installed)
golangci-lint run
```

## 📝 API Response Format

### Success Response

```json
{
  "status": "success",
  "message": "Operation successful",
  "data": {
    // Response data
  }
}
```

### Error Response

```json
{
  "status": "error",
  "message": "Error description",
  "error": {
    // Error details
  }
}
```

## 🔒 Security Best Practices

- ✅ Passwords are hashed using bcrypt
- ✅ JWT tokens for authentication
- ✅ Environment variables for sensitive data
- ✅ Input validation on all endpoints
- ✅ CORS configuration
- ✅ SQL injection prevention using prepared statements
- ✅ Request body size limits

## 🐛 Troubleshooting

### Database Connection Issues

```bash
# Check MySQL connection
mysql -h localhost -u root -p

# Verify database exists
SHOW DATABASES;
```

### Port Already in Use

Change `PORT_API` in `.env` file to a different port.

### Module Import Issues

```bash
# Clean module cache
go clean -modcache

# Re-download dependencies
go mod download
```

## 📚 Documentation

- [RULES.md](RULES.md) - Project guidelines and conventions
- [FOLDER_STRUCTURE.md](FOLDER_STRUCTURE.md) - Detailed folder structure

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👥 Authors

- Your Name - Initial work

## 🙏 Acknowledgments

- [Echo Framework](https://echo.labstack.com/)
- [Go Community](https://golang.org/)
- Clean Architecture by Robert C. Martin

## 📞 Support

For issues and questions:
- Create an issue in the repository
- Contact: your-email@example.com

---

**Happy Coding! 🚀**
