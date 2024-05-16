package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func UpdateCoach(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var coach models.Coach

	//error if not found player
	result := database.Db.First(&coach, params["id"])
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "-CoachRecrods: Not found", http.StatusNotFound)
		return
	}

	//update already existing player
	w.WriteHeader(http.StatusOK)
	json.NewDecoder(r.Body).Decode(&coach)
	database.Db.Save(coach)
	json.NewEncoder(w).Encode(&coach)

}
