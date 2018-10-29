package handler

import (
	"context"
	"net/http"

	pb "github.com/Stolarskis/flood-alert-app/src/app-flood-api"
	"github.com/pkg/errors"
)

//CheckAlerts checks for weather conditions
func CheckAlerts(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	clt, err := pb.CreateClient()
	if err != nil {
		respondError(w, 0, "Failed to create GRPC client")
		return errors.Wrap(err, "Failed to create Client in app-server")
	}
	req := &pb.CheckAlertsRequest{
		AlertMode: 0,
	}
	resp, err := clt.CheckAlerts(ctx, req)
	if err != nil {
		respondError(w, 1, "Failed to check weather alerts")
		return errors.Wrap(err, "Failed to check alerts")
	}
	respondJSON(w, http.StatusOK, resp.Output)
	return nil
}

//Home checks alerts as well
func Home(w http.ResponseWriter, r *http.Request) {
	CheckAlerts(w, r)
}
