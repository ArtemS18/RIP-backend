package repository

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/schemas"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *Repository) RegisterUser(credentials schemas.UserCredentials) (ds.User, error) {
	password := credentials.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ds.User{}, err
	}
	newUser := ds.User{
		Login:          credentials.Login,
		HashedPassword: string(hashedPassword),
		IsModerator:    false,
	}
	err = r.db.Save(&newUser).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return ds.User{}, fmt.Errorf("this user alredy exists")
		}
		return ds.User{}, fmt.Errorf("this user alredy exists")
	}
	return newUser, nil

}

func (r *Repository) AuthUser(credentials schemas.UserCredentials) error {
	var user ds.User
	password := credentials.Password
	err := r.db.Where("login = ?", credentials.Login).First(&user).Error
	if err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("bad password credentials")
	}
	return nil

}

func (r *Repository) LogoutUser(userId uint) error {
	return nil
}

func (r *Repository) GetUserById(userId uint) (ds.User, error) {
	var user ds.User
	err := r.db.Find(&user, userId).Error
	if err != nil {
		return ds.User{}, err
	}
	return user, nil
}
