package handler

import (
	"net/http"

	pb "github.com/Stolarskis/flood-alert-app/src/app-flood-api"
	"github.com/pkg/errors"
)

//Checks alerts for weather conditions ()
func CheckAlerts(w http.ResponseWriter, r *http.Request) error {
	clt, err := pb.CreateClient()
	if err != nil {
		return errors.Wrap(err, "Failed to create Client in app-server")
	}
	respondJSON(w, http.StatusOK, "Testing\n")
}

func Home(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "home")
}
