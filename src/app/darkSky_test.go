package main

import (
	"fmt"
	"testing"

	"github.com/shawntoffel/darksky"
)

func TestConnectionToDarkSky(t *testing.T) {
	//An error will be thrown if its unable to connect to darksky, tested this by turning off my wifi.
	//
	_, err := getForcast(secrets.Location.Lat, secrets.Location.Long)
	if err != nil {
		fmt.Println("Failed to get info via DarkSky,", err)
		t.Fail()
	}
	//fmt.Println(response.Currently.ApparentTemperature)
	//fmt.Println(response.Daily.Summary)
	//fmt.Println(response.Alerts)
}

func TestCheckAlerts(t *testing.T) {
	//Flood alerts are rare, better to test a couple things here.

	//1. Check if our alert is sent the first time.
	mockForcast := darksky.ForecastResponse{}
	alert := &darksky.Alert{
		Title:       "THERES A FLOOD RUN FOR THE HILLS",
		Severity:    "Warning",
		Description: "Test Description of flood alert",
	}
	mockForcast.Alerts = append(mockForcast.Alerts, alert)
	forcast = mockForcast
	CheckFloodAlerts()
	//2. Give the alert again and the message shouldn't be sent a second time.
	CheckFloodAlerts()
	//Note: I should be able to mock the smsAlert call.
}

func TestUpdateInfoDarkSkyFromNull(t *testing.T) {
	//important to note that this is testing if null package variable is updated or not.
	//Lat will always be 0 to start.
	testLat := forcast.Latitude
	updateForcast()
	if testLat == forcast.Latitude {
		fmt.Println("Failed to update")
	}

}
