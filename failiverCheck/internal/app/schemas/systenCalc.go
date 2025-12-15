package schemas

type UpdateSystemCalcFields struct {
	SystemName *string `json:"system_name" validate:"required"`
}

type UpdateSystemCalcAvailable struct {
	Available    *float32 `json:"available_calc" validate:"required"`
	Token        string   `json:"token" validate:"required"`
	SystemCalcID uint     `json:"sustem_calc_id" validate:"required"`
}

type UpdateSystemCalcStatus struct {
	Command string `json:"command" validate:"required"`
}
