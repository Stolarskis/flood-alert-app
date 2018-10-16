package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Stolarskis/flood-alert-app/app-flood-api"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "check-alerts",
			Aliases: []string{"c"},
			Usage:   "Main service call: Calls the main process of flood-app. Which gets info from openweather and then checks for conditions to alert user",
			Action: func(c *cli.Context) error {
				checkInfo()
				return nil
			},
		},

		{
			Name:    "test-alerts",
			Aliases: []string{"t"},
			Usage:   "Sends out test alerts of methods flood-app uses",
			Action: func(c *cli.Context) error {
				testAlerts()
				return nil
			},
		},
		{
			Name:    "mute-alert",
			Aliases: []string{"m"},
			Usage:   "Mutes an alert based on priority. 1 2 or 3",
			Action: func(c *cli.Context) error {
				muteAlert(c)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createClient() (pb.FloodAlertAppServiceClient, error) {
	ip := getIPEnv()
	fmt.Println(ip)
	conn, err := grpc.Dial(ip+":3000", grpc.WithInsecure())
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("Unable to create client")
	}
	return pb.NewFloodAlertAppServiceClient(conn), nil
}

func checkInfo() {
	client, err := createClient()
	if err != nil {
		log.Fatalf("Unable to create grpc client")
	}
	fmt.Println("client created")

	checkAlertsRequest := &pb.CheckAlertsRequest{
		AlertMode: 0,
	}
	ctx := context.Background()
	checkAlertsResponse, err := client.CheckAlerts(ctx, checkAlertsRequest)
	if err != nil {
		fmt.Println("Unable to Check Alerts", err)
	}
	fmt.Println(checkAlertsResponse.Output)
}

// Sends
func testAlerts() error {
	client, err := createClient()
	if err != nil {
		return errors.  ("Failed to create GRPC Client")
	}
	testAlertsRequest := &pb.TestAlertsRequest{TestMessage: "Test Message"}
	ctx := context.Background()
	testAlertsResponse, err := client.TestAlerts(ctx, testAlertsRequest)
	if err != nil {
		fmt.Println("Unable to send Test Alerts", err)
	}
	fmt.Println(testAlertsResponse)
}

func muteAlert(c *cli.Context) {
	client, err := createClient()
	if err != nil {
		log.Fatalf("Unable to create grpc client")
	}
	ctx := context.Background()
	priority, err := strconv.ParseUint(c.Args().Get(0), 10, 64)
	fmt.Println(priority)
	if err != nil {
		log.Fatalf("Unable to convert input, Invalid Input")
	}

	if priority != 1 && priority != 2 && priority != 3 {
		log.Fatalf("Invalid input")
	}

	muteAlertResponse, err := client.MuteAlerts(ctx, &pb.MuteAlertRequest{MuteAlertPriority: priority})
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Unable to send message to app")
	}
	fmt.Println(muteAlertResponse)
}

func getIPEnv() string {
	ip := os.Getenv("KUBERNETES_IP")
	if ip == "" {
		return "localhost"
	}
	return ip
}
