package repository

import (
	"context"
	"errors"
	"victorbetoni/trabalho-web/internal/domain/repository"
	"victorbetoni/trabalho-web/internal/infra/db"
	"victorbetoni/trabalho-web/pkg/uow"
)

var ErrQueriesNotSet = errors.New("queries not set")

type Repository struct {
	*db.Queries
}

func (r *Repository) SetQuery(q *db.Queries) {
	r.Queries = q
}

func (r *Repository) Validate() error {
	if r.Queries == nil {
		return ErrQueriesNotSet
	}
	return nil
}

func GetProfessorRepository(ctx context.Context, u *uow.Uow) repository.ProfessorRepositoryInterface {
	return getRepository[repository.ProfessorRepositoryInterface](ctx, u, "ProfessorRepository")
}

func getRepository[T repository.RepositoryInterface](ctx context.Context, u *uow.Uow, name string) T {
	repo, err := u.GetRepository(ctx, name)
	if err != nil {
		panic(err)
	}
	return repo.(T)
}
