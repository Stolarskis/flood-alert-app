package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Stolarskis/flood-alert-app/src/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFloodAlertAppServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) CheckAlerts(ctx context.Context, r *pb.CheckAlertsRequest) (*pb.CheckAlertsResponse, error) {
	err := CheckFloodAlerts()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.CheckAlertsResponse{Output: "Checked for Alerts"}, nil
}

func (s *server) TestAlerts(ctx context.Context, r *pb.TestAlertsRequest) (*pb.TestAlertsResponse, error) {
	sendEmailAlert(r.TestMessage)
	sendSMSAlert(r.TestMessage)
	return &pb.TestAlertsResponse{Output: "test messages sent"}, nil
}

//Mutes/unmutes alerts -
func (s *server) MuteAlerts(ctx context.Context, r *pb.MuteAlertRequest) (*pb.MuteAlertResponse, error) {

	if r.MuteAlertType == "sms" {
		smsMute = r.MuteAlertStatus
		if r.MuteAlertStatus {
			return &pb.MuteAlertResponse{Output: "Muted SMS Alerts"}, nil
		}
		return &pb.MuteAlertResponse{Output: "Unmuted SMS Alerts"}, nil
	}
	if r.MuteAlertType == "email" {
		emailMute = !emailMute
		if r.MuteAlertStatus {
			return &pb.MuteAlertResponse{Output: "Muted Email Alerts"}, nil
		}
		return &pb.MuteAlertResponse{Output: "Unmuted Email Alerts"}, nil
	}
	return nil, errors.New("Wrong type entered, can only be either SMS or Email, not case sensitive")
}

func (s *server) GetForcast(ctx context.Context, r *pb.GetForcastRequest) (*pb.GetForcastResponse, error) {
	forecast, err := getCurrentForcast()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get forcast")
	}

	forcastString := fmt.Sprintf("Summary:  %s, Temp: %v, Change of Rain: %v", forecast.Currently.Summary, int(forecast.Currently.Temperature), int(forecast.Currently.PrecipProbability))

	return &pb.GetForcastResponse{Forcast: forcastString}, nil
}

func (s *server) UpdateForcast(ctx context.Context, r *pb.UpdateForcastRequest) (*pb.UpdateForcastResponse, error) {
	err := updateForcast()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update forcast")
	}
	return &pb.UpdateForcastResponse{UpdatedForcast: ""}, nil
}
