package dto

import (
	"failiverCheck/internal/app/ds"
	"time"
)

type CurrentUserBucketDTO struct {
	SystemCalculationID *uint `json:"sys_calculation_id"`
	ComponentsCount     uint  `json:"components_count"`
}

type SystemCalcFilters struct {
	DateFormedStart *string    `form:"date_formed_start"`
	DateFormedEnd   *string    `form:"date_formed_end"`
	Status          *ds.Status `form:"status"`
}

type UpdateSystemCalcDTO struct {
	SystemName           *string    `json:"system_name"`
	Status               *ds.Status `json:"status"`
	DateFormed           *time.Time
	DateClosed           *time.Time
	ModeratorId          *uint
	AvailableCalculation *float32
}

type SystemCalculationDTO struct {
	ID                   uint       `json:"id"`
	SystemName           *string    `json:"system_name"`
	AvailableCalculation float32    `json:"available_calculation"`
	Status               string     `json:"status"`
	DateCreated          time.Time  `json:"date_created"`
	DateFormed           *time.Time `json:"date_formed"`
	DateClosed           *time.Time `json:"date_accepted"`

	User                   UserDTO                     `json:"user"`
	Moderator              *UserDTO                    `json:"moderator"`
	ComponentsToSystemCalc []ComponentsToSystemCalcDTO `json:"components"`
}

func ToSystemCalculationDTO(orm ds.SystemCalculation) SystemCalculationDTO {
	moderator := ToUserDTO(*orm.Moderator)
	var componentsToSystemCalc []ComponentsToSystemCalcDTO
	for _, el := range orm.ComponentsToSystemCalc {
		componentsToSystemCalc = append(componentsToSystemCalc, ToComponentsToSystemCalcDTO(el))
	}
	dto := SystemCalculationDTO{
		ID:                     orm.ID,
		SystemName:             orm.SystemName,
		AvailableCalculation:   orm.AvailableCalculation,
		Status:                 orm.Status,
		DateCreated:            orm.DateCreated,
		DateFormed:             orm.DateFormed,
		DateClosed:             orm.DateClosed,
		User:                   ToUserDTO(orm.User),
		Moderator:              &moderator,
		ComponentsToSystemCalc: componentsToSystemCalc,
	}
	return dto
}

func ToSystemCalculationListDTO(arr []ds.SystemCalculation) []SystemCalculationDTO {
	list := make([]SystemCalculationDTO, len(arr))
	for _, el := range arr {
		val := ToSystemCalculationDTO(el)
		list = append(list, val)
	}
	return list
}

type ComponentToSystemCalcByIdDTO struct {
	ComponentID         uint `json:"component_id"`
	SystemCalculationID uint `json:"system_calculation_id"`
}

type UpdateComponentToSystemCalcDTO struct {
	ComponentID         uint `json:"component_id"`
	SystemCalculationID uint `json:"system_calculation_id"`
	ReplicationCount    uint `json:"replication_count"`
}
