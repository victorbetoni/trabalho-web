package handlers

import (
	"fmt"
	"strconv"
	"victorbetoni/trabalho-web/internal/domain/usecase"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"

	"github.com/gin-gonic/gin"
)

func PostProfessor(ctx *gin.Context) util.Response {
	input := usecase.ProfessorCreateCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println(err.Error())
		fmt.Println("FUDEU AQUI")
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
