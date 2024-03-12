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

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var player essences.Player

	//error if not found player
	result := database.Db.First(&player, params["id"])
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	//update already existing player
	w.WriteHeader(http.StatusOK)
	json.NewDecoder(r.Body).Decode(&player)
	database.Db.Save(player)
	json.NewEncoder(w).Encode(&player)

}
