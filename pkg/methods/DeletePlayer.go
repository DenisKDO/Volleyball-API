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

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var player essences.Player

	result := database.Db.First(&player, params["id"])
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
