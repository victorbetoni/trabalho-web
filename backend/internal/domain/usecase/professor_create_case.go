package usecase

import (
	"context"
	"errors"
	"regexp"
	"victorbetoni/trabalho-web/internal/domain/entity"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"
)

type ProfessorCreateCaseInput struct {
	CPF        string          `json:"cpf"`
	Formacao   string          `json:"formacao"`
	Telefone   string          `json:"telefone"`
	Nome       string          `json:"nome"`
	Endereco   entity.Endereco `json:"endereco"`
	AulasDadas int             `json:"aulasDadas"`
}

func (i *ProfessorCreateCaseInput) Valid() error {
	if !util.ValidCPF(i.CPF) {
		return errors.New("CPF inválido.")
	}
	if len(i.Nome) < 3 || len(i.Nome) > 128 {
		return errors.New("O nome deve ter entre 3 e 128 caracteres.")
	}
	return nil
}

type ProfessorCreateCase struct {
	Uow uow.UowInterface
}

func NewProfessorCreateCase(uow uow.UowInterface) *ProfessorCreateCase {
	return &ProfessorCreateCase{
		Uow: uow,
	}
}

func (u *ProfessorCreateCase) Execute(ctx context.Context, in ProfessorCreateCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {

		repo := repository.GetProfessorRepository(ctx, uow)
		if in.Valid() != nil {
			return util.RespBadRequest
		}

		prof := &entity.Professor{
			CPF:      regexp.MustCompile(`\D`).ReplaceAllString(in.CPF, ""),
			Nome:     in.Nome,
			Formacao: in.Formacao,
			Telefone: in.Telefone,
			Endereco: in.Endereco,
		}

		if err := repo.Create(ctx, prof); err != nil {
			return util.InternalServerErr(err)
		}

		return util.Ok(prof)
	})
}
