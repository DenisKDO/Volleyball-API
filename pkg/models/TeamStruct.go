package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model

	Title       string   `json:"title" gorm:"unique" validate:"required"`
	Nickname    string   `json:"nicknames"`
	Association string   `json:"association"`
	Coach       []Coach  `json:"coach"`
	Players     []Player `json:"players"`
	FIVBRanking int      `json:"FIVBRanking"`
}

func (p *Team) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
