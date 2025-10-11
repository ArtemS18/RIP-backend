package usecase

import (
	"failiverCheck/internal/app/dto"
)

func (uc *UseCase) GetSystemCalc(id uint) (dto.SystemCalculationDTO, error) {
	systemCalc, err := uc.Postgres.GetSystemCalcById(id)
	if err != nil {
		return dto.SystemCalculationDTO{}, err
	}
	sysCalcsResp := dto.ToSystemCalculationDTO(systemCalc)

	return sysCalcsResp, nil
}

func (uc *UseCase) GetSystemCalcList(filters dto.SystemCalcFilters) ([]dto.SystemCalculationInfoDTO, error) {
	sysCalcs, err := uc.Postgres.GetSystemCalcList(filters)
	sysCalcsResp := dto.ToSystemCalculationInfoListDTO(sysCalcs)
	if err != nil {
		return nil, err
	}
	return sysCalcsResp, nil
}

func (uc *UseCase) GetSystemCalcBucket(userId uint) (dto.CurrentUserBucketDTO, error) {
	bucket, err := uc.Postgres.GetCurrentSysCalcAndCount(userId)
	if err != nil {
		return dto.CurrentUserBucketDTO{}, err
	}
	return bucket, nil
}

func (uc *UseCase) UpdateSystemCalc(sysCalcId uint, update dto.UpdateSystemCalcDTO) (dto.SystemCalculationDTO, error) {
	system, err := uc.Postgres.UpdateSystemCalc(sysCalcId, update)
	systemDto := dto.ToSystemCalculationDTO(system)
	if err != nil {
		return dto.SystemCalculationDTO{}, err
	}
	return systemDto, nil
}

func (uc *UseCase) UpdateSystemCalcStatusToFormed(sysCalcId uint) (dto.SystemCalculationDTO, error) {
	system, err := uc.Postgres.UpdateSystemCalcStatusToFormed(sysCalcId)
	if err != nil {
		return dto.SystemCalculationDTO{}, err
	}
	dto := dto.ToSystemCalculationDTO(system)
	return dto, nil
}

func (uc *UseCase) UpdateSystemCalcStatusModerator(sysCalcId uint, moderatorId uint, command string) (dto.SystemCalculationDTO, error) {
	system, err := uc.Postgres.UpdateSystemCalcStatusModerator(uint(sysCalcId), moderatorId, command)
	if err != nil {
		return dto.SystemCalculationDTO{}, err
	}
	dto := dto.ToSystemCalculationDTO(system)
	return dto, nil
}

func (uc *UseCase) DeleteSystemCalc(sysCalcId uint) error {
	err := uc.Postgres.DeleteSystemCalc(sysCalcId)
	if err != nil {
		return err
	}
	return nil
}
