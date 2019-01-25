package common

import (
	"fmt"
	"strconv"
	"net/http"
)

func SendJSONError(w http.ResponseWriter, r *http.Request, message string, code int) {
	js := fmt.Sprintf(`{
		"status": "OK",
		"message": %s,
		"code": %d
	}`, strconv.Quote(message), code)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(js))
}

func SendJSON(w http.ResponseWriter, r *http.Request, res []byte, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(res))
}
