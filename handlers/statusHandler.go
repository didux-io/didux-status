package handlers

import (
	"encoding/json"
	"net/http"
	"smilo-status/managers"
)

// handlerFunction for / (root)
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := managers.GetStatus()

	json.NewEncoder(w).Encode(users)
}
