package methods

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DenisKDO/Vollyball-API/internal/database"
	"github.com/DenisKDO/Vollyball-API/internal/validation"
	"github.com/DenisKDO/Vollyball-API/pkg/essences"
	"github.com/go-playground/validator/v10"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	//json
	w.Header().Set("Content-Type", "application/json")

	var player essences.Player

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

	//if ok give response of creative players
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&player)
}
