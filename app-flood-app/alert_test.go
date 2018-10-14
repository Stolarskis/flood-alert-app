package main

import (
	"fmt"
	"testing"
)

func TestGetSecretsAlert(t *testing.T) {
	err := getSecrets("./test.yaml")
	if err != nil {
		t.Error("Failed get get secrets")
	}

	fmt.Println(s)
}

func TestGetSecretsAlertPhones(t *testing.T) {
	err := getSecrets("./test.yaml")
	if err != nil {
		t.Error("Failed get get secrets")
	}

	fmt.Println(s.ApiKeys.ClientPhone)
	fmt.Println(s.ApiKeys.ServerPhone)
}
