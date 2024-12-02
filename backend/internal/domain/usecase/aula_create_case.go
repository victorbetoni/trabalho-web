package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"victorbetoni/trabalho-web/internal/domain/entity"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"

	nanoid "github.com/matoous/go-nanoid"
)

type AulaCreateCaseInput struct {
	CPF          string `json:"cpf"`
	Titulo       string `json:"titulo"`
	Data         string `json:"data"`
	CargaHoraria int    `json:"carga_horario"`
	Horario      string `json:"horario"`
}

func (i *AulaCreateCaseInput) Valid() error {
	if !util.ValidCPF(i.CPF) {
		return errors.New("CPF inválido.")
	}
	if len(i.Titulo) < 3 || len(i.Titulo) > 128 {
		return errors.New("O título deve ter entre 3 e 128 caracteres.")
	}
	return nil
}

type AulaCreateCase struct {
	Uow uow.UowInterface
}

func NewAulaCreateCase(uow uow.UowInterface) *AulaCreateCase {
	return &AulaCreateCase{
		Uow: uow,
	}
}

func (u *AulaCreateCase) Execute(ctx context.Context, in AulaCreateCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {

		profRepo := repository.GetProfessorRepository(ctx, uow)
		aulaRepo := repository.GetAulaRepository(ctx, uow)

		if in.Valid() != nil {
			return util.RespBadRequest
		}

		v, err := profRepo.FindOne(ctx, entity.ProfessorFilter{
			CPF: in.CPF,
		})

		if err != nil && err == sql.ErrNoRows {
			return util.BadRequestErr(err)
		}

		if err != nil {
			return util.InternalServerErr(err)
		}

		id, _ := nanoid.ID(12)

		d, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", in.Data, in.Horario))
		if err != nil {
			return util.BadRequestErr(err)
		}

		aula := &entity.Aula{
			Professor:    v,
			ID:           id,
			Titulo:       in.Titulo,
			Data:         d,
			CargaHoraria: in.CargaHoraria,
		}

		if err := aulaRepo.Create(ctx, aula); err != nil {
			return util.InternalServerErr(err)
		}

		return util.Ok(aula)
	})
}
