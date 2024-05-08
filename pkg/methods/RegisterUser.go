package methods

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
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/go-playground/validator/v10"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	//json
	w.Header().Set("Content-Type", "application/json")

	//creating user struct
	var user essences.User
	//Take json of player from client
	json.NewDecoder(r.Body).Decode(&user)
	user.Activated = false

	//creating validation
	v := validation.New()

	err := user.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				v.AddError("Validation error for field "+err.Field(), "This field is requierd")
				w.WriteHeader(http.StatusBadRequest)
			case "email":
				v.AddError("Validation error for field "+err.Field(), "Please enter a valid email address")
				w.WriteHeader(http.StatusBadRequest)
			case "min":
				v.AddError("Validation error for field "+err.Field(), err.Field()+" must be at least 8 bytes long")
				w.WriteHeader(http.StatusBadRequest)
			default:
				v.AddError("Unknown", "unknown validation error")
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}

	//creating hash for our password
	pswrd := password.New()
	pswrd.Set(user.Password)
	user.PasswordHash = pswrd.Hash
	if !v.Valid() {
		for key, message := range v.Errors {
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUser := database.Db.Create(&user)

	//checking for unique of email
	err = createdUser.Error
	if err != nil {
		v.AddError("email", "email has already registered")
	}

	//generating toker
	token, err := tokens.GenerateToken(int64(user.ID), 3*time.Hour*24, tokens.ScopeActivation)
	if err != nil {
		v.AddError("Token", "Error in creating token")
	}

	if !v.Valid() {
		for key, message := range v.Errors {
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		return
	}

	createdToken := database.Db.Create(&token)
	err = createdToken.Error
	if err != nil {
		v.AddError("database", "error in creating token")
	}

	//email sending
	data := struct {
		Name  string
		ID    int
		Token string
	}{
		Name:  user.Name,
		ID:    int(user.ID),
		Token: token.Plaintext,
	}

	m := mailer.New("sandbox.smtp.mailtrap.io", 587, "5384e1ea4ca8b8", "6011b01060d126", "den.kim04@mail.ru")

	//sending and checking for panic in goroutine
	helper.Background(func() {
		err = m.Send(user.EmailAddress, "user_welcome.tmpl", data)
		if err != nil {
			// Importantly, if there is an error sending the email then we use the
			// http.Error() function to send a generic error message to the client.
			http.Error(w, "Error in sending email", http.StatusInternalServerError)
		}
	})

	//show response of the created user
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&user)

}
