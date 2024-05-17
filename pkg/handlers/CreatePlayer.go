package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/go-playground/validator/v10"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Please enter authorizations token", http.StatusForbidden)
		return
	}
	parts := strings.Split(authHeader, " ")
	tokenOK := parts[1]
	if parts[0] != "Bearer" || parts[1] == "" {
		http.Error(w, "Please enter authorizations token", http.StatusForbidden)
		return
	}
	hash := sha256.Sum256([]byte(tokenOK))
	var token models.Token
	token.Hash = hash[:]
	if err := database.Db.Where("hash = ?", token.Hash).First(&token).Error; err != nil {
		http.Error(w, "Invalid authorization token", http.StatusForbidden)
		return
	}
	//json
	w.Header().Set("Content-Type", "application/json")

	var player models.Player

	//Take json of player from client
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode JSON: %v ", err)
		return
	}

	//validation
	v := validation.New()

	err := player.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				v.AddError("Validation error for field "+err.Field(), "This field is requierd or invalid type of value")
			case "max":
				v.AddError("Validation error for field "+err.Field(), err.Field()+" must be max size - 2 bytes long")
			default:
				v.AddError("Unknown", "unknown validation error")
			}
		}
	}

	if !v.Valid() {
		for key, message := range v.Errors {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		return
	}
	database.Db.Create(&player)

	//response
	response := map[string]interface{}{
		"player": player,
	}

	//if ok give response of creative players
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
