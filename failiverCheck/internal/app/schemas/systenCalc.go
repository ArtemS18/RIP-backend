package schemas

type UpdateSystemCalcFields struct {
	SystemName *string `json:"system_name" validate:"required"`
}

type UpdateSystemCalcStatus struct {
	Command string `json:"command" validate:"required"`
}
