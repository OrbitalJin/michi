package router

import (
	"github.com/OrbitalJin/qmuxr/internal/router/handler"
	"github.com/OrbitalJin/qmuxr/internal/templater"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterIface interface {
	GetEngine() *gin.Engine
	Route()
	Up(port string)
}

type Router struct {
	Engine    *gin.Engine
	handler   *handler.Handler
	templater *templater.Templater
}

func NewRouter(handler *handler.Handler) (*Router, error) {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	templater, err := templater.New()

	if err != nil {
		return nil, err
	}

	engine.SetHTMLTemplate(templater.GetHTMLTemplates())

	return &Router{
		Engine:    engine,
		handler:   handler,
		templater: templater,
	}, nil
}

func (r *Router) Route() {
	r.Engine.GET("/search", r.handler.Root)
	r.Engine.GET("/error", r.handler.Error)
	r.Engine.GET("/session_success", r.handler.SessionOpened)
	r.Engine.GET("/favicon.ico", r.handler.Favicon)
}

func (r *Router) GetEngine() *gin.Engine {
	return r.Engine

}

func (r *Router) Up(port string) {
	r.Engine.Run(port)
}
