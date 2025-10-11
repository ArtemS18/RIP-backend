package usecase

import "failiverCheck/internal/app/ds"

type Postgres interface {
	GetUserByLogin(login string) (ds.User, error)
	GetComponentById(id int) (ds.Component, error)
	GetComponents() ([]ds.Component, error)
	GetComponentsByTitle(title string) ([]ds.Component, error)
}
