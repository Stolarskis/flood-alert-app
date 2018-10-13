package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type apiInfo struct {
	URL string `yaml:"URL"`
}

type openWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeH int `json:"3h"`
	} `json:"rain"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

var a = new(apiInfo)

func init() {
	err := getSecrets()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func getInfo() (*openWeather, error) {
	info := &openWeather{}
	url := a.URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Unable to create request", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Unable to get reponse", err)
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		log.Fatal("Unable to decode response", err)
		return nil, err
	}

	return info, nil
}

func (info *openWeather) printSampleInfo() {
	fmt.Println("Humidity", info.Main.Humidity)
	fmt.Println("CurrentWeather:", info.Weather)
	fmt.Println("clouds", info.Clouds.All)
	fmt.Println("Rain", info.Rain.ThreeH)
}

func (info *openWeather) processData() {
	//Triggers
	//Its going to rain heavy - certain codes  notify 3
	switch info.Weather[0].ID {
	case 201:
		//Thunderstorm with rain - email
		alert(1, "Thunderstorm with rain")
	case 202:
		//Thunderstorm with heavy rain
		alert(2, "Thunderstorm with heavy rain")
	case 211:
		//thunderstorm - email
		alert(1, "Thunderstorm with rain")
	case 212:
		//heavy thunderstorm
		alert(2, "Thunderstorm with heavy rain")
	case 501:
		//moderate rain - email
		alert(1, "Moderate Rain")
	case 502:
		//heavy intensity rain
		alert(2, "Thunderstorm with heavy rain")
	case 503:
		//very heavy rain
		alert(2, "heavy rain")
	case 522:
		//heavy intensity shower rain
		alert(2, "heavy intensity shower rain")

	}
	//for testing:
	if info.Rain.ThreeH == 0 {
		fmt.Println("Its not raining")
	}

	//Inches of rain is 2.5-4 = Warning 2
	if info.Rain.ThreeH >= 3 {
		//Notify
		alert(2, "Rain is above 3 inches, please take car to parking garage")
	}

	//Inches of rain is above 4 = Emergency 1
	if info.Rain.ThreeH >= 4 {
		//Emergency
		alert(3, "Emergency: Rain in the past 3 hours is above 3 inches")
	}
}

func getWeatherUrl() error {
	file, err := ioutil.ReadFile("./secrets/api.yaml")
	if err != nil {
		return errors.Wrap(err, "Failed to open file")
	}
	err = yaml.Unmarshal(file, a)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal file")
	}

	return nil
}
