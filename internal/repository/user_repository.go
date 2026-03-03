package repository

import (
	"database/sql"

	"github.com/iMrKaue/go-payment-api/internal/model"
)

type UserRepositoryInterface interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

// UserRepository é responsável por acessar o banco
type UserRepository struct {
	DB *sql.DB
}

// Construtor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create insere um novo usuário no banco
func (r *UserRepository) Create(user *model.User) error {
	query := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`

	return r.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)

}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := "SELECT id, name, email, password, created_at FROM users WHERE email = $1"

	user := &model.User{}

	err := r.DB.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
