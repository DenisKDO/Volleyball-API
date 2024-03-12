package essences

import "github.com/jinzhu/gorm"

type Team struct {
	gorm.Model

	Title       string   `json:"title" gorm:"unique"`
	Nickname    string   `json:"nicknames"`
	Association string   `json:"association"`
	Coach       string   `json:"coach"`
	Players     []Player `json:"players"`
	FIVBRanking int      `json:"FIVBRanking"`
}
