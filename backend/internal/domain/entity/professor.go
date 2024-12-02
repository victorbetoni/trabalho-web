package entity

type Professor struct {
	CPF        string   `json:"cpf"`
	Nome       string   `json:"nome"`
	Formacao   string   `json:"formacao"`
	Telefone   string   `json:"telefone"`
	Endereco   Endereco `json:"endereco"`
	AulasDadas int      `json:"aulas_dadas"`
}
