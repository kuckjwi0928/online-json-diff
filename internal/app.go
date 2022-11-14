package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online-json-diff/api"
	"online-json-diff/configs"
	"online-json-diff/pkg/service"
)

type App struct {
	serviceContainer *service.Container
	router           *gin.Engine
}

func NewApp(serviceContainer *service.Container) *App {
	if configs.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	app := &App{serviceContainer, router}
	app.registerRouter()
	return app
}

func (a *App) Run(address string) {
	log.Fatal(a.router.Run(address))
}

func (a *App) registerRouter() {
	if configs.Server.Env == "prod" {
		a.router.Static("/static", "web/static")
		a.router.LoadHTMLFiles("web/index.html")
		a.router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}
	v1 := a.router.Group("/v1")
	{
		v1.GET("/diff-target", baseHandler[*api.DiffTargetRequest, *api.DiffTargetResponse](new(api.DiffTargetRequest), a.handleDiffTarget))
		v1.GET("/diff", baseHandler[*api.DiffTargetRequest, *api.DiffTargetResponse](new(api.DiffTargetRequest), a.handleDiff))
	}
}
