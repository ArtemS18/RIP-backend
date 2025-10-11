package postgres

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *Postgres) RegisterUser(credentials schemas.UserCredentials) (ds.User, error) {
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

func (r *Postgres) AuthUser(credentials schemas.UserCredentials) error {
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

func (r *Postgres) GetUserByLogin(login string) (ds.User, error) {
	var user ds.User
	err := r.db.Where("login = ?", login).First(&user).Error
	if err != nil {
		return ds.User{}, err
	}
	return user, nil

}

func (r *Postgres) LogoutUser(userId uint) error {
	return nil
}

func (r *Postgres) GetUserById(userId uint) (ds.User, error) {
	var user ds.User
	err := r.db.Find(&user, userId).Error
	if err != nil {
		return ds.User{}, err
	}
	return user, nil
}

func (r *Postgres) UpdateUserById(userId uint, update dto.UserUpdateDTO) (ds.User, error) {
	var user ds.User
	password := update.Password
	if password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return ds.User{}, err

		}
		hashedPasswordStr := string(hashedPassword)
		user.HashedPassword = hashedPasswordStr
	}
	if update.Login != nil {
		user.Login = *update.Login
	}
	res := r.db.Model(&ds.User{}).Where("id = ?", userId).Updates(user)
	if res.RowsAffected == 0 {
		return ds.User{}, gorm.ErrRecordNotFound
	}
	if res.Error != nil {
		return ds.User{}, res.Error
	}
	user, _ = r.GetUserById(userId)
	return user, nil
}
