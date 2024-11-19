package repository

import (
	"context"
	"database/sql"
	"victorbetoni/trabalho-web/internal/domain/entity"
	"victorbetoni/trabalho-web/internal/infra/db"

	"github.com/victorbetoni/go-streams/streams"
)

type ProfessorRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewProfessorRepository(dbConn *sql.DB) *ProfessorRepository {
	return &ProfessorRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *ProfessorRepository) Find(ctx context.Context, filter entity.ProfessorFilter) ([]*entity.Professor, error) {
	p, err := r.Queries.FindProfessores(ctx, db.FindProfessoresParams{
		Cpf:    filter.CPF,
		Nome:   filter.Nome,
		Limit:  int32(filter.Limit),
		Offset: int32(filter.Page-1) * int32(filter.Limit),
	})

	if err != nil {
		return nil, err
	}

	if len(p) == 0 {
		return nil, sql.ErrNoRows
	}

	return *streams.Map(streams.StreamOf(p...), func(p db.Professore) *entity.Professor {
		return &entity.Professor{
			CPF:  p.Cpf,
			Nome: p.Nome,
		}
	}).ToSlice(), nil
}

func (r *ProfessorRepository) FindOne(ctx context.Context, filter entity.ProfessorFilter) (*entity.Professor, error) {
	filter.Limit = 1
	filter.Page = 1
	p, err := r.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return p[0], nil
}

func (r *ProfessorRepository) Create(ctx context.Context, professor *entity.Professor) error {
	return r.Queries.CreateProfessor(ctx, db.CreateProfessorParams{
		Cpf:      professor.CPF,
		Nome:     professor.Nome,
		Formacao: professor.Formacao,
		Telefone: professor.Telefone,
		Rua:      professor.Endereco.Rua,
		Bairro:   professor.Endereco.Bairro,
		Cep:      professor.Endereco.CEP,
		Cidade:   professor.Endereco.Cidade,
		Numero:   int32(professor.Endereco.Numero),
	})
}

func (r *ProfessorRepository) Update(ctx context.Context, professor *entity.Professor) error {
	return r.Queries.UpdateProfessor(ctx, db.UpdateProfessorParams{
		Nome: professor.Nome,
		Cpf:  professor.CPF,
	})
}

func (r *ProfessorRepository) Delete(ctx context.Context, cpf string) error {
	return r.Queries.DeleteProfessor(ctx, cpf)
}
