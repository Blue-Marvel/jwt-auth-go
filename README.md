# JWT Authentication Tutorial with Go

This tutorial demonstrates how to implement JSON Web Token (JWT) authentication in a Go application. The implementation secures APIs by requiring a valid JWT for access.

---

## Features

- User Registration
- User Login with JWT Token Generation
- Protected Routes (accessible only with a valid JWT)

---

## Prerequisites

- Go installed on your system
- Basic knowledge of Go and REST APIs
- MongoDB (for user data storage)

---

## Project Structure

```
jwt-auth/
├── controllers/
│   ├── user_controller.go
|── database_connection
|   ├── database_connection.go
├── helpers
|   ├── auth_helper.go
|   ├── token_helper.go
├── middleware/
│   └── auth_middleware.go
├── models/
│   └── user_model.go
├── routers/
│   ├── auth_router.go
|   ├── user_router.go
├── main.go
├── go.mod
├── go.sum
```

---

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/jwt-auth.git
   cd jwt-auth
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Environment Configuration**
   Create a `.env` file with the following content:

   ```
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your_jwt_secret
   ```

4. **Run the Application**
   ```bash
   go run main.go
   ```

---

## Endpoints

### 1. **User Registration**

- **URL**: `/api/register`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "message": "User registered successfully!"
  }
  ```

### 2. **User Login**

- **URL**: `/api/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "token": "your.jwt.token"
  }
  ```

### 3. **Protected Route**

- **URL**: `/api/protected`
- **Method**: `GET`
- **Headers**:
  ```
  Authorization: Bearer your.jwt.token
  ```
- **Response**:
  ```json
  {
    "message": "Access granted!"
  }
  ```

---

## Key Files

### 1. **`jwt_helper.go`**

- Handles JWT generation and validation.

### 2. **`auth_middleware.go`**

- Middleware to protect routes by verifying JWTs.

### 3. **`auth_controller.go`**

- Manages user registration and login.

### 4. **`user_controller.go`**

- Handles user-related operations (e.g., fetching user data).

---

## Running Tests

Write unit tests to validate the JWT functionality and other application features:

```bash
go test ./...
```

---

## Further Improvements

- Add password hashing (e.g., using `bcrypt`).
- Implement role-based access control (RBAC).
- Use a config management library for better environment variable handling.

---

This README provides a starting point for building a secure Go application with JWT authentication. Feel free to extend or modify the implementation based on your project needs.
