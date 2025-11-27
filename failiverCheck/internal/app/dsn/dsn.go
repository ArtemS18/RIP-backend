package dsn

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	_ = godotenv.Load()
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

}
