package usecase

import (
	"context"
	"database/sql"
	"victorbetoni/trabalho-web/internal/domain/entity"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"
)

type AulaListCaseInput struct {
	Page int    `json:"page"`
	CPF  string `json:"cpf"`
}

type AulaListCase struct {
	Uow uow.UowInterface
}

func NewAulaListCase(uow uow.UowInterface) *AulaListCase {
	return &AulaListCase{
		Uow: uow,
	}
}

func (u *AulaListCase) Execute(ctx context.Context, in AulaListCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {

		repo := repository.GetAulaRepository(ctx, uow)

		p, err := repo.Find(ctx, entity.AulaFilter{
			Page:      in.Page,
			Professor: in.CPF,
			Limit:     10,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				return util.Ok([]*entity.Aula{})
			}
			return util.InternalServerErr(err)
		}

		return util.Ok(p)
	})
}
