package models

import "github.com/MatheusMoronari/Desafio/banco"

func Insert(todo Todo) (id int64, err error) {

	conn, err := banco.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close() //isso aqui ele executa depois que a função for executada.

	sql := `insert into todos (title, description, done ) VALUES ($1,$2,$3) RETURNING id`
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}
