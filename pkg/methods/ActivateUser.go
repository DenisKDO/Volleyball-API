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
	"github.com/jinzhu/gorm"
)

func ActivateUser(w http.ResponseWriter, r *http.Request) {
	//json response
	w.Header().Set("Content-Type", "application/json")
	v := validation.New()

	var token essences.Token
	var user essences.User
	json.NewDecoder(r.Body).Decode(&token)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	if err := database.Db.Where("hash = ?", token.Hash).First(&token).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		v.AddError("Token", "token has not found")
	}

	result := database.Db.First(&user, token.UserID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		v.AddError("User", "user with that token has not found")
	}

	if time.Now().After(token.Expiry) {
		w.WriteHeader(http.StatusUnauthorized)
		v.AddError("Token", "Token is expired")
	}

	if !v.Valid() {
		for key, message := range v.Errors {
			fmt.Fprintf(w, "-%s: %s\n", key, message)
		}
		return
	}

	user.Activated = true
	database.Db.Delete(&token)
	database.Db.Save(user)
	json.NewEncoder(w).Encode(&user)
}
