package validation

import (
	"net/http"
	"strconv"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

func StrToInt(queryStr string, w http.ResponseWriter, parameter string) int {
	queryInt, err := strconv.Atoi(queryStr)

	if err != nil {
		http.Error(w, "Invalid "+parameter, http.StatusBadRequest)
		return 0
	}
	return queryInt
}
