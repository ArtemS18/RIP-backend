package usecase

import (
	"context"
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"fmt"
	"slices"
	"time"
)

func (uc *UseCase) RegisterUser(credentials schemas.UserCredentials) (ds.User, error) {
	user, err := uc.Postgres.RegisterUser(credentials)
	if err != nil {
		return ds.User{}, err
	}
	return user, nil

}

func (uc *UseCase) LogoutUser(ctx context.Context, token string) error {
	err := uc.Redis.SetBlackListJWT(ctx, token, time.Duration(uc.Config.JWT.ExpiresAtMinutes)*time.Minute)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) GetUser(userId uint) (dto.UserDTO, error) {
	user, err := uc.Postgres.GetUserById(userId)
	if err != nil {
		return dto.UserDTO{}, err
	}
	userDTO := dto.ToUserDTO(user)
	return userDTO, nil
}

func (uc *UseCase) UpdateUser(userId uint, update dto.UserUpdateDTO) (dto.UserDTO, error) {
	user, err := uc.Postgres.UpdateUserById(userId, update)
	if err != nil {
		return dto.UserDTO{}, err
	}
	userDTO := dto.ToUserDTO(user)
	return userDTO, nil

}

func (uc *UseCase) ValidateRole(user dto.UserDTO, allowedRoles ...schemas.Role) error {
	var role schemas.Role
	if user.IsModerator {
		role = schemas.ModeratorRole
	}
	if !slices.Contains(allowedRoles, schemas.ModeratorRole) {
		allowedRoles = append(allowedRoles, schemas.ModeratorRole)
	}

	if !slices.Contains(allowedRoles, role) {
		return fmt.Errorf("not allowed role")
	}
	return nil
}
