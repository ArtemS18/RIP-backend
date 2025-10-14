package usecase

import (
	"failiverCheck/internal/app/dto"

	"github.com/go-playground/validator/v10"
)

func (uc *UseCase) UpdateComponentsToSystemCalc(update dto.UpdateComponentToSystemCalcDTO) (dto.ComponentsToSystemCalcDTO, error) {
	orm, err := uc.Postgres.UpdateComponentsToSystemCalc(update)
	new := dto.ToComponentsToSystemCalcDTO(orm)
	if err != nil {
		return dto.ComponentsToSystemCalcDTO{}, err
	}
	return new, nil
}

func (uc *UseCase) DeleteComponentsToSystemCalc(ids dto.ComponentToSystemCalcByIdDTO) error {
	validate := validator.New()
	if err := validate.Struct(ids); err != nil {
		return err
	}
	err := uc.Postgres.DeleteComponentsToSystemCalc(ids)
	if err != nil {
		return err
	}

	return nil
}
