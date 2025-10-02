package dto

type UserDTO struct {
	ID          uint   `json:"-"`
	Login       string `json:"login"`
	IsModerator bool   `json:"is_moderator"`
}
