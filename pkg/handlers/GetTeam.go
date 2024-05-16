package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetTeam(w http.ResponseWriter, r *http.Request) {

	//params from URL
	params := mux.Vars(r)

	var team models.Team
	var players []models.Player

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID value", http.StatusBadRequest)
		return
	}

	//check if team with id in url exist
	database.Db.First(&team, id)
	if errors.Is(database.Db.First(&team, params["id"]).Error, gorm.ErrRecordNotFound) {
		http.Error(w, "-Team: Not found", http.StatusNotFound)
		return
	}

	//make relation between team and players
	database.Db.Model(&team).Related(&players)

	team.Players = players

	//if found return to client json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(team)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
