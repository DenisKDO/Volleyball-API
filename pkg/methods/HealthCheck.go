package methods

import (
	"encoding/json"
	"net/http"
)

type SystemInfo struct {
	Environment string
	Version     string
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var SystemInfo SystemInfo
	SystemInfo.Environment = "development"
	SystemInfo.Version = "1.0.0"

	response := map[string]interface{}{
		"status":      "available",
		"system_info": SystemInfo,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
