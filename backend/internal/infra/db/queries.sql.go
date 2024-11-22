// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package db

import (
	"context"
	"time"
)

const createAluno = `-- name: CreateAluno :exec
INSERT INTO Alunos (GRR, Nome) VALUES (?,?)
`

type CreateAlunoParams struct {
	Grr  string `json:"grr"`
	Nome string `json:"nome"`
}

func (q *Queries) CreateAluno(ctx context.Context, arg CreateAlunoParams) error {
	_, err := q.db.ExecContext(ctx, createAluno, arg.Grr, arg.Nome)
	return err
}

const createAula = `-- name: CreateAula :exec
INSERT INTO Aulas (ID, MateriaID, Professor, CargaHorariaMinutos, Data) VALUES (?,?,?,?,?)
`

type CreateAulaParams struct {
	ID                  string    `json:"id"`
	Materiaid           string    `json:"materiaid"`
	Professor           string    `json:"professor"`
	Cargahorariaminutos int32     `json:"cargahorariaminutos"`
	Data                time.Time `json:"data"`
}

func (q *Queries) CreateAula(ctx context.Context, arg CreateAulaParams) error {
	_, err := q.db.ExecContext(ctx, createAula,
		arg.ID,
		arg.Materiaid,
		arg.Professor,
		arg.Cargahorariaminutos,
		arg.Data,
	)
	return err
}

const createMateria = `-- name: CreateMateria :exec
INSERT INTO Materias (ID, Nome, CargaHorariaMinutos) VALUES (?,?,?)
`

type CreateMateriaParams struct {
	ID                  string `json:"id"`
	Nome                string `json:"nome"`
	Cargahorariaminutos int32  `json:"cargahorariaminutos"`
}

func (q *Queries) CreateMateria(ctx context.Context, arg CreateMateriaParams) error {
	_, err := q.db.ExecContext(ctx, createMateria, arg.ID, arg.Nome, arg.Cargahorariaminutos)
	return err
}

const createPresenca = `-- name: CreatePresenca :exec
INSERT INTO Presencas (AulaID, AlunoGRR) VALUES (?,?)
`

type CreatePresencaParams struct {
	Aulaid   string `json:"aulaid"`
	Alunogrr string `json:"alunogrr"`
}

func (q *Queries) CreatePresenca(ctx context.Context, arg CreatePresencaParams) error {
	_, err := q.db.ExecContext(ctx, createPresenca, arg.Aulaid, arg.Alunogrr)
	return err
}

const createProfessor = `-- name: CreateProfessor :exec
INSERT INTO Professores (CPF, Nome, Formacao, Telefone, Rua, Bairro, CEP, Cidade, Numero) VALUES (?,?,?,?,?,?,?,?,?)
`

type CreateProfessorParams struct {
	Cpf      string `json:"cpf"`
	Nome     string `json:"nome"`
	Formacao string `json:"formacao"`
	Telefone string `json:"telefone"`
	Rua      string `json:"rua"`
	Bairro   string `json:"bairro"`
	Cep      string `json:"cep"`
	Cidade   string `json:"cidade"`
	Numero   int32  `json:"numero"`
}

func (q *Queries) CreateProfessor(ctx context.Context, arg CreateProfessorParams) error {
	_, err := q.db.ExecContext(ctx, createProfessor,
		arg.Cpf,
		arg.Nome,
		arg.Formacao,
		arg.Telefone,
		arg.Rua,
		arg.Bairro,
		arg.Cep,
		arg.Cidade,
		arg.Numero,
	)
	return err
}

const deleteAluno = `-- name: DeleteAluno :exec
DELETE FROM Alunos WHERE GRR = ?
`

func (q *Queries) DeleteAluno(ctx context.Context, grr string) error {
	_, err := q.db.ExecContext(ctx, deleteAluno, grr)
	return err
}

const deleteAula = `-- name: DeleteAula :exec
DELETE FROM Aulas WHERE ID = ?
`

