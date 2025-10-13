package usecase

import (
	"context"
	"failiverCheck/internal/app/dto"
	utils "failiverCheck/internal/pkg/jwtUtils"
	"fmt"
	"time"
)

func (uc *UseCase) ValidateUser(ctx context.Context, token string) (dto.UserDTO, error) {
	claims, err := utils.ValidateJwtToken(token, uc.Config.JWT.SecretKey)
	if err != nil {
		return dto.UserDTO{}, fmt.Errorf("unauthorized: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := uc.Redis.GetBlackListJWT(ctx, token); err == nil {
		return dto.UserDTO{}, fmt.Errorf("unauthorized: token in blacklist")
	}
	userId, ok := claims["sub"].(float64)
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("bad jwt credentials")
	}
	is_moderator, ok := claims["is_moderator"].(bool)
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("bad jwt credentials")
	}
	login, ok := claims["login"].(string)
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("bad jwt credentials")
	}
	user := dto.UserDTO{
		ID:          uint(userId),
		IsModerator: is_moderator,
		Login:       login,
	}
	return user, nil
}
