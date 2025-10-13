package usecase

import (
	"context"
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"time"
)

type Postgres interface {
	GetUserByLogin(login string) (ds.User, error)
	GetComponentById(id int) (ds.Component, error)
	GetComponents() ([]ds.Component, error)
	GetComponentsByTitle(title string) ([]ds.Component, error)
	UpdateComponentById(id uint, update dto.UpdateComponentDTO) (ds.Component, error)
	CreateComponent(create dto.CreateComponentDTO) (ds.Component, error)
	DeletedComponentById(id uint) error
	AddComponentInSystemCalc(componentId, userId uint) error

	GetSystemCalcById(id uint) (ds.SystemCalculation, error)
	GetSystemCalcList(dto dto.SearchSystemCalcDTO) ([]ds.SystemCalculation, error)
	GetCurrentSysCalcAndCount(userId uint) (dto.CurrentUserBucketDTO, error)
	UpdateSystemCalcStatusToFormed(userId uint) (ds.SystemCalculation, error)
	DeleteSystemCalc(userId uint, id uint) error
	UpdateSystemCalcStatusModerator(sysCaclId uint, moderatorId uint, command string) (ds.SystemCalculation, error)
	UpdateSystemCalc(sysCalcId uint, update dto.UpdateSystemCalcDTO) (ds.SystemCalculation, error)

	RegisterUser(credentials schemas.UserCredentials) (ds.User, error)
	LogoutUser(userId uint) error
	GetUserById(id uint) (ds.User, error)
	UpdateUserById(id uint, update dto.UserUpdateDTO) (ds.User, error)
}

type Minio interface {
	UploadComponentImg(ctx context.Context, uploadDTO dto.ComponentImgCreateDTO) (string, error)
	DeleteComponentImg(ctx context.Context, objectName *string) error
}

type Redis interface {
	SetBlackListJWT(ctx context.Context, token string, jwtTTL time.Duration) error
	GetBlackListJWT(ctx context.Context, token string) error
}
