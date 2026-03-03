package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "postgres://postgres:postgres@127.0.0.1:5433/paymentdb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco")
	}

	fmt.Println("Banco conectado com sucesso!")

	return db
}
