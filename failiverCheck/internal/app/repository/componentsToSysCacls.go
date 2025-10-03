package repository

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"

	"gorm.io/gorm"
)

func (r *Repository) UpdateComponentsToSystemCalc(update dto.UpdateComponentToSystemCalcDTO) (ds.ComponentsToSystemCalc, error) {
	var new ds.ComponentsToSystemCalc
	res := r.db.Model(&new).Where("system_calculation_id = ? AND component_id = ?", update.SystemCalculationID, update.ComponentID).Update("replication_count", update.ReplicationCount)
	r.db.Preload("Component").Where("system_calculation_id = ? AND component_id = ?", update.SystemCalculationID, update.ComponentID).First(&new)
	if res.RowsAffected == 0 {
		return ds.ComponentsToSystemCalc{}, gorm.ErrRecordNotFound
	}
	return new, res.Error
}

func (r *Repository) DeleteComponentsToSystemCalc(ids dto.ComponentToSystemCalcByIdDTO) error {
	var deletedComponent ds.ComponentsToSystemCalc

	res := r.db.Where("system_calculation_id = ? AND component_id = ?", ids.SystemCalculationID, ids.ComponentID).Delete(&deletedComponent)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
