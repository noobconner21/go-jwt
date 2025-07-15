# Go JWT Authentication API

A simple yet robust REST API built with Go (Gin framework) that provides JWT-based authentication with PostgreSQL database integration.

## 🚀 Features

- **User Registration** - Secure user signup with email and password
- **User Authentication** - JWT-based login system
- **Password Security** - Bcrypt hashing for password storage
- **Token Validation** - Middleware for protecting routes
- **Cookie-based JWT** - Secure token storage in HTTP-only cookies
- **PostgreSQL Integration** - GORM ORM with PostgreSQL database
- **Environment Configuration** - Secure configuration management

## 🛠️ Tech Stack

- **Language**: Go 1.24.5
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment Management**: godotenv
- **Development Tools**: CompileDaemon (hot reload)

## 📁 Project Structure

```
go-jwt/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── go.sum                  # Go module dependencies
├── controllers/
│   └── usersController.go  # User-related HTTP handlers
├── initializers/
│   ├── connectToDb.go      # Database connection setup
│   ├── loadEnvVariables.go # Environment variables loader
│   └── syncDb.go           # Database synchronization
├── middleware/
│   └── requireAuth.go      # JWT authentication middleware
└── models/
    └── userModel.go        # User data model
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.24.5 or higher
- PostgreSQL database
- Git
- CompileDaemon (optional, for development with hot reload)
  ```bash
  go install github.com/githubnemo/CompileDaemon@latest
  ```

### 1. Clone the Repository

```bash
git clone https://github.com/noobconner21/go-jwt.git
cd go-jwt
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
DB=postgres://username:password@localhost:5432/database_name?sslmode=disable
SECRET=your-jwt-secret-key-here
```

**Environment Variables:**

- `DB`: PostgreSQL connection string
- `SECRET`: JWT signing secret (use a strong, random string)

### 4. Database Setup

Ensure PostgreSQL is running and create your database:

```sql
CREATE DATABASE your_database_name;
```

### 5. Run the Application

#### Production Mode

```bash
go run main.go
```

#### Development Mode (with hot reload)

For development with automatic recompilation on file changes:

```bash
# First, build the application
go build -o go-jwt

# Then run with CompileDaemon for hot reload
CompileDaemon --command="./go-jwt"
```

The server will start on `http://localhost:8080`

## 📚 API Endpoints

### Authentication Endpoints

#### Register User

```http
POST /signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**

- `200 OK` - User created successfully
- `400 Bad Request` - Invalid request or user already exists

#### Login User

```http
POST /login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response:**

- `200 OK` - Login successful (JWT token set as HTTP-only cookie)
- `400 Bad Request` - Invalid credentials

#### Validate Token

```http
GET /validate
Cookie: Authorization=jwt-token-here
```

**Response:**

- `200 OK` - Token valid, returns user information
- `401 Unauthorized` - Invalid or expired token

## 🔐 Security Features

- **Password Hashing**: Passwords are hashed using bcrypt with salt rounds
- **JWT Security**: Tokens are signed with HMAC SHA256
- **HTTP-Only Cookies**: JWT tokens stored securely in HTTP-only cookies
- **Token Expiration**: Tokens expire after 30 days
- **Input Validation**: Request body validation and sanitization

## 🏗️ Architecture

### Models

- **User Model**: Defines user structure with email and password fields
- **GORM Integration**: Automatic table creation and migration

### Controllers

- **Signup**: Handles user registration with password hashing
- **Login**: Authenticates users and generates JWT tokens
- **Validate**: Protected endpoint for token validation

### Middleware

- **RequireAuth**: Validates JWT tokens and protects routes
- **User Context**: Injects authenticated user into request context

### Initializers

- **Database Connection**: PostgreSQL connection with GORM
- **Environment Loading**: Secure environment variable management
- **Database Sync**: Automatic table creation and migration

## 🚀 Usage Examples

### Register a new user

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"mypassword"}'
```

### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"mypassword"}' \
  -c cookies.txt
```

### Access protected route

```bash
curl -X GET http://localhost:8080/validate \
  -b cookies.txt
```

## 📝 License

This project is open source and available under the [MIT License](LICENSE).


---

**Built with ❤️ in Go with ShayC**
