package models

import "github.com/MatheusMoronari/Desafio/banco"

func GetAll() (pessoas []pessoa, err error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	rows, err := conn.Query(`SELECT * FROM pessoas`)
	if err != nil {
		return
	}
	for rows.Next() {
		var pessoa pessoa

		err = rows.Scan(&pessoa.id, &pessoa.Nome, &pessoa.Codigo, &pessoa.Tipo_pessoa)
		if err != nil {
			continue
		}
		pessoas = append(pessoas, pessoa)
	}
	return
}
