package models

type pessoa struct {
	id          int64  `json: "id"`
	Nome        string `json: "nome"`
	Codigo      string `json:"codigo"`
	Tipo_pessoa int    `json: "tipo_pessoa"`
}
