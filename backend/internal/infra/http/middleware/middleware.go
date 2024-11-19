package middleware

import (
	"encoding/json"
	"net/http"
	"victorbetoni/trabalho-web/internal/infra/http/jwt"
	"victorbetoni/trabalho-web/internal/infra/util"

	"github.com/gin-gonic/gin"
)

func Parse(hf util.CustomHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := hf(ctx)
		if resp.Body == "" {
			resp.Body = "{}"
		}
		marshaled, err := json.Marshal(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": resp.Status, "message": "Não foi possível fazer parse do body.", "body": "{}"})
			return
		}
		ctx.JSON(resp.Status, gin.H{"status": resp.Status, "message": resp.Message, "body": string(marshaled)})
	}
}

func Auth(next util.CustomHandlerFunc) util.CustomHandlerFunc {
	return func(ctx *gin.Context) util.Response {
		token, err := ctx.Cookie("access_token")
		if err != nil {
			return util.Response{
				Status:  http.StatusUnauthorized,
				Message: "Token não encontrado.",
				Body:    "{}",
			}
		}

		if _, err = jwt.ExtractUserIdentifier(token); err != nil {
			return util.Response{
				Status:  http.StatusUnauthorized,
				Message: "Token inválido ou expirado.",
				Body:    "{}",
			}
		}

		return next(ctx)
	}
}
