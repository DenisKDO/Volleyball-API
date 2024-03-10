package pkg

import "github.com/jinzhu/gorm"

type Player struct {
	gorm.Model

	FirstName   string
	SecondName  string
	TeamID      int
	DateOfBirth string
	Age         int
	SpikeHeight int
	BlockHeight int
	Height      int
}

type Team struct {
	gorm.Model

	Title   string
	Coach   string
	Players []Player
}
