package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegiserInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type formatterUserRegister struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	IsVerified  bool   `json:"is_verified"`
}

func FomatterUserRegister(user *User, token string) formatterUserRegister {
	resp := formatterUserRegister{
		ID:          user.ID,
		Name:        user.Email,
		Email:       user.Email,
		AccessToken: token,
		IsVerified:  user.IsVerified,
	}
	return resp
}

type formatterGetUser struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	IsVerified bool   `json:"is_verified"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func FormatterGetUser(users []*User) []formatterGetUser {
	var resp []formatterGetUser

	for _, user := range users {
		response := formatterGetUser{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			IsVerified: user.IsVerified,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		}
		resp = append(resp, response)
	}
	return resp
}

func FormatterDetailUser(user *User) formatterGetUser {
	response := formatterGetUser{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
	return response
}

type ParamsID struct {
	ID string `uri:"id" binding:"required"`
}

type UserUpdateInput struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	PasswordConfirma string `json:"password_confirm" binding:"required"`
}

func (user *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Model(User{}).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) UpdateUser(id uint, db *gorm.DB) error {
	err := db.Model(User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (user *User) DeleteUser(id uint, db *gorm.DB) error {
	err := db.Model(User{}).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindUserById(id uint, db *gorm.DB) (*User, error) {
	var user *User
	err := db.Model(User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) FindUserByEmail(email string, db *gorm.DB) (*User, error) {
	err := db.Model(User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) GetUser(db *gorm.DB) ([]*User, error) {
	listUser := []*User{}
	err := db.Model(User{}).Find(&listUser).Error
	if err != nil {
		return listUser, err
	}
	return listUser, nil
}
