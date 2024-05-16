package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Coach struct {
	gorm.Model

	FirstName   string `json:"firstName" gorm:"not null" validate:"required"`
	SecondName  string `json:"secondName" gorm:"not null" validate:"required"`
	Position    string `json:"position"`
	DateOfBirth string `json:"dateOfBirth"`
	Age         int    `json:"age" validate:"numeric"`
	TeamID      int    `json:"teamID" validate:"numeric"`
}

func (c *Coach) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
