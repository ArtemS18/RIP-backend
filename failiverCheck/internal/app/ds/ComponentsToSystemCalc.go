package ds

type ComponentsToSystemCalc struct {
	ComponentID         uint              `gorm:"primaryKey; autoIncrement:false" json:"component_id"`
	SystemCalculationID uint              `gorm:"primaryKey; autoIncrement:false" json:"system_calculation_id"`
	ReplicationCount    uint              `gorm:"default:1" json:"replication_count"`
	Component           Component         `gorm:"foreignKey:ComponentID" json:"component_data"`
	SystemCalculation   SystemCalculation `gorm:"foreignKey:SystemCalculationID" json:"-"`
}
