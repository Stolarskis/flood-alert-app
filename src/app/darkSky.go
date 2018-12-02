package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/shawntoffel/darksky"
)

func getInfoDarkSky() (*darksky.ForecastResponse, error) {

	fmt.Println(secrets.ApiKeys.DarkSky)
	fmt.Println("Hello World")
	//Create client
	client := darksky.New(secrets.ApiKeys.DarkSky)
	//Build Request
	request := darksky.ForecastRequest{}
	//32.7765° N, 79.9311° W
	request.Longitude = 32.7765
	request.Latitude = 79.9311
	//Get the forcast
	forcast, err := client.Forecast(request)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get forcast info")
	}
	return forcast, nil

}
