package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/Stolarskis/flood-alert-app/src/app-flood-api"
	"github.com/pkg/errors"
)

//CheckAlerts checks for weather conditions
func CheckAlerts(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	clt, err := pb.CreateClient()
	if err != nil {
		respondError(w, "Failed to create Client in app-server")
		return errors.Wrap(err, "Failed to create Client in app-server")
	}
	req := &pb.CheckAlertsRequest{
		AlertMode: 0,
	}
	resp, err := clt.CheckAlerts(ctx, req)
	if err != nil {
		respondError(w, "Failed to check alerts")
		return errors.Wrap(err, "Failed to check alerts")
	}
	respondJSON(w, http.StatusOK, resp.Output)
	return nil
}

//Home checks alerts as well
func Home(w http.ResponseWriter, r *http.Request) {
	CheckAlerts(w, r)
}

//AddUser adds a user to the datastore instance
func AddUser(w http.ResponseWriter, r *http.Request) {
	user := new(user)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("failed to decode", err)
	}
	fmt.Println("printing user", user)

}
