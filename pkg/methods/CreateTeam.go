package methods

import (
	"encoding/json"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var team essences.Team

	//Take JSON file from client
	json.NewDecoder(r.Body).Decode(&team)
	createdTeam := database.Db.Create(&team)
	err := createdTeam.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&team)
	}
}
