package usecase

import (
	"failiverCheck/internal/app/config"
)

type UseCase struct {
	Postgres Postgres
	Minio    Minio
	Config   *config.Config
	Redis    Redis
}

func NewUseCase(pg Postgres, minio Minio, c *config.Config, r Redis) *UseCase {
	return &UseCase{pg, minio, c, r}
}
