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

func DeleteTeam(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var team models.Team
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID value", http.StatusBadRequest)
		return
	}

	result := database.Db.First(&team, id)
	//error if id not existing
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Team not found", http.StatusNotFound)
	} else {
		//delete and show response (deleted team)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		database.Db.Delete(&team)
		json.NewEncoder(w).Encode(&team)
	}
}
