package repository

import (
	"failiverCheck/internal/app/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	db    *gorm.DB
	minio *Minio
}

func NewRepository(dsn string, config *config.Config) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	logrus.Info(config.Minio.Host)
	minio, err := NewMinio(config)
	if err != nil {
		return nil, err
	}
	return &Repository{db, minio}, nil
}
