package service

import (
	"errors"
	"testing"

	"github.com/iMrKaue/go-payment-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Mock do repository
type MockUserRepository struct {
	CreateFunc      func(user *model.User) error
	FindByEmailFunc func(email string) (*model.User, error)
}

func (m *MockUserRepository) Create(user *model.User) error {
	return m.CreateFunc(user)
}

func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	return m.FindByEmailFunc(email)
}

func TestCreateUserSuccess(t *testing.T) {

	mockRepo := &MockUserRepository{
		CreateFunc: func(user *model.User) error {
			user.ID = 1
			return nil
		},
	}

	service := NewUserService(mockRepo)

	user, err := service.CreateUser("Kaue", "test@gmail.com", "123456")

	if err != nil {
		t.Errorf("esperava sucesso, mas deu erro: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("esperava ID 1, recebeu %d", user.ID)
	}
}

func TestCreateUserDuplicate(t *testing.T) {

	mockRepo := &MockUserRepository{
		CreateFunc: func(user *model.User) error {
			return errors.New("duplicate key")
		},
	}

	service := NewUserService(mockRepo)

	_, err := service.CreateUser("Kaue", "test@gmail.com", "123456")

	if err == nil {
		t.Errorf("esperava erro de duplicidade")
	}
}

func TestLoginSuccess(t *testing.T) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	mockRepo := &MockUserRepository{
		FindByEmailFunc: func(email string) (*model.User, error) {
			return &model.User{
				ID:       1,
				Email:    email,
				Password: string(hashedPassword),
			}, nil
		},
	}

	service := NewUserService(mockRepo)

	token, err := service.Login("test@email.com", "123456")

	if err != nil {
		t.Errorf("esperava sucesso, mas deu erro: %v", err)
	}

	if token == "" {
		t.Errorf("esperava token válido")
	}
}

func TestLoginUserNotFound(t *testing.T) {
	mockRepo := &MockUserRepository{
		FindByEmailFunc: func(email string) (*model.User, error) {
			return nil, errors.New("not found")
		},
	}

	service := NewUserService(mockRepo)

	_, err := service.Login("test@email.com", "123456")

	if err == nil {
		t.Errorf("esperava erro para usuário inexistente")
	}
}

func TestLoginInvalidPassword(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	mockRepo := &MockUserRepository{
		FindByEmailFunc: func(email string) (*model.User, error) {
			return &model.User{
				ID:       1,
				Email:    email,
				Password: string(hashedPassword),
			}, nil
		},
	}

	service := NewUserService(mockRepo)

	_, err := service.Login("test@gmail.com", "senha_errada")

	if err == nil {
		t.Errorf("esperava erro de senha inválida")
	}
}
