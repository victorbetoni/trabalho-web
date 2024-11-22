// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Aluno struct {
	Grr  string `json:"grr"`
	Nome string `json:"nome"`
}

type Aula struct {
	ID                  string    `json:"id"`
	Data                time.Time `json:"data"`
	Cargahorariaminutos int32     `json:"cargahorariaminutos"`
	Materiaid           string    `json:"materiaid"`
	Professor           string    `json:"professor"`
}

type Login struct {
	Username string `json:"username"`
	Senha    string `json:"senha"`
}

type Materia struct {
	ID                  string `json:"id"`
	Nome                string `json:"nome"`
	Cargahorariaminutos int32  `json:"cargahorariaminutos"`
}

type Presenca struct {
	Aulaid   string `json:"aulaid"`
	Alunogrr string `json:"alunogrr"`
	Horas    int32  `json:"horas"`
}

type Professore struct {
	Cpf      string `json:"cpf"`
	Nome     string `json:"nome"`
	Formacao string `json:"formacao"`
	Telefone string `json:"telefone"`
	Rua      string `json:"rua"`
	Bairro   string `json:"bairro"`
	Cidade   string `json:"cidade"`
	Cep      string `json:"cep"`
	Numero   int32  `json:"numero"`
}
