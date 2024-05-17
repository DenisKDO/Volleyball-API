package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func DeleteCoach(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Please enter authorizations token", http.StatusForbidden)
		return
	}
	parts := strings.Split(authHeader, " ")
	tokenOK := parts[1]
	hash := sha256.Sum256([]byte(tokenOK))
	var token models.Token
	token.Hash = hash[:]
	if err := database.Db.Where("hash = ?", token.Hash).First(&token).Error; err != nil {
		http.Error(w, "Invalid authorization token", http.StatusForbidden)
		return
	}
	params := mux.Vars(r)

	var coach models.Coach
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "-ID: Invalid value", http.StatusBadRequest)
		return
	}

	result := database.Db.First(&coach, id)
	//error if id not existing
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "-CoachRecords: Not found", http.StatusNotFound)
	} else {
		//delete and show response (deleted team)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		database.Db.Delete(&coach)
		json.NewEncoder(w).Encode(&coach)
	}
}
