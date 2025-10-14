package schemas

type ComponentToSystemCalcByIdReq struct {
	ComponentID         uint `json:"component_id" validate:"required" binding:"required"`
	SystemCalculationID uint `json:"system_calculation_id" validate:"required" binding:"required"`
}

type UpdateComponentToSystemCalcReq struct {
	ComponentID         uint `json:"component_id" validate:"required" binding:"required"`
	SystemCalculationID uint `json:"system_calculation_id" validate:"required" binding:"required"`
	ReplicationCount    uint `json:"replication_count"`
}
