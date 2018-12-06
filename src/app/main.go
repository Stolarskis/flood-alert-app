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
	fmt.Println("Check alerts called")
	err := CheckFloodAlerts()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.CheckAlertsResponse{Output: "Checked for Alerts"}, nil
}

func (s *server) TestAlerts(ctx context.Context, r *pb.TestAlertsRequest) (*pb.TestAlertsResponse, error) {
	fmt.Println("test Alerts called")
	sendEmailAlert(r.TestMessage)
	sendSMSAlert(r.TestMessage)
	return &pb.TestAlertsResponse{Output: "test messages sent"}, nil
}

//Mutes/unmutes alerts -
func (s *server) MuteAlerts(ctx context.Context, r *pb.MuteAlertRequest) (*pb.MuteAlertResponse, error) {

	fmt.Println("Mute Alerts called")
	if r.MuteAlertType == "SMS" {
		smsMute = !smsMute
		if smsMute {
			return &pb.MuteAlertResponse{Output: "Muted SMS Alerts"}, nil
		}
		return &pb.MuteAlertResponse{Output: "Unmuted SMS Alerts"}, nil
	}
	if r.MuteAlertType == "Email" {
		emailMute = !emailMute
		if emailMute {
			return &pb.MuteAlertResponse{Output: "Muted Email Alerts"}, nil
		}
		return &pb.MuteAlertResponse{Output: "Unmuted Email Alerts"}, nil
	}
	return nil, errors.New("Wrong type entered, can only be either SMS or Email")
}
