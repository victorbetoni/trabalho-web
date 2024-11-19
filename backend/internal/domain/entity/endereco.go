package entity

type Endereco struct {
	CEP    string `json:"cep"`
	Rua    string `json:"rua"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	Numero string `json:"numero"`
}
