package main

import (
	"fmt"
	"testing"
)

func TestGetSecretsAlert(t *testing.T) {
	err := getSecrets()
	if err != nil {
		t.Error("Failed get get secrets")
	}

	fmt.Println(s)
}
