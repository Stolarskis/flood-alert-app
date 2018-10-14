package main

import (
	"fmt"
	"testing"
)

func TestWeatherUrl(t *testing.T) {
	err := getWeatherUrl("./test.yaml")
	if err != nil {
		t.Error("Failed get get secrets")
	}
	fmt.Println("hello")
	fmt.Println(a)
}
