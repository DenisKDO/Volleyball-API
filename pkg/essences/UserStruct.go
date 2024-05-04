package essences

import (
	"github.com/DenisKDO/Vollyball-API/internal/password"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name         string            `json:"name" gorm:"not null" validate:"required"`
	EmailAddress string            `json:"emailAddress" gorm:"unique;not null" validate:"required,email"`
	PasswordOK   password.Password `json:"passwordOk" gorm:"not null" validate:"required"`
	Activated    bool              `json:"activated" gorm:"not null" validate:"required"`
	Password     string            `json:"password" gorm:"not null" validate:"required,min=8"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
