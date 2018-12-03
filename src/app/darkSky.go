package main

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
)

var weatherInfo = ""

//Updates the package level variable currentWeatherInfo with current weather info
func UpdateWeatherInfo() error {
	result, err := getInfoDarkSky(secrets.Location.Lat, secrets.Location.Long)
	if err != nil {
		return errors.Wrap(err, "Failed to check alerts")
	}
	weatherInfo = result.Hourly.Summary

	return nil
}

func checkAlertsDarkSky() error {
	result, err := getInfoDarkSky(secrets.Location.Lat, secrets.Location.Long)
	if err != nil {
		return errors.Wrap(err, "Failed to check alerts")
	}
	alerts := result.Alerts
	//There is a lot of things I can do with this function. Extracting different severities. For now we just go through each of the alerts.
	//Extracting any that have flood in them

	for i := 0; i < len(alerts); i++ {
		alerts[i].Title = strings.ToLower(alerts[i].Title)
		if strings.Contains(alerts[i].Title, "Flood") {
			sendSMSAlert(alerts[i].Description)
			//For now, break so I don't get a spam of text messages
			break
		}
	}
	return nil
}

func getInfoDarkSky(lat, long float64) (darksky.ForecastResponse, error) {

	//Create client
	client := darksky.New(secrets.ApiKeys.DarkSky)
	//Build Request
	request := darksky.ForecastRequest{}
	//32.7765° N, 79.9311° W
	dLong := darksky.Measurement(long)
	dLat := darksky.Measurement(lat)
	request.Longitude = dLong
	request.Latitude = dLat
	//TODO: This line breaks it for some reason
	//request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely,daily"}
	//Get the forcast
	forcast, err := client.Forecast(request)
	if err != nil {
		return darksky.ForecastResponse{}, errors.Wrap(err, "Failed to get forcast info")
	}

	return forcast, nil
}
