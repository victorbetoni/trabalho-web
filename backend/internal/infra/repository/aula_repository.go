package repository

import (
	"context"
	"database/sql"
	"regexp"
	"victorbetoni/trabalho-web/internal/domain/entity"
	"victorbetoni/trabalho-web/internal/infra/db"

	"github.com/victorbetoni/go-streams/streams"
)

type AulaRepository struct {
	dbConn *sql.DB
	*db.Queries
}

func NewAulaRepository(dbConn *sql.DB) *AulaRepository {
	return &AulaRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (r *AulaRepository) Find(ctx context.Context, filter entity.AulaFilter) ([]*entity.Aula, error) {
	filter.Professor = regexp.MustCompile(`\D`).ReplaceAllString(filter.Professor, "")
	p, err := r.Queries.FindAula(ctx, db.FindAulaParams{
		Professor: filter.Professor,
		Limit:     int32(filter.Limit),
		Offset:    int32(filter.Page-1) * int32(filter.Limit),
	})

	if err != nil {
		return nil, err
	}

	if len(p) == 0 {
		return nil, sql.ErrNoRows
	}

	return *streams.Map(streams.StreamOf(p...), func(p db.FindAulaRow) *entity.Aula {

		return &entity.Aula{
			ID:           p.ID,
			Titulo:       p.Titulo,
			Data:         p.Data,
			CargaHoraria: int(p.Cargahorariaminutos),
			Professor: &entity.Professor{
				CPF:      p.Cpf,
				Nome:     p.Nome,
				Formacao: p.Formacao,
				Telefone: p.Telefone,
				Endereco: entity.Endereco{
					CEP:    p.Cep,
					Rua:    p.Rua,
					Bairro: p.Bairro,
					Cidade: p.Cidade,
					Numero: int(p.Numero),
				},
			},
		}
	}).ToSlice(), nil
}

func (r *AulaRepository) Create(ctx context.Context, aula *entity.Aula) error {
	return r.Queries.CreateAula(ctx, db.CreateAulaParams{
		ID:                  aula.ID,
		Data:                aula.Data,
		Professor:           aula.Professor.CPF,
		Titulo:              aula.Titulo,
		Cargahorariaminutos: int32(aula.CargaHoraria),
	})
}

func (r *AulaRepository) Delete(ctx context.Context, id string) error {
	return r.Queries.DeleteAula(ctx, id)
}
