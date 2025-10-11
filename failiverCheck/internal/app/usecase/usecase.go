package usecase

import (
	"failiverCheck/internal/app/config"
	"failiverCheck/internal/app/ds"
)

type UseCase struct {
	Postgres Postgres
	Config   *config.Config
}

type Postgres interface {
	GetUserByLogin(login string) (ds.User, error)
}

func NewUseCase(pg Postgres, c *config.Config) *UseCase {
	return &UseCase{pg, c}
}
