package methods

import (
	"encoding/json"
	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"net/http"
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	//json
	w.Header().Set("Content-Type", "application/json")

	//status of request
	w.WriteHeader(http.StatusOK)

	//adding changes to database
	var teams []essences.Team

	database.Db.Find(&teams)

	//writing response
	json.NewEncoder(w).Encode(&teams)
}
