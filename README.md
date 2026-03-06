# 🚀 Go Payment API

REST API built with **Go (Golang)** featuring JWT authentication, layered architecture, unit testing with mocks, and PostgreSQL integration via Docker.

---

## 📌 Overview

This project demonstrates backend engineering best practices:

- Layered architecture (**Handler → Service → Repository**)
- Dependency inversion using interfaces
- JWT authentication
- Password hashing with bcrypt
- Middleware for protected routes
- Unit testing with mocks
- PostgreSQL running via Docker
- Clean and maintainable project structure

The goal of this project is to demonstrate backend fundamentals, authentication flows, testability and production-oriented practices.

---

## 🏗 Architecture

### Request Flow


Request → Handler → Service → Repository → Database


### Project Structure


cmd/
configs/
internal/
├── handler/
├── service/
├── repository/
├── middleware/
├── model/
├── dto/
└── utils/


**Handler**
- Receives HTTP requests
- Validates input
- Returns responses

**Service**
- Contains business logic
- Handles authentication logic
- Coordinates repositories

**Repository**
- Responsible for database operations
- Abstracted via interfaces

---

## 🔐 Authentication Flow

1. User registers with hashed password (bcrypt)
2. Login validates password hash
3. API generates a JWT token
4. Token includes expiration
5. Middleware validates JWT
6. Protected routes require valid token

Example protected route:


GET /profile


Required header:


Authorization: Bearer JWT_TOKEN


---

## 🧪 Tests

Unit tests are implemented in the **Service layer** using manual mocks.

Coverage:


85%+ coverage of business logic


Run tests:


go test -cover ./...


---

## 🐳 Running the Project

### 1. Start PostgreSQL with Docker


docker compose up -d


### 2. Run the application


go run cmd/main.go


Server runs at:


http://localhost:8080


---

## 📮 Endpoints

### Create User


POST /users


Request body:

```json
{
  "name": "User Name",
  "email": "user@email.com",
  "password": "123456"
}
Login
POST /login

Request body:

{
  "email": "user@email.com",
  "password": "123456"
}

Response:

{
  "token": "JWT_TOKEN"
}
Protected Route
GET /profile
```
---
### Header required:

Authorization: Bearer JWT_TOKEN
🛠 Technologies

- Go (Golang)
- PostgreSQL
- Docker
- JWT
- Bcrypt
- net/http
- Go testing package

---
### 🎯 Purpose

This project highlights:

- Backend architecture design
- Authentication with JWT
- Secure password handling
- Dependency inversion
- Testable business logic
- Containerized database environment

---
### 👨‍💻 Author

Kaue Ferreira Macedo

LinkedIn:
https://www.linkedin.com/in/kaue-macedo

GitHub:
https://github.com/iMrKaue
