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
	"github.com/jinzhu/gorm"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
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
	//json
	w.Header().Set("Content-Type", "application/json")

	var team models.Team

	//Take JSON file from client
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode JSON: %v ", err)
		return
	}

	//validation
	v := validation.New()

	err := team.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				v.AddError("Validation error for field "+err.Field(), "This field is requierd or invalid type of value")
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

	//check for unique
	if err := database.Db.Where("title = ?", team.Title).First(&models.Team{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Failed to check uniqueness of team title: %v\n", err)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Team with title '%s' already exists\n", team.Title)
		return
	}

	database.Db.Create(&team)

	//status 200
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&team)

}
