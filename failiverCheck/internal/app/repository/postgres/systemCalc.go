package postgres

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"fmt"
	"slices"
	"time"

	"gorm.io/gorm"
)

func (r *Postgres) CreateSystemCalc(userId uint) (ds.SystemCalculation, error) {
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

func (r *Postgres) GetSystemCalcByUserId(userId uint) (ds.SystemCalculation, error) {
	var exist_calc ds.SystemCalculation
	findErr := r.db.Where("user_id = ? AND status = ?", userId, ds.DRAFT).First(&exist_calc).Error
	if findErr != nil {
		return ds.SystemCalculation{}, findErr
	}
	return exist_calc, nil

}
func (r *Postgres) GetSystemCalcById(id uint) (ds.SystemCalculation, error) {
	var sysCalc ds.SystemCalculation
	findErr := r.db.Preload("Moderator").Preload("User").Preload("ComponentsToSystemCalc.Component").Where("id = ? AND status <> ?", id, ds.DELETED).First(&sysCalc).Error
	if findErr != nil {
		return ds.SystemCalculation{}, findErr
	}
	return sysCalc, nil
}

func (r *Postgres) CreateOrGetSystemCalc(userId uint) (ds.SystemCalculation, error) {
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
func (r *Postgres) GetCountInSysCalc(sysCalcId uint) (int64, error) {
	var count int64
	err := r.db.Model(&ds.ComponentsToSystemCalc{}).Where("system_calculation_id = ?", sysCalcId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Postgres) DeleteSystemCalc(sysCalcId uint, userId uint) error {
	date := time.Now()
	res := r.db.Model(ds.SystemCalculation{}).Where(
		"id = ? AND user_id = ? AND status <> ?", sysCalcId, userId, ds.DELETED).Updates(
		ds.SystemCalculation{
			Status:     string(ds.DELETED),
			DateFormed: &date,
		},
	)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return res.Error
}
func (r *Postgres) GetSystemCalcList(filters dto.SearchSystemCalcDTO) ([]ds.SystemCalculation, error) {
	var sys_cacls []ds.SystemCalculation
	allowedStatus := []string{string(ds.COMPLETED), string(ds.FORMED), string(ds.REJECTED)}
	query := r.db.Preload("Moderator").Preload("User").Preload("ComponentsToSystemCalc.Component").Where("status IN (?)", allowedStatus)
	if filters.UserID != nil {
		query = query.Where("user_id = ?", *filters.UserID)
	}
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

func (r *Postgres) UpdateSystemCalc(id uint, dto dto.UpdateSystemCalcDTO) (ds.SystemCalculation, error) {
	var system ds.SystemCalculation
	res := r.db.Model(&system).Where("id = ? AND status != ?", id, ds.DELETED).Updates(dto)
	if res.RowsAffected == 0 {
		return ds.SystemCalculation{}, gorm.ErrRecordNotFound
	}
	if res.Error != nil {
		return ds.SystemCalculation{}, res.Error
	}

	var updatedSystem ds.SystemCalculation
	findErr := r.db.Preload("Moderator").Preload("User").First(&updatedSystem, id).Error
	if findErr != nil {
		return ds.SystemCalculation{}, findErr
	}

	return updatedSystem, nil
}
func (r *Postgres) UpdateSystemCalcStatusModerator(sysCaclId uint, moderatorId uint, command string) (ds.SystemCalculation, error) {
	var sys_cacl ds.SystemCalculation
	err := r.db.Preload("ComponentsToSystemCalc.Component").Where("id = ? AND status = ?", sysCaclId, ds.FORMED).First(&sys_cacl).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return sys_cacl, fmt.Errorf("only sys_calc with status FORMED can be moderating")
		}
		return sys_cacl, err
	}
	var status ds.Status
	var available float32
	switch command {
	case "confirm":
		available, err = calculateAvailable(&sys_cacl)
		if err != nil {
			return sys_cacl, err
		}
		status = ds.COMPLETED
	case "reject":
		status = ds.REJECTED
	default:
		return sys_cacl, fmt.Errorf("invalide command: %v", command)
	}
	timeClosed := time.Now()
	dto := dto.UpdateSystemCalcDTO{Status: &status, DateClosed: &timeClosed, ModeratorId: &moderatorId}
	if available != 0 {
		dto.AvailableCalculation = &available
	}
	return r.UpdateSystemCalc(sysCaclId, dto)

}

func (r *Postgres) UpdateSystemCalcStatusToFormed(id uint) (ds.SystemCalculation, error) {
	sys_cacl, err := r.GetSystemCalcById(id)
	if err != nil {
		return sys_cacl, err
	}
	if sys_cacl.Status != string(ds.DRAFT) {
		return sys_cacl, fmt.Errorf("current sys calc`s status should be the %v", ds.DRAFT)
	}
	if sys_cacl.SystemName == nil || *sys_cacl.SystemName == "" {
		return sys_cacl, fmt.Errorf("system_name should be not null")
	}
	dateFormed := time.Now()
	status := ds.FORMED

	return r.UpdateSystemCalc(id, dto.UpdateSystemCalcDTO{Status: &status, DateFormed: &dateFormed})
}
