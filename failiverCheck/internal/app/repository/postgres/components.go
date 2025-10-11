package postgres

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"fmt"

	"gorm.io/gorm/clause"
)

func (r *Postgres) GetComponents() ([]ds.Component, error) {
	var components []ds.Component
	err := r.db.Where("is_deleted = ?", false).Find(&components).Error
	if err != nil {
		return nil, err
	}
	if len(components) == 0 {
		return nil, fmt.Errorf("records not found")
	}
	return components, nil
}

func (r *Postgres) GetComponentById(id int) (ds.Component, error) {
	var component ds.Component

	err := r.db.Where("id = ? AND is_deleted = ?", id, false).First(&component).Error
	if err != nil {
		return ds.Component{}, err
	}

	return component, nil
}

func (r *Postgres) GetComponentsByTitle(title string) ([]ds.Component, error) {
	var components []ds.Component

	err := r.db.Where("title ILIKE ? AND is_deleted = ?", "%"+title+"%", false).Find(&components).Error

	if err != nil {
		return nil, err
	}

	return components, nil
}

func (r *Postgres) UpdateComponentById(id uint, update dto.UpdateComponentDTO) (ds.Component, error) {
	var component ds.Component
	res := r.db.Model(&component).Clauses(clause.Returning{}).Where("id = ? AND is_deleted = ?", id, false).Updates(update)
	if res.Error != nil {
		return ds.Component{}, res.Error
	}
	if res.RowsAffected == 0 {
		return ds.Component{}, fmt.Errorf("record not found")
	}
	return component, nil
}

func (r *Postgres) CreateComponent(create dto.CreateComponentDTO) (ds.Component, error) {

	if create.MTBF <= 0 || create.MTTR <= 0 {
		return ds.Component{}, fmt.Errorf("mtbf and mttr shouldn`t be least nil")
	}
	availableCalc := float32(create.MTBF) / float32(create.MTTR+create.MTBF)

	component := ds.Component{
		Title:       create.Title,
		Type:        create.Type,
		MTBF:        create.MTBF,
		MTTR:        create.MTTR,
		Available:   float32(availableCalc),
		Description: create.Description,
	}

	if err := r.db.Create(&component).Error; err != nil {
		return ds.Component{}, err
	}
	return component, nil
}

func (r *Postgres) DeletedComponentById(id uint) error {
	var component ds.Component
	res := r.db.Model(&component).Where("id = ? AND is_deleted = ?", id, false).Update("is_deleted", true)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}
