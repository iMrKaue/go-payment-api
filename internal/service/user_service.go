package service

import (
	"fmt"
	"strings"

	"github.com/iMrKaue/go-payment-api/internal/model"
	"github.com/iMrKaue/go-payment-api/internal/repository"
	"github.com/iMrKaue/go-payment-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(name, email, password string) (*model.User, error) {

	//Hash da senha antes de salvar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao criptografar senha")
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.Repo.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, fmt.Errorf("email já cadastrado")
		}
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("usuário não encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("senha inválida")
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token")
	}

	return token, nil
}
