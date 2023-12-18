package models

import (
	"github.com/MatheusMoronari/Desafio/banco"
	"github.com/MatheusMoronari/Desafio/entidades"
)

func Update(id int64, pessoa entidades.Pessoa) (int64, error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.Exec(`UPDATE pessoas SET nome = $1 WHERE id=$2`,
		pessoa.Nome, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
