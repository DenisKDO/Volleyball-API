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

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var team models.Team

	//finding team that we want to update else error 404
	result := database.Db.First(&team, params["id"])
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	//update already existing team
	w.WriteHeader(http.StatusOK)
	json.NewDecoder(r.Body).Decode(&team)
	database.Db.Save(team)
	json.NewEncoder(w).Encode(&team)
}
