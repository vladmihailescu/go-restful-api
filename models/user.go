package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string
}

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UserDTO) ToModel() User {
	return User{
		Name:  u.Name,
		Email: u.Email,
	}
}

type UserRegisterDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserRegisterDTO) ToModel() User {
	return User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserLoginDTO) ToModel() User {
	return User{
		Email:    u.Email,
		Password: u.Password,
	}
}
