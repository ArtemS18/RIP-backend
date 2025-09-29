package repository

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"fmt"

	"gorm.io/gorm/clause"
)

func (r *Repository) GetComponents() ([]ds.Component, error) {
	var components []ds.Component
	err := r.db.Where("is_deleted = ?", false).Find(&components).Error
	if err != nil {
		return nil, err
	}
	if len(components) == 0 {
		return nil, fmt.Errorf("massive is empty")
	}
	return components, nil
}

func (r *Repository) GetComponentById(id int) (ds.Component, error) {
	var component ds.Component

	err := r.db.Where("id = ? AND is_deleted = ?", id, false).First(&component).Error
	if err != nil {
		return ds.Component{}, err
	}

	return component, nil
}

func (r *Repository) GetComponentsByTitle(title string) ([]ds.Component, error) {
	var components []ds.Component

	err := r.db.Where("title ILIKE ? AND is_deleted = ?", "%"+title+"%", false).Find(&components).Error

	if err != nil {
		return nil, err
	}

	return components, nil
}

func (r *Repository) UpdateComponentById(id uint, update dto.UpdateComponentDTO) (ds.Component, error) {
	var component ds.Component
	err := r.db.Model(&ds.Component{}).Clauses(clause.Returning{}).Where("id = ?", id).Updates(update).Scan(&component).Error
	if err != nil {
		return ds.Component{}, err
	}
	if component.Title == "" {
		return ds.Component{}, fmt.Errorf("record not found")
	}
	return component, nil
}

func (r *Repository) CreateComponent(create dto.CreateComponentDTO) (ds.Component, error) {

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

func (r *Repository) DeletedComponentById(id uint) error {
	var component ds.Component
	err := r.db.Model(&component).Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "title"},
		}}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	if component.Title == "" {
		return fmt.Errorf("record not found")
	}
	return nil
}
