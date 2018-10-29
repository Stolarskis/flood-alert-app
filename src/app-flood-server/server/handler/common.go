package handler

import (
	"encoding/json"
	"net/http"
)

type user struct {
	name  string `json:"name"`
	email string `json:"email"`
	phone string `json:"phone"`
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, message string) {
	respondJSON(w, http.StatusOK, message)
}
