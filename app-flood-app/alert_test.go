package main

import (
	"fmt"
	"testing"
)

func TestGetSecretsAlert(t *testing.T) {
	info, err := getSecrets("./test.yaml")
	if err != nil {
		t.Error("Failed get get secrets")
	}

	if info.ApiKeys.SMSToken != "SMSToken" {
		fmt.Println("Incorrect SMSToken Values")
		t.Fail()
	}

}

func TestGetSecretsAlertPhones(t *testing.T) {
	info, err := getSecrets("./test.yaml")
	if err != nil {
		t.Error("Failed get get secrets")
	}

	//fmt.Println(info.ApiKeys.ClientPhone)
	//fmt.Println(info.ApiKeys.ServerPhone)
	if info.ApiKeys.ClientPhone != "ClientPhone" {
		fmt.Println("Incorrect ClientPhone Value")
		t.Fail()
	}
	if info.ApiKeys.ServerPhone != "ServerPhone" {
		fmt.Println("Incorrect ServerPhone Value")
		t.Fail()
	}
}
