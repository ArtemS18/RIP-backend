package main

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dsn"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()
	dsn := dsn.LoadEnv()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&ds.Component{},
		&ds.User{},
		&ds.SystemCalculation{},
		&ds.ComponentsToSystemCalc{},
	)
	if err != nil {
		panic("cant migrate db")
	}
}
