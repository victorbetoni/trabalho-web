package handlers

import (
	"strconv"
	"victorbetoni/trabalho-web/internal/domain/usecase"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"

	"github.com/gin-gonic/gin"
)

func PostAula(ctx *gin.Context) util.Response {
	input := usecase.AulaCreateCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewAulaCreateCase(uow.Current()).Execute(ctx, input)
}

func DeleteAula(ctx *gin.Context) util.Response {
	input := usecase.AulaDeleteCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewAulaDeleteCase(uow.Current()).Execute(ctx, input)
}

func ListAulas(ctx *gin.Context) util.Response {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return util.BadRequestErr(err)
	}

	cpf := ctx.Query("cpf")
	if err != nil {
		return util.BadRequestErr(err)
	}

	input := usecase.AulaListCaseInput{
		Page: page,
		CPF:  cpf,
	}
	return usecase.NewAulaListCase(uow.Current()).Execute(ctx, input)
}
