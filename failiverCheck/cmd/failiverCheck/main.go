package main

import (
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/dsn"
	"failiverCheck/internal/app/handler"
	"failiverCheck/internal/app/repository"
	"failiverCheck/internal/pkg"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("App start")

	router := gin.Default()
	config, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("error loading config: %v", err)
	}
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
