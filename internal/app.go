package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"online-json-diff/api"
	"online-json-diff/pkg/service"
)

type App struct {
	serviceContainer *service.Container
	router           *gin.Engine
}

func NewApp(serviceContainer *service.Container) *App {
	router := gin.Default()
	app := &App{serviceContainer, router}
	app.registerRouter()
	return app
}

func (a *App) Run(address string) {
	log.Fatal(a.router.Run(address))
}

func (a *App) registerRouter() {
	v1 := a.router.Group("/v1")
	{
		v1.GET("/diff-target", func(ctx *gin.Context) {
			baseHandler[*api.DiffTargetRequest, *api.DiffTargetResponse](ctx, new(api.DiffTargetRequest), a.diffTargetHandler)
		})
	}
}
