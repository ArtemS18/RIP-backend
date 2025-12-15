package main

import (
	"failiverCheck/internal/app/adapters/http"
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/dsn"
	"failiverCheck/internal/app/repository/minio"
	"failiverCheck/internal/app/repository/postgres"
	"failiverCheck/internal/app/repository/redis"
	"failiverCheck/internal/app/usecase"
	"failiverCheck/internal/pkg/app"
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
	configCORS.AllowOrigins = []string{"*"}
	configCORS.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}
	//configCORS.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(configCORS))
	dsnPg := dsn.LoadEnv()
	fmt.Println(dsnPg)
	pg, errRep := postgres.NewPostgers(dsnPg)
	if errRep != nil {
		logrus.Fatalf("error initializing repository: %v", errRep)
	}
	minio, err := minio.NewMinio(config.Minio)
	if err != nil {
		logrus.Fatalf("error initializing minio: %v", errRep)
	}
	redis, err := redis.New(config.Redis)
	if err != nil {
		logrus.Fatalf("error initializing redis: %v", errRep)
	}
	uc := usecase.NewUseCase(pg, minio, config, &redis)
	handler := http.NewHandler(pg, minio, uc, config)

	app := app.NewApplication(config, router, handler)
	_, err = handler.UseCase.UpdateSystemCalcStatusModerator(1, 1, "confirm")
	if err != nil {
		logrus.Fatalf("error in test update system calc: %v", errRep)
	}
	app.RunApplication()
	log.Println("App terminated")
}
