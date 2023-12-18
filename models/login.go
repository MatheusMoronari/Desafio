package models

import (
	"database/sql"
	"fmt"

	"github.com/MatheusMoronari/Desafio/banco"
	_ "github.com/lib/pq"
)

// func ValidaUsuario(username, password string) bool {
// 	conn, err := banco.OpenConnection()
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()

// 	rows, error := conn.Query("SELECT id FROM users WHERE username = $1 AND password = $2", username, password)
// 	fmt.Println(rows)
// 	if error == sql.ErrNoRows {
// 		return false

// 	} else if error != nil {
// 		log.Fatal(error)
// 	}
// 	return true
// }

func ValidaUsuario(dadoslogin Dadoslogin) bool {
	conn, err := banco.OpenConnection()
	if err != nil {
		return false
	}
	defer conn.Close()
	fmt.Println("Username:", dadoslogin.Username)
	fmt.Println("Password:", dadoslogin.Password)
	var userId int
	err = conn.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", dadoslogin.Username, dadoslogin.Password).Scan(&userId)
	println(err)
	if err == sql.ErrNoRows {
		fmt.Println("Usuário não encontrado ou senha incorreta")
		return false
	} else if err != nil {
		fmt.Printf("Erro ao realizar a consulta: %s\n", err)
		return false
	}
	return true // Usuário encontrado e senha corresponde
}
