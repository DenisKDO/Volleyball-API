package essences

import "github.com/jinzhu/gorm"

type Player struct {
	gorm.Model

	UniformNumber int    `json:"uniformNumber"`
	FirstName     string `json:"firstName"`
	SecondName    string `json:"secondName"`
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
