package usecase

import (
	"failiverCheck/internal/app/config"
)

type UseCase struct {
	Postgres Postgres
	Minio    Minio
	Config   *config.Config
}

func NewUseCase(pg Postgres, minio Minio, c *config.Config) *UseCase {
	return &UseCase{pg, minio, c}
}
