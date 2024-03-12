package methods

import (
	"encoding/json"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var players []essences.Player
	database.Db.Find(&players)

	json.NewEncoder(w).Encode(&players)
}
