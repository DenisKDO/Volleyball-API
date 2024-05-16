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

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var player models.Player
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "-ID: Invalid value", http.StatusBadRequest)
		return
	}

	result := database.Db.First(&player, id)
	//error if id not existing
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "-PlayerRecords: Not found", http.StatusNotFound)
	} else {
		//delete and show response (deleted team)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		database.Db.Delete(&player)
		json.NewEncoder(w).Encode(&player)
	}
}
