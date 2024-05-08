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

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	//json response
	w.Header().Set("Content-Type", "application/json")

	//parameter from url
	params := mux.Vars(r)

	var player essences.Player

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID value", http.StatusBadRequest)
		return
	}

	//finding player by id else error 404
	result := database.Db.First(&player, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Player not found", http.StatusNotFound)
		return
	}

	//write response to client in json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&player)
}
