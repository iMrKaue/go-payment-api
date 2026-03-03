🚀 Go Payment API

API REST desenvolvida em Go com autenticação JWT, arquitetura em camadas, testes unitários com mocks e integração com PostgreSQL via Docker.

REST API built with Go featuring JWT authentication, layered architecture, unit tests with mocks, and PostgreSQL integration via Docker.

📌 Sobre o Projeto | About the Project

Este projeto foi desenvolvido com foco em boas práticas de backend, incluindo:

Arquitetura em camadas (Handler → Service → Repository)

Inversão de dependência com interfaces

Testes unitários isolados com mocks

Autenticação com JWT

Hash de senha com bcrypt

Banco de dados PostgreSQL via Docker

Middleware para proteção de rotas

This project was built focusing on backend best practices, including:

Layered architecture (Handler → Service → Repository)

Dependency inversion using interfaces

Isolated unit tests with mocks

JWT authentication

Password hashing with bcrypt

PostgreSQL database via Docker

Middleware for route protection

🏗 Arquitetura | Architecture
cmd/
internal/
 ├── handler/
 ├── service/
 ├── repository/
 ├── middleware/
 ├── model/
 ├── dto/
 ├── utils/
configs/

Fluxo de requisição:

Request → Handler → Service → Repository → Database

Request → Handler → Service → Repository → Database

🔐 Autenticação | Authentication

Registro de usuário com senha criptografada (bcrypt)

Login validando hash de senha

Geração de JWT com expiração

Middleware validando token

Rota protegida /profile

User registration with hashed password (bcrypt)

Login with password validation

JWT generation with expiration

Token validation middleware

Protected route /profile

🧪 Testes | Tests

Testes unitários aplicados na camada de Service

Mock manual do repository

Cobertura de testes: 85%+ na camada de negócio

Unit tests applied to the Service layer

Manual repository mock

Test coverage: 85%+ on business logic layer

Para rodar testes:

To run tests:

go test -cover ./...
🐳 Como Rodar o Projeto | Running the Project
1️⃣ Subir o banco com Docker

Start database with Docker:

docker compose up -d
2️⃣ Rodar a aplicação

Run the application:

go run cmd/main.go

Servidor disponível em:

Server available at:

http://localhost:8080
📮 Endpoints | Endpoints
Criar Usuário | Create User

POST /users

{
  "name": "User Name",
  "email": "user@email.com",
  "password": "123456"
}
Login

POST /login

{
  "email": "user@email.com",
  "password": "123456"
}

Retorna:

Returns:

{
  "token": "JWT_TOKEN"
}
Rota Protegida | Protected Route

GET /profile

Header:

Authorization: Bearer JWT_TOKEN
🛠 Tecnologias | Technologies

Go (Golang)

PostgreSQL

Docker

JWT

Bcrypt

net/http

Testing (Go testing package)

🎯 Objetivo

Este projeto demonstra:

Estruturação profissional de backend

Separação de responsabilidades

Segurança básica aplicada

Testabilidade

Boas práticas em Go

This project demonstrates:

Professional backend structuring

Separation of concerns

Applied security practices

Testability

Go best practices

👨‍💻 Autor | Author

Kaue Ferreira Macedo
LinkedIn: www.linkedin.com/in/kaue-macedo

GitHub: https://github.com/iMrKaue