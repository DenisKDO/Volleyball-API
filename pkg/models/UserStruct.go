package models

import (
	"github.com/DenisKDO/Vollyball-API/internal/password"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name         string            `json:"name" gorm:"not null" validate:"required"`
	EmailAddress string            `json:"emailAddress" gorm:"unique;not null" validate:"required,email"`
	Activated    bool              `json:"activated" gorm:"not null"`
	Password     password.Password `json:"-"`
	PasswordHash []byte            `json:"passwordHash" gorm:"not null"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
