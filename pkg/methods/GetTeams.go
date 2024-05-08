package methods

import (
	"encoding/json"
	"net/http"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	//json
	w.Header().Set("Content-Type", "application/json")

	//status of request
	w.WriteHeader(http.StatusOK)

	//adding changes to database
	var teams []essences.Team
	var players []essences.Player
	database.Db.Find(&teams)
	for index := range teams {

		database.Db.Model(&teams[index]).Related(&players)
		teams[index].Players = players
	}
	//writing response
	json.NewEncoder(w).Encode(&teams)
}
