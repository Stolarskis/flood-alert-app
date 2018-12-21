package main

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
)

var pAlerts map[string]string

var forcast darksky.ForecastResponse

func init() {
	pAlerts = make(map[string]string)
}

//CheckFloodAlertsDarkSky Checks the current flood alerts
func CheckFloodAlerts() error {
	alerts := forcast.Alerts

	//There is a lot of things I can do with this function. Extracting different severities. For now we just go through each of the alerts.
	//Extracting any that have flood in them
	for i := 0; i < len(alerts); i++ {
		alerts[i].Title = strings.ToLower(alerts[i].Title)
		//Check if flood exists in the title
		if strings.Contains(alerts[i].Title, "flood") {
			//Check if the alert is already in the hash table
			if _, ok := pAlerts[alerts[i].Title]; !ok {
				sendSMSAlert(alerts[i].Description)
				//Add it to the hash table
				pAlerts[alerts[i].Title] = alerts[i].Severity
			}
			//Otherwise don't notify
		}
	}
	return nil
}

//getCurrentInfo returns the current info from the last update
func getCurrentForcast() (darksky.ForecastResponse, error) {
	forcast, err := getForcast(secrets.Location.Lat, secrets.Location.Long)
	if err != nil {
		return darksky.ForecastResponse{}, err
	}
	return forcast, nil

}

func updateForcast() error {
	result, err := getForcast(secrets.Location.Lat, secrets.Location.Long)
	if err != nil {
		return errors.Wrap(err, "Failed to update forcast")
	}

	forcast = result
	return nil
}

func getForcast(lat, long float64) (darksky.ForecastResponse, error) {

	//Create client
	client := darksky.New(secrets.ApiKeys.DarkSky)
	//Build Request
	request := darksky.ForecastRequest{}

	dLong := darksky.Measurement(long)
	dLat := darksky.Measurement(lat)
	request.Longitude = dLong
	request.Latitude = dLat
	forcast, err := client.Forecast(request)
	if err != nil {
		return darksky.ForecastResponse{}, errors.Wrap(err, "Failed to get forcast info")
	}

	return forcast, nil
}
