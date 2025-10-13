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

func (uc *UseCase) GetSystemCalcList(user dto.UserDTO, filters dto.SystemCalcFilters) ([]dto.SystemCalculationInfoDTO, error) {
	searchDto := dto.SearchSystemCalcDTO{
		DateFormedStart: filters.DateFormedStart,
		DateFormedEnd:   filters.DateFormedEnd,
		Status:          filters.Status,
	}
	if !user.IsModerator {
		searchDto.UserID = &user.ID
	}
	sysCalcs, err := uc.Postgres.GetSystemCalcList(searchDto)
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

func (uc *UseCase) DeleteSystemCalc(userId uint, sysCalcId uint) error {
	err := uc.Postgres.DeleteSystemCalc(userId, sysCalcId)
	if err != nil {
		return err
	}
	return nil
}
