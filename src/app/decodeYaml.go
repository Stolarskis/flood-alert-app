package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type SValues struct {
	ApiKeys struct {
		SMSAccount      string `yaml:"SMSAccount"`
		SMSToken        string `yaml:"SMSToken"`
		EmailPrivateKey string `yaml:"EmailPrivateKey"`
		EmailPublicKey  string `yaml:"EmailPublicKey"`
		EmailDomain     string `yaml:"EmailDomain"`
		ClientPhone     string `yaml:"ClientPhone"`
		ServerPhone     string `yaml:"ServerPhone"`
		DarkSky         string `yaml:"DarkSky"`
	} `yaml:"apiKeys"`
	URL string `yaml:"URL"`
}

var secrets = new(SValues)

//populate the secret instance variable
func init() {
	secretInfo, err := getSecrets("./secrets/api.yaml")
	if err != nil {
		fmt.Println(err, ", exiting")
		os.Exit(1)
	}
	secrets = secretInfo
}

//Get Secret values from the yaml file
func getSecrets(yamlPath string) (*SValues, error) {
	secretInfo := new(SValues)
	file, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read secrets file")
	}
	err = yaml.Unmarshal(file, secretInfo)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to unmarshal secrets yaml")
	}

	return secretInfo, nil
}
