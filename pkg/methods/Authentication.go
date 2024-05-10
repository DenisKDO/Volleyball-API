package methods

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/tokens"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

// input struct
type input struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type minToken struct {
	Tokenok   string    `json:"token"`
	Expieryok time.Time `json:"expiery"`
}

func (i *input) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

func Authentication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input input
	var user essences.User
	var minToken minToken
	v := validation.New()

	//Taking json from user
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Error in json", http.StatusBadRequest)
	}

	//check for validation
	err = input.Validate()
	if err != nil {
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
		return
	}

	//check if there any user with that email
	result := database.Db.Where("email_address = ?", input.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Invalid user email or password", http.StatusNotFound)
		return
	}

	//Check if password valid
	if input.Password != *user.Password.Plaintext {
		http.Error(w, "Invalid user email or password", http.StatusBadRequest)
		return
	}

	//Creating a token for user
	token, err := tokens.GenerateToken(int64(user.ID), 24*time.Hour, tokens.ScopeAuthentication)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	minToken.Expieryok = token.Expiry
	minToken.Tokenok = token.Plaintext.Ptext
	//encode json token
	json.NewEncoder(w).Encode(&minToken)
}
