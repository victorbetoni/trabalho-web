package entity

type Professor struct {
	CPF      string   `json:"id"`
	Nome     string   `json:"nome"`
	Formacao string   `json:"formacao"`
	Telefone string   `json:"telefone"`
	Endereco Endereco `json:"endereco"`
}
