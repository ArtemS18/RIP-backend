package usecase

import (
	"context"
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"

	log "github.com/sirupsen/logrus"
)

func (uc *UseCase) GetComponent(componetId int) (ds.Component, error) {
	component, err := uc.Postgres.GetComponentById(componetId)
	if err != nil {
		return ds.Component{}, err
	}

	return component, nil
}

func (uc *UseCase) GetComponents(searchQuery string) ([]ds.Component, error) {
	var components []ds.Component
	var err error
	log.Info(searchQuery)
	if searchQuery == "" {
		components, err = uc.Postgres.GetComponents()
		if err != nil {
			return nil, err
		}
	} else {
		components, err = uc.Postgres.GetComponentsByTitle(searchQuery)
		if err != nil {
			return nil, err
		}
	}
	return components, nil
}

func (uc *UseCase) UpdateComponent(componentId uint, update dto.UpdateComponentDTO) (ds.Component, error) {
	component, err := uc.Postgres.UpdateComponentById(componentId, update)
	if err != nil {
		return ds.Component{}, err
	}
	return component, nil
}
func (uc *UseCase) CreateComponent(create dto.CreateComponentDTO) (ds.Component, error) {
	component, err := uc.Postgres.CreateComponent(create)
	if err != nil {
		return ds.Component{}, err
	}
	return component, nil
}

func (uc *UseCase) DeleteComponent(componentId uint) error {

	component, err := uc.Postgres.GetComponentById(int(componentId))
	if err != nil {
		return err
	}
	imgUrl := component.Img
	if err := uc.Postgres.DeletedComponentById(componentId); err != nil {
		return err
	}
	ctx := context.Background()
	if err := uc.Minio.DeleteComponentImg(ctx, &imgUrl); err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) AddComponentInSystemCalc(userId uint, componentId uint) error {
	err := uc.Postgres.AddComponentInSystemCalc(componentId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) UpdateComponentImg(componentId uint, uploadDTO dto.ComponentImgCreateDTO) (string, error) {
	ctx := context.Background()
	location, err := uc.Minio.UploadComponentImg(ctx, uploadDTO)
	log.Info(location)
	if err != nil {
		return "", err
	}
	component, err := uc.Postgres.GetComponentById(int(componentId))
	if err != nil {
		return "", err
	}
	if component.Img != "" {
		if err = uc.Minio.DeleteComponentImg(ctx, &component.Img); err != nil {
			log.Error(err)
		}
	}
	uc.Postgres.UpdateComponentById(componentId, dto.UpdateComponentDTO{Img: &location})
	return location, nil
}
