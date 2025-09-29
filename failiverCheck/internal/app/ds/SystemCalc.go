package ds

import "time"

type SystemCalculation struct {
	ID                   uint      `gorm:"autoIncrement; primaryKey"`
	SystemName           string    `gorm:"size:256; defaul:null"`
	AvailableCalculation float32   `gorm:"defaul:null"`
	UserID               uint      `gorm:"not null"`
	Status               string    `gorm:"size:256; default:DRAFT; not null"`
	DateCreated          time.Time `gorm:"autoCreateTime; not null"`
	DateFormed           time.Time `gorm:"default:null"`
	DateAcceped          time.Time `gorm:"default:null"`
	ModeratorID          *uint

	User       User        `gorm:"foreignKey:UserID"`
	Moderator  User        `gorm:"foreignKey:ModeratorID"`
	Components []Component `gorm:"many2many:components_to_system_calcs"`
}

type Status string

const (
	DRAFT     Status = "DRAFT"
	DELETED   Status = "DELETED"
	COMPLETED Status = "COMPLETED"
	FORMED    Status = "FORMED"
	REJECTED  Status = "REJECTED"
)
