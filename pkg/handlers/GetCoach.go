package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetCoach(w http.ResponseWriter, r *http.Request) {
	//json response
	w.Header().Set("Content-Type", "application/json")

	//parameter from url
	params := mux.Vars(r)

	var coach models.Coach

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "-ID: Invalid value", http.StatusBadRequest)
		return
	}

	//finding player by id else error 404
	result := database.Db.First(&coach, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "-CoachRecord: Not found", http.StatusNotFound)
		return
	}

	//write response to client in json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&coach)
}
