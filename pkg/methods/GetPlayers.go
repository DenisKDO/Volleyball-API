package methods

import (
	"encoding/json"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	//json response and status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//getting all players from database
	var players []essences.Player
	database.Db.Find(&players)

	json.NewEncoder(w).Encode(&players)
}
