package ds

type Component struct {
	ID          uint32  `gorm:"autoIncrement; primaryKey" json:"id"`
	Title       string  `gorm:"size:256" json:"title"`
	Type        string  `gorm:"size:256" json:"type"`
	MTBF        uint32  `gorm:"type:integer" json:"mtbf"`
	MTTR        uint32  `gorm:"type:integer" json:"mttr"`
	Available   float32 `json:"available"`
	Img         string  `gorm:"size:512; default: null" json:"img"`
	Description string  `json:"description"`
	IsDeleted   bool    `gorm:"default:false" json:"-"`
}
