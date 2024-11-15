package router

import (
	"victorbetoni/trabalho-web/internal/infra/http/handlers"
	"victorbetoni/trabalho-web/internal/infra/http/middleware"

	"github.com/gin-gonic/gin"
)

func Build() *gin.Engine {
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "http://frontend:8090")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	engine.POST("/professor", middleware.Parse(handlers.PostProfessor))
	return engine
}
