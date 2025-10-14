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

type SearchSystemCalcDTO struct {
	DateFormedStart *string    `form:"date_formed_start"`
	DateFormedEnd   *string    `form:"date_formed_end"`
	Status          *ds.Status `form:"status"`
	UserID          *uint      `form:"-"`
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

	User                   string                      `json:"user"`
	Moderator              *string                     `json:"moderator"`
	ComponentsToSystemCalc []ComponentsToSystemCalcDTO `json:"components"`
}

type SystemCalculationInfoDTO struct {
	ID                   uint       `json:"id"`
	SystemName           *string    `json:"system_name"`
	AvailableCalculation float32    `json:"available_calculation"`
	Status               string     `json:"status"`
	DateCreated          time.Time  `json:"date_created"`
	DateFormed           *time.Time `json:"date_formed"`
	DateClosed           *time.Time `json:"date_accepted"`

	User      string  `json:"user"`
	Moderator *string `json:"moderator"`
}

func ToSystemCalculationInfoDTO(orm ds.SystemCalculation) SystemCalculationInfoDTO {
	var moderator *string = nil
	if orm.Moderator != nil {
		ptr := orm.Moderator.Login
		moderator = &ptr
	}
	dto := SystemCalculationInfoDTO{
		ID:                   orm.ID,
		SystemName:           orm.SystemName,
		AvailableCalculation: orm.AvailableCalculation,
		Status:               orm.Status,
		DateCreated:          orm.DateCreated,
		DateFormed:           orm.DateFormed,
		DateClosed:           orm.DateClosed,
		User:                 orm.User.Login,
		Moderator:            moderator,
	}
	return dto
}

func ToSystemCalculationDTO(orm ds.SystemCalculation) SystemCalculationDTO {
	var moderator *string = nil
	if orm.Moderator != nil {
		ptr := orm.Moderator.Login
		moderator = &ptr
	}
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
		User:                   orm.User.Login,
		Moderator:              moderator,
		ComponentsToSystemCalc: componentsToSystemCalc,
	}
	return dto
}

func ToSystemCalculationListDTO(arr []ds.SystemCalculation) []SystemCalculationDTO {
	list := make([]SystemCalculationDTO, len(arr))
	for i, el := range arr {
		val := ToSystemCalculationDTO(el)
		list[i] = val
	}
	return list
}

func ToSystemCalculationInfoListDTO(arr []ds.SystemCalculation) []SystemCalculationInfoDTO {
	list := make([]SystemCalculationInfoDTO, len(arr))
	for i, el := range arr {
		val := ToSystemCalculationInfoDTO(el)
		list[i] = val
	}
	return list
}

type ComponentToSystemCalcByIdDTO struct {
	ComponentID         uint
	SystemCalculationID uint
	UserID              uint
}

type UpdateComponentToSystemCalcDTO struct {
	ComponentID         uint
	SystemCalculationID uint
	ReplicationCount    uint
	UserID              uint
}
