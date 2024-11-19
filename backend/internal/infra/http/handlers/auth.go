package handlers

import (
	"net/http"
	"victorbetoni/trabalho-web/internal/domain/usecase"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) util.Response {
	http.SetCookie(ctx.Writer, &http.Cookie{Name: "access_token", Value: "", Path: "/", HttpOnly: true})
	return util.Ok(nil)
}

func Login(ctx *gin.Context) util.Response {
	input := usecase.LoginCaseInput{}
	if err := ctx.BindJSON(&input); err != nil {
		return util.RespBadRequest
	}
	return usecase.NewLoginCase(uow.Current()).Execute(ctx, input)
}

func CheckSession(ctx *gin.Context) util.Response {
	// Se a requisição passou do middleware e chegou aqui, quer dizer que a sessão esta OK.
	return util.Ok(nil)
}
