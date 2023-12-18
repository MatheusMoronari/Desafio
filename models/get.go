package models

import (
	"github.com/MatheusMoronari/Desafio/banco"
	"github.com/MatheusMoronari/Desafio/entidades"
)

func Get(id int64) (pessoa entidades.Pessoa, err error) {
	conn, err := banco.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	row := conn.QueryRow(`SELECT p.id,p.nome,p.codigo,tp.descricao FROM pessoas as p inner join tipo_pessoa tp on(p.tipo_pessoa = tp.id_tipo_pessoa) WHERE id=$1`, id)
	err = row.Scan(&pessoa.Id, &pessoa.Nome, &pessoa.Codigo, &pessoa.Tipo_pessoa_desc)
	return
}
