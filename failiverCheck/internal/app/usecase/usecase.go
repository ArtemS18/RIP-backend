package usecase

import (
	"failiverCheck/internal/app/config"
)

type UseCase struct {
	Postgres Postgres
	Config   *config.Config
}

func NewUseCase(pg Postgres, c *config.Config) *UseCase {
	return &UseCase{pg, c}
}
