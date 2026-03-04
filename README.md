# 🚀 Go Payment API

REST API built with Go featuring JWT authentication, layered architecture, unit testing with mocks, and PostgreSQL integration via Docker.

API REST desenvolvida em Go com autenticação JWT, arquitetura em camadas, testes unitários com mocks e integração com PostgreSQL via Docker.

## 📌 Overview | Visão Geral
This project demonstrates backend best practices including:

- Layered architecture (Handler → Service → Repository)
- Dependency inversion using interfaces
- JWT authentication
- Password hashing with bcrypt
- Middleware for protected routes
- Unit tests with 85%+ coverage on business logic
- PostgreSQL running via Docker

Este projeto demonstra boas práticas de backend incluindo:

- Arquitetura em camadas (Handler → Service → Repository)
- Inversão de dependência com interfaces
- Autenticação JWT
- Criptografia de senha com bcrypt
- Middleware para rotas protegidas
- Testes unitários com 85%+ de cobertura na camada de negócio
- PostgreSQL via Docker

## 🏗 Architecture | Arquitetura
Request Flow:
Request → Handler → Service → Repository → Database

Project Structure:
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

## 🔐 Authentication Flow | Fluxo de Autenticação
1. User registration with hashed password (bcrypt)
2. Login validates password hash
3. JWT token is generated with expiration
4. Middleware validates token
5. Protected route `/profile` requires valid token

## 🧪 Tests | Testes
Unit tests implemented in the Service layer using manual mocks.

Coverage:
85%+ of business logic statements covered

Run tests:
go test -cover ./...
🐳 Running the Project | Executando o Projeto
1️⃣ Start PostgreSQL with Docker
docker compose up -d
2️⃣ Run the application
go run cmd/main.go

Server runs at:
http://localhost:8080
📮 Endpoints
➤ Create User

POST /users

{
  "name": "User Name",
  "email": "user@email.com",
  "password": "123456"
}
➤ Login

POST /login

{
  "email": "user@email.com",
  "password": "123456"
}

Response:

{
  "token": "JWT_TOKEN"
}
➤ Protected Route
GET /profile
Header:
Authorization: Bearer JWT_TOKEN

🛠 Technologies
Go (Golang)
PostgreSQL
Docker
JWT
Bcrypt
net/http
Go testing package

🎯 Purpose
This project highlights backend fundamentals, clean structure, authentication mechanisms, testability, and production-oriented practices.

Este projeto evidencia fundamentos sólidos de backend, estrutura organizada, autenticação segura e testabilidade.

👨‍💻 Author
Kaue Ferreira Macedo

LinkedIn: https://www.linkedin.com/in/kaue-macedo
GitHub: https://github.com/iMrKaue


