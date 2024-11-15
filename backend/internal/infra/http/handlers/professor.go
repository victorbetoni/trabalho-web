package handlers

import (
	"victorbetoni/trabalho-web/internal/domain/usecase"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"

	"github.com/gin-gonic/gin"
)

func PostProfessor(ctx *gin.Context) util.Response {
	input := usecase.ProfessorCreateCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewProfessorCreateCase(uow.Current()).Execute(ctx, input)
}

func ListProfessores(ctx *gin.Context) util.Response {
	input := usecase.ProfessorListCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewProfessorListCase(uow.Current()).Execute(ctx, input)
}
