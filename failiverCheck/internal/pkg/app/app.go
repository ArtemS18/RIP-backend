package app

import (
	"failiverCheck/internal/app/adapters/http"
	"failiverCheck/internal/app/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Config  *config.Config
	Router  *gin.Engine
	Handler *http.Handler
}

func NewApplication(c *config.Config, r *gin.Engine, h *http.Handler) *Application {
	return &Application{
		Config:  c,
		Router:  r,
		Handler: h,
	}
}

func (app *Application) RunApplication() {
	logrus.Println("Server start up")
	app.Handler.RegisterHandlers(app.Router)
	address := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	if err := app.Router.Run(address); err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("Server down")

}
