package methods

import (
	"encoding/json"
	"errors"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetTeam(w http.ResponseWriter, r *http.Request) {

	//params from URL
	params := mux.Vars(r)

	var team essences.Team
	var players []essences.Player

	//check if team with id in url exist
	database.Db.First(&team, params["id"])
	if errors.Is(database.Db.First(&team, params["id"]).Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Team not found", http.StatusNotFound)
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
