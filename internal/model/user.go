package model

import "time"

// User representa a tabela users no banco
type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
