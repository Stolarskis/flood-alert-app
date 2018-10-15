package main

import (
	"fmt"
	"testing"
)

func TestDecodingWeatherUrl(t *testing.T) {
	info, err := getSecrets("./test.yaml")
	if err != nil {
		t.Error("Failed get get secrets")
	}

	if info.URL != "URL" {
		fmt.Println("Incorrect secrets URL")
		t.Fail()
	}
}

func TestGettingInfoFromAPI(t *testing.T) {
	_, err := getInfo()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
