package ds

type User struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Login          string `gorm:"type:varchar(25);unique;not null" json:"login"`
	HashedPassword string `gorm:"type:varchar(100);not null" json:"-"`
	IsModerator    bool   `gorm:"type:boolean;default:false" json:"is_moderator"`
}
