package dto

import (
	"failiverCheck/internal/app/ds"
)

type UserDTO struct {
	ID          uint   `json:"-"`
	Login       string `json:"login"`
	IsModerator bool   `json:"is_moderator"`
}

func ToUserDTO(orm ds.User) UserDTO {
	dto := UserDTO{
		ID:          orm.ID,
		Login:       orm.Login,
		IsModerator: orm.IsModerator,
	}
	return dto
}
