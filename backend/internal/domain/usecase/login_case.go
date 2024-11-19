package usecase

import (
	"context"
	"net/http"
	"victorbetoni/trabalho-web/internal/infra/http/jwt"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/internal/infra/util"
	"victorbetoni/trabalho-web/pkg/uow"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginCaseInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginCase struct {
	Uow uow.UowInterface
}

func NewLoginCase(uow uow.UowInterface) *LoginCase {
	return &LoginCase{
		Uow: uow,
	}
}

func (u *LoginCase) Execute(ctx context.Context, in LoginCaseInput) util.Response {
	return u.Uow.Do(ctx, func(uow *uow.Uow) util.Response {
		loginRepo := repository.GetLoginRepository(ctx, uow)

		senha, err := loginRepo.FindLogin(ctx, in.Username)
		if err != nil {
			return util.Response{
				Status:  http.StatusUnauthorized,
				Message: "Usuário não necontrado.",
			}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(senha), []byte(in.Password)); err != nil {
			return util.Response{
				Status:  http.StatusUnauthorized,
				Message: "Senha errada.",
			}
		}

		token := jwt.GenerateToken(in.Username)

		ctx, _ := ctx.(*gin.Context)
		http.SetCookie(ctx.Writer, &http.Cookie{Name: "access_token", Path: "/", Secure: false, Value: token, MaxAge: 604800, HttpOnly: true})

		return util.Ok(in.Username)
	})
}
