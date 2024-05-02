package helper

import (
	"math"
	"net/http"
	"strconv"

	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/jinzhu/gorm"
)

func StrToInt(queryStr string, w http.ResponseWriter, parameter string) int {
	queryInt, err := strconv.Atoi(queryStr)

	if err != nil {
		http.Error(w, "Invalid "+parameter, http.StatusBadRequest)
		return 0
	}
	return queryInt
}

func NoRecordsFind(db *gorm.DB, w http.ResponseWriter, parameter string) int {
	var count int
	db.Model(&essences.Player{}).Count(&count)
	if count == 0 {
		http.Error(w, "No records found matching the "+parameter, http.StatusNotFound)
		return count
	}
	return 1
}

func RoundUp(x float64) int {
	return int(math.Ceil(x))
}
