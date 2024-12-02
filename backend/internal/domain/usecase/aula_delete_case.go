package usecase

import (
	"context"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"
)

type AulaDeleteCaseInput struct {
	ID string `json:"id"`
}

type AulaDeleteCase struct {
	Uow uow.UowInterface
}

func NewAulaDeleteCase(uow uow.UowInterface) *AulaDeleteCase {
	return &AulaDeleteCase{
		Uow: uow,
	}
}

func (u *AulaDeleteCase) Execute(ctx context.Context, in AulaDeleteCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {

		aulaRepo := repository.GetAulaRepository(ctx, uow)

		if err := aulaRepo.Delete(ctx, in.ID); err != nil {
			return util.InternalServerErr(err)
		}

		return util.Ok(nil)
	})
}
