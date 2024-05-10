package methods

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type MinToken struct {
	Token string `json:"token" validate:"required"`
}

func (u *MinToken) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	//json response
	w.Header().Set("Content-Type", "application/json")

	v := validation.New()

	var minToken MinToken
	var token essences.Token
	var user essences.User

	//take json from user
	if err := json.NewDecoder(r.Body).Decode(&minToken); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode JSON: %v ", err)
		return
	}

	//validation
	err := minToken.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				v.AddError("Validation error for field "+err.Field(), "This field is requierd")
			}
		}
	}
	if !v.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		for key, message := range v.Errors {
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		return
	}

	hash := sha256.Sum256([]byte(minToken.Token))
	token.Hash = hash[:]
	if err := database.Db.Where("hash = ?", token.Hash).First(&token).Error; err != nil {
		http.Error(w, "Token has expired, has already been used, or does not exist", http.StatusNotFound)
		return
	}

	result := database.Db.First(&user, token.UserID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "User with that token has not found", http.StatusNotFound)
		return
	}

	if time.Now().After(token.Expiry) {
		http.Error(w, "Token is expired", http.StatusUnauthorized)
		return
	}

	user.Activated = true
	fmt.Println(token)
	database.Db.Delete(&token)
	database.Db.Save(user)
	json.NewEncoder(w).Encode(&user)
}
