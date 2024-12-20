package handlers

import (
	"strconv"
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

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return util.BadRequestErr(err)
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		return util.BadRequestErr(err)
	}

	input := usecase.ProfessorListCaseInput{
		CPF:   ctx.Query("cpf"),
		Nome:  ctx.Query("nome"),
		Limit: limit,
		Page:  page,
	}
	return usecase.NewProfessorListCase(uow.Current()).Execute(ctx, input)
}

func UpdateProfessor(ctx *gin.Context) util.Response {
	input := usecase.ProfessorUpdateCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewProfessorUpdateCase(uow.Current()).Execute(ctx, input)
}

func DeleteProfessor(ctx *gin.Context) util.Response {
	input := usecase.ProfessorDeleteCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewProfessorDeleteCase(uow.Current()).Execute(ctx, input)
}
