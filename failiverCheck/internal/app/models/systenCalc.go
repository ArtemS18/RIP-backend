package models

type UpdateSystemCalcFields struct {
	SystemName *string `json:"system_name"`
}

type UpdateSystemCalcStatus struct {
	Command string `json:"command"`
}
