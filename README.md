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

- **URL**: `/api/auth/signup`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "first_name": "",
    "last_name": "",
    "password": "",
    "email": "",
    "phone": "",
    "user_type": "USER"|"ADMIN"
  }
  ```
- **Response**:
  ```json
  {
    "InsertedID": "123245"
  }
  ```

### 2. **User Login**

- **URL**: `/api/auth/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

### 3. \*_All Users_

- **URL**: `/api/users`
- **Method**: `GET`
- **Headers**:
  ```
  Authorization: Bearer your.jwt.token
  ```

### 4. \*_Users by Id_

- **URL**: `/api/users/user_id`
- **Method**: `GET`
- **Headers**:
  ```
  Authorization: Bearer your.jwt.token
  ```

---

## Key Files

### 1. **`token_helper.go`**

### 2. **`auth_helper.go`**

- Handles JWT generation and validation.

### 3. **`auth_middleware.go`**

- Middleware to protect routes by verifying JWTs.

### 4. **`user_controller.go`**

- Handles user-related operations (e.g.,user auth and fetching user data).

---

<!--
## Running Tests

Write unit tests to validate the JWT functionality and other application features:

```bash
go test ./...
```

--- -->
<!--
## Further Improvements

- Add password hashing (e.g., using `bcrypt`).
- Implement role-based access control (RBAC).
- Use a config management library for better environment variable handling.

---

This README provides a starting point for building a secure Go application with JWT authentication. Feel free to extend or modify the implementation based on your project needs. -->
