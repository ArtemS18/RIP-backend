package usecase

import (
	"context"
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (uc *UseCase) GetComponent(componetId int) (ds.Component, error) {
	component, err := uc.Postgres.GetComponentById(componetId)
	if err != nil {
		return ds.Component{}, err
	}

	return component, nil
}

func (uc *UseCase) GetComponents(filters dto.ComponentsFiltersDTO) ([]ds.Component, error) {
	var components []ds.Component
	var err error
	components, err = uc.Postgres.GetComponents(filters)
	if err != nil {
		return nil, err
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
	systemCal, err := uc.Postgres.CreateOrGetSystemCalc(userId)
	if err != nil {
		return err
	}

	_, err = uc.Postgres.GetComponentsToSystemCalc(dto.ComponentToSystemCalcByIdDTO{
		ComponentID:         componentId,
		SystemCalculationID: systemCal.ID,
	})
	if err == nil {
		return fmt.Errorf("component (id = %d) alredy added in system calculation (id = %d)", componentId, systemCal.ID)
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	componentsToSystemCalc := ds.ComponentsToSystemCalc{
		ComponentID:         componentId,
		SystemCalculationID: systemCal.ID,
	}
	_, err = uc.Postgres.CreateComponentsToSystemCalc(componentsToSystemCalc)
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
