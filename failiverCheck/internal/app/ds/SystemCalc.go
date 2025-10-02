package ds

import "time"

type SystemCalculation struct {
	ID                   uint       `gorm:"autoIncrement; primaryKey" json:"id"`
	SystemName           *string    `gorm:"size:256; defaul:null" json:"system_name"`
	AvailableCalculation float32    `gorm:"defaul:null" json:"available_calculation"`
	UserID               uint       `gorm:"not null" json:"-"`
	Status               string     `gorm:"size:256; default:DRAFT; not null" json:"status"`
	DateCreated          time.Time  `gorm:"autoCreateTime; not null" json:"date_created"`
	DateFormed           *time.Time `gorm:"default:null" json:"date_formed"`
	DateClosed           *time.Time `gorm:"default:null" json:"date_accepted"`
	ModeratorID          *uint      `json:"-"`

	User                   User                     `gorm:"foreignKey:UserID" json:"user"`
	Moderator              *User                    `gorm:"foreignKey:ModeratorID" json:"moderator"`
	ComponentsToSystemCalc []ComponentsToSystemCalc `gorm:"foreignKey:SystemCalculationID" json:"components"`
}

type Status string

const (
	DRAFT     Status = "DRAFT"
	DELETED   Status = "DELETED"
	COMPLETED Status = "COMPLETED"
	FORMED    Status = "FORMED"
	REJECTED  Status = "REJECTED"
)
