package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/tokens"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// input struct
type input struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (i *input) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input input
	var user models.User
	v := validation.New()

	//take json from user
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode JSON: %v ", err)
		return
	}

	//check for validation
	if err := input.Validate(); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				v.AddError("Validation error for field "+err.Field(), "This field is requierd")
			case "email":
				v.AddError("Validation error for field "+err.Field(), "Please enter a valid email address")
			case "min":
				v.AddError("Validation error for field "+err.Field(), err.Field()+" must be at least 8 bytes long")
			default:
				v.AddError("Unknown", "unknown validation error")
			}
		}
	}

	if !v.Valid() {
		for key, message := range v.Errors {
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check if there any user with that email
	result := database.Db.Where("email_address = ?", input.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Invalid user email or password", http.StatusUnauthorized)
		return
	}

	//if not activated0
	if !user.Activated {
		http.Error(w, "Your profile has not activated", http.StatusForbidden)
		return
	}

	//check for password
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(input.Password)); err != nil {
		http.Error(w, "Invalid user email or password", http.StatusUnauthorized)
		return
	}

	//create token for user
	token, err := tokens.GenerateToken(int64(user.ID), 24*time.Hour, tokens.ScopeAuthentication)
	if err != nil {
		http.Error(w, "Server error in creating authentication token", http.StatusInternalServerError)
		return
	}

	//if user already have not expired token
	var tokenCheck models.Token
	if ok := database.Db.Where("user_id = ?", user.ID).First(&tokenCheck).Error; ok == nil {
		if token.Scope == tokens.ScopeAuthentication {
			http.Error(w, "This user already has authentication token", http.StatusConflict)
			return
		}
	}

	//create token
	database.Db.Create(&token)

	//json response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"authentication_token": map[string]interface{}{
			"token":  token.Plaintext,
			"expiry": token.Expiry,
		},
	})
}
