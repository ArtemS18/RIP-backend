package dto

import (
	"failiverCheck/internal/app/ds"
)

type UserDTO struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	IsModerator bool   `json:"is_moderator"`
}

type UserUpdateDTO struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
}

func ToUserDTO(orm ds.User) UserDTO {
	dto := UserDTO{
		ID:          orm.ID,
		Login:       orm.Login,
		IsModerator: orm.IsModerator,
	}
	return dto
}
