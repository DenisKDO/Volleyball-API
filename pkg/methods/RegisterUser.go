package methods

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DenisKDO/Vollyball-API/internal/database"
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

	err := user.Validate()
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Fprintf(w, "Validation error for field %s:\n", err.Field())
			switch err.Tag() {
			case "required":
				fmt.Fprintf(w, "This field is required.")
			case "email":
				fmt.Fprintf(w, "Please enter a valid email address.\n")
			case "min":
				fmt.Fprintf(w, err.Field()+" must be at least 8 bytes long")
			default:
				fmt.Fprintf(w, "Unknown validation error: %s\n", err.Tag())
			}
			fmt.Fprintf(w, "")
		}
		return
	}
	user.PasswordOK.Set(user.Password)

	createdUser := database.Db.Create(&user)
	err = createdUser.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		//show response of the created team
		json.NewEncoder(w).Encode(&user)
	}

}
