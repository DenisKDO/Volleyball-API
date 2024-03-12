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

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var team essences.Team

	result := database.Db.First(&team, params["id"])
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Team not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewDecoder(r.Body).Decode(&team)
	database.Db.Save(team)
	json.NewEncoder(w).Encode(&team)
}
