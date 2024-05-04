package helper

import (
	"math"
	"net/http"

	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/jinzhu/gorm"
)

func NoRecordsFind(db *gorm.DB, w http.ResponseWriter, parameter string) int {
	var count int
	db.Model(&essences.Player{}).Count(&count)
	if count == 0 {
		return count
	}
	return 1
}

func RoundUp(x float64) int {
	return int(math.Ceil(x))
}
