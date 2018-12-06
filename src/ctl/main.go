package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/Stolarskis/flood-alert-app/src/api"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "check-alerts",
			Aliases: []string{"c"},
			Usage:   "Main service call: Calls the main process of flood-app. Which gets info from openweather and then checks for conditions to alert user",
			Action: func(c *cli.Context) error {
				err := checkInfo()
				if err != nil {
					fmt.Println("Request to check info failed", err)
				}
				return nil
			},
		},

		{
			Name:    "test-alerts",
			Aliases: []string{"t"},
			Usage:   "Sends out test alerts of methods flood-app uses",
			Action: func(c *cli.Context) error {
				err := testAlerts()
				if err != nil {
					fmt.Println("Request to test alerts failed", err)
				}
				return nil
			},
		},
		{
			Name:    "mute-alert",
			Aliases: []string{"m"},
			Usage:   "Mutes an alert based on priority. 1 2 or 3",
			Action: func(c *cli.Context) error {
				err := muteAlert(c)
				if err != nil {
					fmt.Println("Request to mute alerts Failed:", err)
				}
				return nil
			},
		},
		{
			Name: "getForcast",
			Aliases: []string["g"],
			Usage: "Gets the current weather information",
			Action: func(c *cli.Context) error {
				err := muteAlert(c)
			}

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func checkInfo() error {
	fmt.Println("Checking Info")
	ctx := context.Background()
	client, err := pb.CreateClient()
	if err != nil {
		return err
	}
	checkAlertsRequest := &pb.CheckAlertsRequest{
		AlertMode: 0,
	}
	checkAlertsResponse, err := client.CheckAlerts(ctx, checkAlertsRequest)
	if err != nil {
		return err
	}
	fmt.Println(checkAlertsResponse.Output)
	return nil
}

// Tests methods of alerts 
func testAlerts() error {
	client, err := pb.CreateClient()
	if err != nil {
		return err
	}
	testAlertsRequest := &pb.TestAlertsRequest{TestMessage: "Test Message"}
	ctx := context.Background()
	testAlertsResponse, err := client.TestAlerts(ctx, testAlertsRequest)
	if err != nil {
		return err
	}
	fmt.Println(testAlertsResponse)
	return nil
}

//Mutes alerts, enter "Email" or "SMS"
func muteAlert(c *cli.Context) error {
	client, err := pb.CreateClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	muteAlertResponse, err := client.MuteAlerts(ctx, &pb.MuteAlertRequest{MuteAlertType: c.Args().Get(0)})
	if err != nil {
		return errors.Wrap(err, "Failed to send message to app")
	}
	fmt.Println(muteAlertResponse)
	return nil
}

func getForcast(c *cli.Context) error{
	client, err := pb.CreateClient()
	if err != nil{
		return err
	}
	ctx := context.Background()

	forcastResponse,err := client.
}