package main

import (
	"fmt"
	"testing"
)

func TestGetInfoDarkSky(t *testing.T) {

	_, err := getInfoDarkSky(32.7765, -79.9311)
	if err != nil {
		fmt.Println("Failed to get info via DarkSky,", err)
		t.Fail()
	}
	//fmt.Println(response.Currently.ApparentTemperature)
	//fmt.Println(response.Daily.Summary)
	//fmt.Println(response.Alerts)
}

func TestCheckAlertsDarkSky(t *testing.T) {
	checkAlertsDarkSky()
}

func TestUpdateInfoDarkSky(t *testing.T) {
	fmt.Println("Should be empty", weatherInfo)
	UpdateWeatherInfo()
	fmt.Println(weatherInfo)

}
