package models

import "github.com/MatheusMoronari/Desafio/banco"

func Get(id int64) (pessoa pessoa, err error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	row := conn.QueryRow(`SELECT * FROM pessoas WHERE id=$1`, id)
	err = row.Scan(&pessoa.id, &pessoa.Nome, &pessoa.Codigo, &pessoa.Tipo_pessoa)
	return
}
