package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Body    any    `json:"body"`
}

func (r *Response) Ok() bool {
	return r.Status == 200
}

type CustomHandlerFunc func(ctx *gin.Context) Response

var (
	RespBadRequest = Response{
		Status:  http.StatusBadRequest,
		Message: "Preencha todos os campos.",
		Body:    "",
	}
)

func Ok(body any) Response {
	return Response{
		Status:  http.StatusOK,
		Message: "",
		Body:    body,
	}
}

func InternalServerErr(err error) Response {
	fmt.Println(err.Error())
	return Response{
		Status:  http.StatusInternalServerError,
		Message: "Algo deu errado.",
		Body:    "",
	}
}

func BadRequestErr(err error) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
		Body:    "",
	}
}
