package models

import "github.com/MatheusMoronari/Desafio/banco"

func Update(id int64, pessoa pessoa) (int64, error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.Exec(`UPDATE pessoas SET nome = $1, codigo= $2, tipo_pessoa= $3 WHERE id=$4`,
		pessoa.Nome, pessoa.Codigo, pessoa.Tipo_pessoa, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
