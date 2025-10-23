package dto

import (
	"failiverCheck/internal/app/ds"
)

type CreateComponentDTO struct {
	Title       string `json:"title" validate:"required"`
	Type        string `json:"type" validate:"required"`
	MTBF        uint32 `json:"mtbf" validate:"required,numeric"`
	MTTR        uint32 `json:"mttr" validate:"required,numeric"`
	Description string `json:"description" validate:"required"`
}

type UpdateComponentDTO struct {
	Title       *string  `json:"title"`
	Type        *string  `json:"type"`
	MTBF        *uint32  `json:"mtbf"`
	MTTR        *uint32  `json:"mttr"`
	Available   *float32 `json:"available"`
	Img         *string  `json:"img"`
	Description *string  `json:"description"`
}
type ComponentInSystemCalcDTO struct {
	ID    uint32 `json:"id"`
	Title string `json:"title"`
}

type ComponentsToSystemCalcDTO struct {
	ReplicationCount uint                      `json:"replication_count"`
	Component        *ComponentInSystemCalcDTO `json:"component"`
}

type ComponentsFiltersDTO struct {
	Title  string
	Limit  int
	Offset int
}

func ToComponentDTO(orm ds.Component) ComponentInSystemCalcDTO {
	dto := ComponentInSystemCalcDTO{
		ID:    orm.ID,
		Title: orm.Title,
	}
	return dto
}

func ToComponentsToSystemCalcDTO(orm ds.ComponentsToSystemCalc) ComponentsToSystemCalcDTO {
	var component *ComponentInSystemCalcDTO = nil
	if orm.Component != nil {
		dto := ToComponentDTO(*orm.Component)
		component = &dto
	}

	dto := ComponentsToSystemCalcDTO{
		ReplicationCount: orm.ReplicationCount,
		Component:        component,
	}
	return dto
}
