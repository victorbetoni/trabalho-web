package usecase

import (
	"context"
	"database/sql"
	"errors"
	"victorbetoni/trabalho-web/internal/domain/entity"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"
)

type ProfessorListCaseInput struct {
	CPF   string `json:"cpf"`
	Nome  string `json:"nome"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

func (i *ProfessorListCaseInput) Valid() error {
	if i.Limit < 1 || i.Limit > 100 {
		return errors.New("o limite deve estar entre 1 e 100")
	}
	return nil
}

type ProfessorListCase struct {
	Uow uow.UowInterface
}

func NewProfessorListCase(uow uow.UowInterface) *ProfessorListCase {
	return &ProfessorListCase{
		Uow: uow,
	}
}

func (u *ProfessorListCase) Execute(ctx context.Context, in ProfessorListCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {

		repo := repository.GetProfessorRepository(ctx, uow)
		if err := in.Valid(); err != nil {
			return util.BadRequestErr(err)
		}

		filter := entity.ProfessorFilter{
			CPF:   in.CPF,
			Nome:  in.Nome,
			Limit: in.Limit,
			Page:  in.Page,
		}

		p, err := repo.Find(ctx, filter)
		if err != nil {
			if err == sql.ErrNoRows {
				return util.Ok([]*entity.Professor{})
			}
			return util.InternalServerErr(err)
		}

		return util.Ok(p)
	})
}
