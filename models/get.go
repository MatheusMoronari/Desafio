package models

import "github.com/MatheusMoronari/Desafio/banco"

func Get(id int64) (todo Todo, err error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)
	err = row.Scan(&todo.id, &todo.Title, &todo.Description, &todo.Done)
	return
}
