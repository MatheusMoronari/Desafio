package entidades

type Pessoa struct {
	id               int64  `json: "id"`
	Nome             string `json: "nome"`
	Codigo           string `json:"codigo"`
	Tipo_pessoa      int    `json: "tipo_pessoa"`
	Tipo_pessoa_desc string `json: "tipo_pessoa_desc"`
}
