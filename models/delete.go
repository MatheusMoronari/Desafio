package models

import "github.com/MatheusMoronari/Desafio/banco"

func Delete(id int64) (int64, error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.Exec(`DEKETE FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
