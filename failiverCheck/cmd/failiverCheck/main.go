package main

import (
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/dsn"
	"failiverCheck/internal/app/handler"
	"failiverCheck/internal/app/repository"
	"failiverCheck/internal/pkg"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT

// @host      localhost:8080
// @BasePath  /api/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	log.Println("App start")

	router := gin.Default()
	config, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("error loading config: %v", err)
	}
	configCORS := cors.DefaultConfig()
	configCORS.AllowOrigins = []string{"*"} // Allow any origin (development only)
	router.Use(cors.New(configCORS))
	dsnPg := dsn.LoadEnv()
	fmt.Println(dsnPg)
	repo, errRep := repository.NewRepository(dsnPg, config)
	if errRep != nil {
		logrus.Fatalf("error initializing repository: %v", errRep)
	}

	handler := handler.NewHandler(repo)

	app := pkg.NewApplication(config, router, handler)
	app.RunApplication()
	log.Println("App terminated")
}
