package common

import (
  "encoding/json"
  "net/http"
)

func SendJSON(w http.ResponseWriter, r *http.Request, res interface{}) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
}