package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
			Action: func(c *cli.Context) {
				err := checkInfo()
				if err != nil {
					fmt.Println("Request to check info failed", err)
				}
			},
		},

		{
			Name:    "test-alerts",
			Aliases: []string{"t"},
			Usage:   "Sends out test alerts of methods flood-app uses",
			Action: func(c *cli.Context) {
				err := testAlerts()
				if err != nil {
					fmt.Println("Request to test alerts failed", err)
				}
			},
		},
		{
			Name:    "mute-alert",
			Aliases: []string{"m"},
			Usage:   "Mutes an alert based on priority. 1 2 or 3",
			Action: func(c *cli.Context) {
				err := muteAlert(c)
				if err != nil {
					fmt.Println("Request to mute alerts Failed:", err)
				}
			},
		},
		{
			Name:    "getForcast",
			Aliases: []string{"g"},
			Usage:   "Gets the current weather information",
			Action: func(c *cli.Context) {
				err := getForcast(c)
				if err != nil {
					fmt.Println("Failed to get forcast", err)
				}
			},
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

	//Get Arguments
	muteType := strings.ToLower(c.Args().Get(0))
	muteStatus, err := strconv.ParseBool(strings.ToLower(c.Args().Get(1)))
	if err != nil {
		return errors.Wrap(err, "Invalid mute alert status entered")
	}

	//Make sure that type is valid
	if muteType != "sms" && muteType != "email" {
		return errors.New("Invalid alert type entered")
	}
	//Make request to app server
	muteAlertResponse, err := client.MuteAlerts(ctx, &pb.MuteAlertRequest{MuteAlertType: c.Args().Get(0), MuteAlertStatus: muteStatus})
	if err != nil {
		return err
	}
	fmt.Println(muteAlertResponse.Output)
	return nil
}

func getForcast(c *cli.Context) error {
	client, err := pb.CreateClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	forcastResponse, err := client.GetForcast(ctx, &pb.GetForcastRequest{})
	if err != nil {
		return errors.Wrap(err, "Failed to get forcast")
	}
	fmt.Println(forcastResponse.Forcast)
	return nil
}
