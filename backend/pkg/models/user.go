package models

import (
	"backend/pkg/database"
	"backend/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username" validate:"required,min=3,max=20" gorm:"not null; unique"`
	Password  string `json:"password" validate:"required,min=8" gorm:"not null"`
	Mail      string `json:"mail" validate:"required,email" gorm:"not null; unique"`
}

func (user *UserModel) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *UserModel) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *UserModel) GetUserByUsername(username string) (*UserModel, error) {
	err := database.Database.
		Where("username = ?", username).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *UserModel) GetUserByMail(mail string) (*UserModel, error) {
	err := database.Database.
		Where("mail = ?", mail).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *UserModel) GetUserById(id string) (*UserModel, error) {
	err := database.Database.
		Where("id = ?", id).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *UserModel) GetUserByIdOrUsername(idOrUsername string) (*UserModel, error) {
	err := database.Database.
		Where("id = ? OR username = ?", idOrUsername, idOrUsername).
		First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *UserModel) CreateUser() error {
	err := database.Database.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *UserModel) DeleteUser() error {
	err := database.Database.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *UserModel) UpdateUser(newUser *UserModel) error {
	if newUser.Password != "" {
		err := newUser.HashPassword()
		if err != nil {
			return err
		}
	}

	utils.Replace(user, newUser, "Username", "Mail", "Password")

	err := database.Database.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
