package main

import (
	"fmt"
	"net/http"

	"github.com/iMrKaue/go-payment-api/configs"
	"github.com/iMrKaue/go-payment-api/internal/handler"
	"github.com/iMrKaue/go-payment-api/internal/middleware"
	"github.com/iMrKaue/go-payment-api/internal/model"
	"github.com/iMrKaue/go-payment-api/internal/repository"
	"github.com/iMrKaue/go-payment-api/internal/service"
)

func main() {
	db := configs.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/login", userHandler.Login)
	http.HandleFunc("/profile", middleware.AuthMiddleware(userHandler.Profile))

	fmt.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)

	newUser := &model.User{
		Name:     "Kaue 2",
		Email:    "kaue2@gmail.com",
		Password: "123456",
	}

	newUser, err := userService.CreateUser(
		newUser.Name,
		newUser.Email,
		newUser.Password,
	)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Usuário criado com ID:", newUser.ID)
}
