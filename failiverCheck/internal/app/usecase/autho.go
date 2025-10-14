package usecase

import (
	"context"
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"failiverCheck/internal/pkg/jwtUtils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) ValidateUser(ctx context.Context, token string) (dto.UserDTO, error) {
	claims, err := jwtUtils.ValidateJwtToken(token, uc.Config.JWT.SecretKey)
	if err != nil {
		return dto.UserDTO{}, fmt.Errorf("unauthorized: %w", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := uc.Redis.GetBlackListJWT(ctx, token); err == nil {
		return dto.UserDTO{}, fmt.Errorf("unauthorized: token in blacklist")
	}
	scope, ok := claims["scope"].(string)
	if !ok || scope != "access" {
		return dto.UserDTO{}, fmt.Errorf("bad jwt credentials")
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

func (uc *UseCase) Autho(credentials schemas.UserCredentials) (schemas.AuthoResp, error) {
	password := credentials.Password
	user, err := uc.Postgres.GetUserByLogin(credentials.Login)
	if err != nil {
		return schemas.AuthoResp{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return schemas.AuthoResp{}, fmt.Errorf("bad password credentials")
	}
	duration := time.Duration(uc.Config.JWT.ExpiresAtMinutes) * time.Minute
	exp := time.Now().Add(duration).Unix()
	tokenStr, err := createAccessToken(user, exp, uc.Config.JWT.SecretKey)
	if err != nil {
		return schemas.AuthoResp{}, err
	}
	dto := schemas.AuthoResp{
		TokenType:   "access",
		ExpiresIn:   int(exp),
		AccessToken: tokenStr,
	}
	return dto, nil
}

func createAccessToken(user ds.User, exp int64, secretKey string) (string, error) {
	claims := jwt.MapClaims{}
	logrus.Info(time.Unix(int64(exp), 0).Format(time.ANSIC))
	claims["sub"] = user.ID
	claims["exp"] = exp
	claims["login"] = user.Login
	claims["is_moderator"] = user.IsModerator
	claims["scope"] = "access"
	tokenStr, err := jwtUtils.CreateJwtToken(claims, secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// func createRefreshClaims(userID uint, tokenUID uuid.UUID, exp int64) jwt.Claims {
// 	claims := jwt.MapClaims{}
// 	claims["sub"] = userID
// 	claims["exp"] = exp
// 	claims["uuid"] = tokenUID
// 	claims["scope"] = "refresh"
// 	return claims
// }
