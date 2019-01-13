package common

import (
	"encoding/json"
	"net/http"
)

// Send JSON response based on dynamic response
func SendJSON(w http.ResponseWriter, r *http.Request, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Send common response include status and data
func CommonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commonResponse{
		Status: status,
		Data:   data,
	})
}
