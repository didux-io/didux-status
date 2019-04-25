package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smilo-status/managers"
)

// handlerFunction for root URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Smilo Status!")
}

// handlerFunction for /users/ url path
func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := managers.GetStatus()

	json.NewEncoder(w).Encode(users)
}
