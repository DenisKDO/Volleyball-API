package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model

	UniformNumber int    `json:"uniformNumber" gorm:"not null" validate:"required,numeric"`
	FirstName     string `json:"firstName" gorm:"not null" validate:"required"`
	SecondName    string `json:"secondName" gorm:"not null" validate:"required"`
	Position      string `json:"position"`
	DateOfBirth   string `json:"dateOfBirth"`
	Age           int    `json:"age" validate:"numeric"`
	Height        int    `json:"height" validate:"numeric"`
	Weight        int    `json:"weight" validate:"numeric"`
	SpikeHeight   int    `json:"spikeHeight" validate:"numeric"`
	BlockHeight   int    `json:"blockHeight" validate:"numeric"`
	TeamID        int    `json:"teamID" validate:"numeric"`
	TeamTitle     string `json:"teamName"`
}

func (p *Player) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
