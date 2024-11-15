package entity

type ProfessorFilter struct {
	CPF   string `json:"cpf"`
	Nome  string `json:"nome"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
