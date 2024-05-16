package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/helper"
	"github.com/DenisKDO/Vollyball-API/internal/mailer"
	"github.com/DenisKDO/Vollyball-API/internal/password"
	"github.com/DenisKDO/Vollyball-API/internal/tokens"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/models"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type tempUser struct {
	gorm.Model

	Name         string `json:"name" gorm:"not null" validate:"required"`
	EmailAddress string `json:"emailAddress" gorm:"unique;not null" validate:"required,email"`
	Activated    bool   `json:"activated" gorm:"not null"`
	Password     string `json:"password" gorm:"not null" validate:"required,min=8"`
}

func (u *tempUser) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	//json
	w.Header().Set("Content-Type", "application/json")

	//creating user struct
	var user tempUser
	var trueUser models.User
	//Take JSON file from client
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to decode JSON: %v ", err)
		return
	}
	user.Activated = false

	//creating validation
	v := validation.New()

	err := user.Validate()
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

	//creating hash for our password
	pswrd := password.New()
	pswrd.Set(user.Password)
	trueUser.Password.Hash = pswrd.Hash
	trueUser.Password.Plaintext = &user.Password
	trueUser.Activated = false
	trueUser.EmailAddress = user.EmailAddress
	trueUser.Name = user.Name
	trueUser.PasswordHash = pswrd.Hash
	if !v.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		for key, message := range v.Errors {
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		return
	}

	if ok := database.Db.Where("email_address = ?", trueUser.EmailAddress).First(&trueUser).Error; ok == nil {
		http.Error(w, "email has already registered", http.StatusConflict)
		return
	}

	createdUser := database.Db.Create(&trueUser)

	//take error if user can't create
	err = createdUser.Error
	if err != nil {
		http.Error(w, "Error in creating user", http.StatusInternalServerError)
		return
	}

	//generating token
	token, err := tokens.GenerateToken(int64(trueUser.ID), 3*time.Hour*24, tokens.ScopeActivation)
	if err != nil {
		http.Error(w, "Error in creating token", http.StatusInternalServerError)
		return
	}

	createdToken := database.Db.Create(&token)
	err = createdToken.Error
	if err != nil {
		http.Error(w, "Error in creating token", http.StatusInternalServerError)
		return
	}

	//email sending
	data := struct {
		Name  string
		ID    int
		Token string
	}{
		Name:  user.Name,
		ID:    int(trueUser.ID),
		Token: token.Plaintext.Ptext,
	}

	m := mailer.New("sandbox.smtp.mailtrap.io", 587, "5384e1ea4ca8b8", "6011b01060d126", "den.kim04@mail.ru")

	//sending and checking for panic in goroutine
	helper.Background(func() {
		err = m.Send(user.EmailAddress, "user_welcome.tmpl", data)
		if err != nil {
			// Importantly, if there is an error sending the email then we use the
			// http.Error() function to send a generic error message to the client.
			http.Error(w, "Error in sending email", http.StatusInternalServerError)
			return
		}
	})

	//show response of the created user
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&trueUser)

}
