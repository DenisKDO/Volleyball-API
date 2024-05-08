package methods

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var player essences.Player
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID value", http.StatusBadRequest)
		return
	}

	result := database.Db.First(&player, id)
	//error if id not existing
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Player not found", http.StatusNotFound)
	} else {
		//delete and show response (deleted team)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		database.Db.Delete(&player)
		json.NewEncoder(w).Encode(&player)
	}
}
