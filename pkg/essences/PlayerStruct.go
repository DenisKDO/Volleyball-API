package essences

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model

	UniformNumber int    `json:"uniformNumber" gorm:"not null" validate:"required,max=3"`
	FirstName     string `json:"firstName" gorm:"not null" validate:"required"`
	SecondName    string `json:"secondName" gorm:"not null" validate:"required"`
	Position      string `json:"position"`
	DateOfBirth   string `json:"dateOfBirth"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	SpikeHeight   int    `json:"spikeHeight"`
	BlockHeight   int    `json:"blockHeight"`
	TeamID        int    `json:"teamID"`
	TeamTitle     string `json:"teamName"`
}

func (p *Player) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
