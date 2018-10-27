package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Stolarskis/flood-alert-app/src/app-flood-api"
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

// Sends
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

func muteAlert(c *cli.Context) error {
	client, err := pb.CreateClient()
	if err != nil {
		return err
	}
	ctx := context.Background()
	priority, err := strconv.ParseUint(c.Args().Get(0), 10, 64)
	fmt.Println(priority)
	if err != nil {
		return errors.Wrap(err, "Failed to convert input to correct format")
	}

	if priority < 1 && priority > 3 {
		return errors.New("Incorrect input from user")
	}

	muteAlertResponse, err := client.MuteAlerts(ctx, &pb.MuteAlertRequest{MuteAlertPriority: priority})
	if err != nil {
		return errors.Wrap(err, "Failed to send message to app")
	}
	fmt.Println(muteAlertResponse)
	return nil
}
