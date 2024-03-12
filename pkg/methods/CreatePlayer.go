package methods

import (
	"encoding/json"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	//json
	w.Header().Set("Content-Type", "application/json")
	//status 200
	w.WriteHeader(http.StatusOK)

	var player essences.Player

	//Take json of player from client
	json.NewDecoder(r.Body).Decode(&player)
	//Create player to database
	createdPlayer := database.Db.Create(&player)
	err := createdPlayer.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		//show response of the created team
		json.NewEncoder(w).Encode(&player)
	}
}
