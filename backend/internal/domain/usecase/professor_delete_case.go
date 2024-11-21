package usecase

import (
	"context"
	"errors"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"
)

type ProfessorDeleteCaseInput struct {
	CPF string `json:"cpf"`
}

func (i *ProfessorDeleteCaseInput) Valid() error {
	if !util.ValidCPF(i.CPF) {
		return errors.New("CPF inválido.")
	}
	return nil
}

type ProfessorDeleteCase struct {
	Uow uow.UowInterface
}

func NewProfessorDeleteCase(uow uow.UowInterface) *ProfessorDeleteCase {
	return &ProfessorDeleteCase{
		Uow: uow,
	}
}

func (u *ProfessorDeleteCase) Execute(ctx context.Context, in ProfessorDeleteCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {

		repo := repository.GetProfessorRepository(ctx, uow)
		if in.Valid() != nil {
			return util.RespBadRequest
		}

		if err := repo.Delete(ctx, in.CPF); err != nil {
			return util.InternalServerErr(err)
		}

		return util.Ok(nil)
	})
}