func (q *Queries) DeleteAula(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAula, id)
	return err
}

const deleteMateria = `-- name: DeleteMateria :exec
DELETE FROM Materias WHERE ID = ?
`

func (q *Queries) DeleteMateria(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteMateria, id)
	return err
}

const deleteProfessor = `-- name: DeleteProfessor :exec
DELETE FROM Professores WHERE CPF = ?
`

func (q *Queries) DeleteProfessor(ctx context.Context, cpf string) error {
	_, err := q.db.ExecContext(ctx, deleteProfessor, cpf)
	return err
}

const findAluno = `-- name: FindAluno :many
SELECT grr, nome FROM Alunos
WHERE (? = "" OR Nome LIKE ?) 
AND (? = "" OR GRR LIKE ?) 
LIMIT ? OFFSET ?
`

type FindAlunoParams struct {
	Nome   string `json:"nome"`
	Grr    string `json:"grr"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) FindAluno(ctx context.Context, arg FindAlunoParams) ([]Aluno, error) {
	rows, err := q.db.QueryContext(ctx, findAluno,
		arg.Nome,
		arg.Nome,
		arg.Grr,
		arg.Grr,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Aluno
	for rows.Next() {
		var i Aluno
		if err := rows.Scan(&i.Grr, &i.Nome); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAula = `-- name: FindAula :many
SELECT aulas.id, data, aulas.cargahorariaminutos, materiaid, professor, materias.id, nome, materias.cargahorariaminutos FROM Aulas
INNER JOIN Materias ON Materias.ID = MateriaID
WHERE (? = "" OR Materias.Nome LIKE ?) 
AND (? = "" OR Materias.ID LIKE ?) 
AND (? = "" OR Aulas.ID LIKE ?)
LIMIT ? OFFSET ?
`

type FindAulaParams struct {
	Materianome string `json:"materianome"`
	Materiaid   string `json:"materiaid"`
	Aulaid      string `json:"aulaid"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

type FindAulaRow struct {
	ID                    string    `json:"id"`
	Data                  time.Time `json:"data"`
	Cargahorariaminutos   int32     `json:"cargahorariaminutos"`
	Materiaid             string    `json:"materiaid"`
	Professor             string    `json:"professor"`
	ID_2                  string    `json:"id_2"`
	Nome                  string    `json:"nome"`
	Cargahorariaminutos_2 int32     `json:"cargahorariaminutos_2"`
}

func (q *Queries) FindAula(ctx context.Context, arg FindAulaParams) ([]FindAulaRow, error) {
	rows, err := q.db.QueryContext(ctx, findAula,
		arg.Materianome,
		arg.Materianome,
		arg.Materiaid,
		arg.Materiaid,
		arg.Aulaid,
		arg.Aulaid,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAulaRow
	for rows.Next() {
		var i FindAulaRow
		if err := rows.Scan(
			&i.ID,
			&i.Data,
			&i.Cargahorariaminutos,
			&i.Materiaid,
			&i.Professor,
			&i.ID_2,
			&i.Nome,
			&i.Cargahorariaminutos_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findLogin = `-- name: FindLogin :one
SELECT Senha FROM Login WHERE Username = ?
`

func (q *Queries) FindLogin(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, findLogin, username)
	var senha string
	err := row.Scan(&senha)
	return senha, err
}

const findMateria = `-- name: FindMateria :many
SELECT id, nome, cargahorariaminutos FROM Materias
WHERE (? = "" OR Nome LIKE ?) 
AND (? = "" OR ID LIKE ?) 
LIMIT ? OFFSET ?
`

type FindMateriaParams struct {
	Nome   string `json:"nome"`
	ID     string `json:"id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) FindMateria(ctx context.Context, arg FindMateriaParams) ([]Materia, error) {
	rows, err := q.db.QueryContext(ctx, findMateria,
		arg.Nome,
		arg.Nome,
		arg.ID,
		arg.ID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Materia
	for rows.Next() {
		var i Materia
		if err := rows.Scan(&i.ID, &i.Nome, &i.Cargahorariaminutos); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findProfessores = `-- name: FindProfessores :many
SELECT cpf, nome, formacao, telefone, rua, bairro, cidade, cep, numero FROM Professores
WHERE (? = "" OR CPF LIKE ?) 
AND (? = "" OR Nome LIKE ?)
LIMIT ? OFFSET ?
`

type FindProfessoresParams struct {
	Cpf    string `json:"cpf"`
	Nome   string `json:"nome"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) FindProfessores(ctx context.Context, arg FindProfessoresParams) ([]Professore, error) {
	rows, err := q.db.QueryContext(ctx, findProfessores,
		arg.Cpf,
		arg.Cpf,
		arg.Nome,
		arg.Nome,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Professore
	for rows.Next() {
		var i Professore
		if err := rows.Scan(
			&i.Cpf,
			&i.Nome,
			&i.Formacao,
			&i.Telefone,
			&i.Rua,
			&i.Bairro,
			&i.Cidade,
			&i.Cep,
			&i.Numero,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAluno = `-- name: UpdateAluno :exec
UPDATE Alunos SET Nome = ? WHERE GRR = ?
`

type UpdateAlunoParams struct {
	Nome string `json:"nome"`
	Grr  string `json:"grr"`
}

func (q *Queries) UpdateAluno(ctx context.Context, arg UpdateAlunoParams) error {
	_, err := q.db.ExecContext(ctx, updateAluno, arg.Nome, arg.Grr)
	return err
}

const updateAula = `-- name: UpdateAula :exec
UPDATE Aulas SET MateriaID = ?, Professor = ?, Data = ? WHERE ID = ?
`

type UpdateAulaParams struct {
	Materiaid string    `json:"materiaid"`
	Professor string    `json:"professor"`
	Data      time.Time `json:"data"`
	ID        string    `json:"id"`
}

func (q *Queries) UpdateAula(ctx context.Context, arg UpdateAulaParams) error {
	_, err := q.db.ExecContext(ctx, updateAula,
		arg.Materiaid,
		arg.Professor,
		arg.Data,
		arg.ID,
	)
	return err
}

const updateMateria = `-- name: UpdateMateria :exec
UPDATE Materias SET Nome = ?, CargaHorariaMinutos = ? WHERE ID = ?
`

type UpdateMateriaParams struct {
	Nome                string `json:"nome"`
	Cargahorariaminutos int32  `json:"cargahorariaminutos"`
	ID                  string `json:"id"`
}

func (q *Queries) UpdateMateria(ctx context.Context, arg UpdateMateriaParams) error {
	_, err := q.db.ExecContext(ctx, updateMateria, arg.Nome, arg.Cargahorariaminutos, arg.ID)
	return err
}

const updateProfessor = `-- name: UpdateProfessor :exec
UPDATE Professores SET Nome = ?, Formacao = ?, Telefone = ?, CEP = ?, Rua = ?, Bairro = ?, Cidade = ?, Numero = ? WHERE CPF = ?
`

type UpdateProfessorParams struct {
	Nome     string `json:"nome"`
	Formacao string `json:"formacao"`
	Telefone string `json:"telefone"`
	Cep      string `json:"cep"`
	Rua      string `json:"rua"`
	Bairro   string `json:"bairro"`
	Cidade   string `json:"cidade"`
	Numero   int32  `json:"numero"`
	Cpf      string `json:"cpf"`
}

func (q *Queries) UpdateProfessor(ctx context.Context, arg UpdateProfessorParams) error {
	_, err := q.db.ExecContext(ctx, updateProfessor,
		arg.Nome,
		arg.Formacao,
		arg.Telefone,
		arg.Cep,
		arg.Rua,
		arg.Bairro,
		arg.Cidade,
		arg.Numero,
		arg.Cpf,
	)
	return err
}
