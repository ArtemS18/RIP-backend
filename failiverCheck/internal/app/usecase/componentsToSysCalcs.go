package usecase

import (
	"failiverCheck/internal/app/dto"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (uc *UseCase) UpdateComponentsToSystemCalc(update dto.UpdateComponentToSystemCalcDTO) (dto.ComponentsToSystemCalcDTO, error) {
	ids := dto.ComponentToSystemCalcByIdDTO{
		ComponentID:         update.ComponentID,
		SystemCalculationID: update.SystemCalculationID,
		UserID:              update.UserID,
	}
	component, err := uc.Postgres.GetComponentsToSystemCalc(ids)
	if err != nil {
		return dto.ComponentsToSystemCalcDTO{}, err
	}
	if component.SystemCalculation.UserID != ids.UserID {
		return dto.ComponentsToSystemCalcDTO{}, fmt.Errorf("access denite")
	}
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
	component, err := uc.Postgres.GetComponentsToSystemCalc(ids)
	if err != nil {
		return err
	}
	if component.SystemCalculation.UserID != ids.UserID {
		return fmt.Errorf("access denite")
	}
	err = uc.Postgres.DeleteComponentsToSystemCalc(ids)
	if err != nil {
		return err
	}

	return nil
}
