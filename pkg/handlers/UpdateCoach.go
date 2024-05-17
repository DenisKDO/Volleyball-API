package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func UpdateCoach(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var coach models.Coach

	//error if not found player
	result := database.Db.First(&coach, params["id"])
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "-CoachRecrods: Not found", http.StatusNotFound)
		return
	}

	//update already existing player
	w.WriteHeader(http.StatusOK)
	json.NewDecoder(r.Body).Decode(&coach)
	database.Db.Save(coach)
	json.NewEncoder(w).Encode(&coach)

}
