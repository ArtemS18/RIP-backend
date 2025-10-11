package usecase

import (
	"failiverCheck/internal/app/schemas"
	utils "failiverCheck/internal/pkg/jwt"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Autho(credentials schemas.UserCredentials) (schemas.AuthoResp, error) {
	password := credentials.Password
	user, err := uc.Postgres.GetUserByLogin(credentials.Login)
	if err != nil {
		return schemas.AuthoResp{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return schemas.AuthoResp{}, fmt.Errorf("bad password credentials")
	}
	claims := jwt.MapClaims{}
	duration := time.Duration(uc.Config.JWT.ExpiresAtMinutes) * time.Minute
	exp := time.Now().Add(duration).Unix()
	logrus.Info(time.Unix(int64(exp), 0).Format(time.ANSIC))
	claims["sub"] = user.ID
	claims["exp"] = exp
	claims["login"] = user.Login
	claims["is_moderator"] = user.IsModerator
	claims["token_type"] = "access"

	tokenStr, err := utils.CreateJwtToken(claims, uc.Config.JWT.SecretKey)
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
