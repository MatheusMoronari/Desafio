package models

import (
	"github.com/MatheusMoronari/Desafio/banco"
	"github.com/MatheusMoronari/Desafio/entidades"
)

func Insert(pessoa entidades.Pessoa) (id int64, err error) {

	conn, err := banco.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close() //isso aqui ele executa depois que a função for executada.

	sql := `insert into pessoas (nome, codigo, tipo_pessoa, data_inclusao ) VALUES ($1,$2,$3,current_timestamp) RETURNING id`
	err = conn.QueryRow(sql, pessoa.Nome, pessoa.Codigo, pessoa.Tipo_pessoa).Scan(&id)
	return
}
