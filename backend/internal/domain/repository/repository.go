package repository

import (
	"context"
	"victorbetoni/trabalho-web/internal/domain/entity"
)

type RepositoryInterface interface{}

type ProfessorRepositoryInterface interface {
	Find(ctx context.Context, filter entity.ProfessorFilter) ([]*entity.Professor, error)
	FindOne(ctx context.Context, filter entity.ProfessorFilter) (*entity.Professor, error)
	Create(ctx context.Context, professor *entity.Professor) error
	Update(ctx context.Context, professor *entity.Professor) error
	Delete(ctx context.Context, cpf string) error
	RepositoryInterface
}

type AulaRepositoryInterface interface {
	Find(ctx context.Context, filter entity.AulaFilter) ([]*entity.Aula, error)
	Create(ctx context.Context, aula *entity.Aula) error
	Delete(ctx context.Context, id string) error
	RepositoryInterface
}

type LoginRepositoryInterface interface {
	FindLogin(ctx context.Context, username string) (string, error)
}
