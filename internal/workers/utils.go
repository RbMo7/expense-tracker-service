package workers

import (
	"encoding/json"
	"net/http"
)

func SendJSONError(w http.ResponseWriter, message string, errorType string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"message": message,
		"error":   errorType,
	})
}

func SendJSONSuccess(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]interface{}{
		"success": true,
		"message": message,
	}
	if data != nil {
		response["data"] = data
	}
	json.NewEncoder(w).Encode(response)
}
