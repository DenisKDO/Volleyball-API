package methods

import (
	"encoding/json"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var teams []essences.Team

	database.Db.Find(&teams)

	jsonResponse, err := json.Marshal(teams)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
