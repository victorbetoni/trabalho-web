package router

import (
	"victorbetoni/trabalho-web/internal/infra/http/handlers"
	"victorbetoni/trabalho-web/internal/infra/http/middleware"
	"victorbetoni/trabalho-web/internal/infra/util"

	"github.com/gin-gonic/gin"
)

type route struct {
	path        string
	requireAuth bool
	method      string
	handler     util.CustomHandlerFunc
}

var routes = []route{
	{path: "/professor", requireAuth: true, method: "GET", handler: handlers.ListProfessores},
	{path: "/professor", requireAuth: true, method: "POST", handler: handlers.PostProfessor},
	{path: "/professor", requireAuth: true, method: "PUT", handler: handlers.UpdateProfessor},
	{path: "/login", requireAuth: false, method: "POST", handler: handlers.Login},
	{path: "/logout", requireAuth: false, method: "GET", handler: handlers.Logout},
	{path: "/checkSession", requireAuth: true, method: "GET", handler: handlers.CheckSession},
}

type httpMethod func(e *gin.Engine, path string, hf gin.HandlerFunc)

var methods = map[string]httpMethod{
	"POST": post,
	"GET":  get,
	"PUT":  put,
}

func Build() *gin.Engine {
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "http://localhost:8091")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	for _, r := range routes {
		if r.requireAuth {
			methods[r.method](engine, r.path, middleware.Parse(middleware.Auth(r.handler)))
		} else {
			methods[r.method](engine, r.path, middleware.Parse(r.handler))
		}
	}

	return engine
}

func post(e *gin.Engine, path string, hf gin.HandlerFunc) {
	e.POST(path, hf)
}

func get(e *gin.Engine, path string, hf gin.HandlerFunc) {
	e.GET(path, hf)
}

func put(e *gin.Engine, path string, hf gin.HandlerFunc) {
	e.PUT(path, hf)
}
