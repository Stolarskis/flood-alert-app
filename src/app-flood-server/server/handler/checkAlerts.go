package handler

import (
	"net/http"

	pb "github.com/Stolarskis/flood-alert-app/src/app-flood-api"
)

func CheckAlerts(w http.ResponseWriter, r *http.Request) {

	pb.
		respondJSON(w, http.StatusOK, "Testing\n")
}

func Home(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "HomePage")
}
