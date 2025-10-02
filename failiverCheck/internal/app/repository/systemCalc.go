package repository

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"fmt"
	"slices"
	"time"

	"gorm.io/gorm"
)

func (r *Repository) CreateSystemCalc(userId uint) (ds.SystemCalculation, error) {
	newSystemCalc := ds.SystemCalculation{
		UserID:      userId,
		ModeratorID: nil,
	}
	createErr := r.db.Create(&newSystemCalc).Error
	if createErr != nil {
		return ds.SystemCalculation{}, createErr
	}
	return newSystemCalc, nil
}

func (r *Repository) GetSystemCalcByUserId(userId uint) (ds.SystemCalculation, error) {
	var exist_calc ds.SystemCalculation
	findErr := r.db.Where("user_id = ? AND status = ?", userId, ds.DRAFT).First(&exist_calc).Error
	if findErr != nil {
		return ds.SystemCalculation{}, findErr
	}
	return exist_calc, nil

}
func (r *Repository) GetSystemCalcById(id uint) (ds.SystemCalculation, error) {
	var sysCalc ds.SystemCalculation
	findErr := r.db.Preload("Moderator").Preload("User").Preload("ComponentsToSystemCalc.Component").Where("id = ? AND status <> ?", id, ds.DELETED).First(&sysCalc).Error
	if findErr != nil {
		return ds.SystemCalculation{}, findErr
	}
	return sysCalc, nil
}

func (r *Repository) CreateOrGetSystemCalc(userId uint) (ds.SystemCalculation, error) {
	exist_calc, findErr := r.GetSystemCalcByUserId(userId)
	if findErr != nil {
		if findErr == gorm.ErrRecordNotFound {
			return r.CreateSystemCalc(userId)
		} else {
			return ds.SystemCalculation{}, findErr
		}
	}

	return exist_calc, nil

}

func (r *Repository) AddComponentInSystemCalc(componentID uint, userId uint) error {
	systemCal, err := r.CreateOrGetSystemCalc(userId)
	if err != nil {
		return err
	}

	var existing ds.ComponentsToSystemCalc
	check := r.db.Where("component_id = ? AND system_calculation_id = ?", componentID, systemCal.ID).First(&existing)
	if check.Error == nil {
		return fmt.Errorf("component (id = %d) alredy added in system calculation (id = %d)", componentID, systemCal.ID)
	}
	if check.Error != nil && check.Error != gorm.ErrRecordNotFound {
		return check.Error
	}

	componentsToSystemCalc := ds.ComponentsToSystemCalc{
		ComponentID:         componentID,
		SystemCalculationID: systemCal.ID,
	}
	err = r.db.Create(&componentsToSystemCalc).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetCurrentSysCalcAndCount(userId uint) (dto.CurrentUserBucketDTO, error) {
	systemCalc, err := r.GetSystemCalcByUserId(userId)
	dto := dto.CurrentUserBucketDTO{}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto, nil
		}
		return dto, err
	}
	var count int64
	err = r.db.Model(&ds.ComponentsToSystemCalc{}).Where("system_caclulation_id = ?", systemCalc.ID).Count(&count).Error
	if err != nil {
		return dto, err
	}
	dto.SystemCalculationID = &systemCalc.ID
	dto.ComponentsCount = uint(count)
	return dto, nil

}

func (r *Repository) DeleteComponentFromSystemCalc(sysCalcId uint, componentId uint) error {
	var deletedComponent ds.Component

	findErr := r.db.Where("system_caclulation_id = ?, component_id = ?", sysCalcId, componentId).First(&deletedComponent).Error
	if findErr != nil {
		return findErr
	}
	deleteErr := r.db.Delete(deletedComponent).Error
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
func (r *Repository) DeleteSystemCalc(sysCalcId uint) error {
	var err error
	var ids []uint
	query := "UPDATE system_calculations SET status=$1 WHERE id = $2 AND status!=$1"
	res := r.db.Raw(query, ds.DELETED, sysCalcId).Scan(&ids)
	if err = res.Error; err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetSystemCalcList(filters dto.SystemCalcFilters) ([]ds.SystemCalculation, error) {
	var sys_cacls []ds.SystemCalculation
	allowedStatus := []string{string(ds.COMPLETED), string(ds.FORMED), string(ds.REJECTED)}
	query := r.db.Preload("Moderator").Preload("User").Preload("ComponentsToSystemCalc.Component").Where("status IN (?)", allowedStatus)
	if filters.Status != nil {
		if slices.Contains(allowedStatus, string(*filters.Status)) {
			query = query.Where("status = ?", filters.Status)
		} else {
			return nil, fmt.Errorf("not valid status")
		}
	}
	if filters.DateFormedStart != nil {
		dateFormedStartValue := *filters.DateFormedStart
		s, err := time.Parse("2006-01-02", dateFormedStartValue)
		if err != nil {
			return nil, fmt.Errorf("invalid DateFormedStart format: %w", err)
		}
		query = query.Where("date(date_formed) >= ?", s)
	}

	if filters.DateFormedEnd != nil {
		dateFormedEndValue := *filters.DateFormedEnd
		e, err := time.Parse("2006-01-02", dateFormedEndValue)
		if err != nil {
			return nil, fmt.Errorf("invalid DateFormedEnd format: %w", err)
		}
		query = query.Where("date(date_formed) <= ?", e)
	}

	err := query.Find(&sys_cacls).Error

	if err != nil {
		return nil, err
	}
	if len(sys_cacls) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return sys_cacls, nil
}

func (r *Repository) UpdateSystemCalc(id uint, dto dto.UpdateSystemCalcDTO) error {
	res := r.db.Model(&ds.SystemCalculation{}).Where("id = ? AND status != ?", id, ds.DELETED).Updates(dto)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if res.Error != nil {
		return res.Error
	}
	return nil

}
func (r *Repository) UpdateSystemCalcStatusModerator(sysCaclId uint, moderatorId uint, command string) error {
	var sys_cacl ds.SystemCalculation
	err := r.db.Preload("ComponentsToSystemCalc.Component").Where("id = ? AND status = ?", sysCaclId, ds.FORMED).First(&sys_cacl).Error
	if err != nil {
		return err
	}
	var status ds.Status
	var available float32
	switch command {
	case "confirm":
		available, err = calculateAvailable(&sys_cacl)
		if err != nil {
			return err
		}
		status = ds.COMPLETED
	case "reject":
		status = ds.REJECTED
	default:
		return fmt.Errorf("invalide command: %v", command)
	}
	timeClosed := time.Now()
	dto := dto.UpdateSystemCalcDTO{Status: &status, DateClosed: &timeClosed, ModeratorId: &moderatorId}
	if available != 0 {
		dto.AvailableCalculation = &available
	}
	return r.UpdateSystemCalc(sysCaclId, dto)

}

func (r *Repository) UpdateSystemCalcStatusToFormed(id uint) error {
	sys_cacl, err := r.GetSystemCalcById(id)
	if err != nil {
		return err
	}
	if sys_cacl.Status != string(ds.DRAFT) {
		return fmt.Errorf("current sys calc`s status should be the %v", ds.DRAFT)
	}
	if sys_cacl.SystemName == nil || *sys_cacl.SystemName == "" {
		return fmt.Errorf("system_name should be not null")
	}
	dateFormed := time.Now()
	status := ds.FORMED
	return r.UpdateSystemCalc(id, dto.UpdateSystemCalcDTO{Status: &status, DateFormed: &dateFormed})
}
