package usecase

import (
	"failiverCheck/internal/app/dto"

	"gorm.io/gorm"
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
		Limit:           filters.Limit,
		Offset:          filters.Offset,
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
	systemCalc, err := uc.Postgres.GetSystemCalcByUserId(userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.CurrentUserBucketDTO{}, nil
		}
		return dto.CurrentUserBucketDTO{}, err
	}

	count, err := uc.Postgres.GetCountInSysCalc(systemCalc.ID)
	if err != nil {
		return dto.CurrentUserBucketDTO{}, err
	}
	bucket := dto.CurrentUserBucketDTO{
		SystemCalculationID: &systemCalc.ID,
		ComponentsCount:     uint(count),
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
