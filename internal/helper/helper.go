package helper

import (
	"fmt"
	"math"
	"net/http"

	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/jinzhu/gorm"
)

var (
	Status bool
)

func NoRecordsFind(db *gorm.DB, w http.ResponseWriter, parameter string) int {
	var count int
	db.Model(&models.Player{}).Count(&count)
	if count == 0 {
		return count
	}
	return 1
}

func NoRecordsFindCoach(db *gorm.DB, w http.ResponseWriter, parameter string) int {
	var count int
	db.Model(&models.Coach{}).Count(&count)
	if count == 0 {
		return count
	}
	return 1
}

func NoRecordsFindTeam(db *gorm.DB, w http.ResponseWriter, parameter string) int {
	var count int
	db.Model(&models.Team{}).Count(&count)
	if count == 0 {
		return count
	}
	return 1
}

func RoundUp(x float64) int {
	return int(math.Ceil(x))
}

func Background(fn func()) {
	go func() {
		// Recover any panic
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Panic recovered in goroutine:", err)
			}
		}()
		// Execute the arbitrary function that we passed as the parameter.
		fn()
	}()
}
