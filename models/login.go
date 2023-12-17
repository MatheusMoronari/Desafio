package models

import (
	"database/sql"
	"log"

	"github.com/MatheusMoronari/Desafio/banco"
	_ "github.com/lib/pq"
)

func ValidateUser(username, password string) bool {
	conn, err := banco.OpenConnection()
	if err != nil {
		return false
	}
	defer conn.Close()

	var userId int
	err = conn.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", username, password).Scan(&userId)

	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Fatal(err)
	}
	return true
}
