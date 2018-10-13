package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/sfreiberg/gotwilio"
	"gopkg.in/mailgun/mailgun-go.v1"
	"gopkg.in/yaml.v2"
)

type SValues struct {
	ApiKeys struct {
		SMSAccount      string `yaml:"SMSAccount"`
		SMSToken        string `yaml:"SMSToken"`
		EmailPrivateKey string `yaml:"EmailPrivateKey"`
		EmailPublicKey  string `yaml:"EmailPublicKey"`
		EmailDomain     string `yaml:"EmailDomain"`
	} `yaml:"apiKeys"`
}

var callMute = false
var smsMute = false
var emailMute = false

var s = new(SValues)

func init() {
	err := getSecrets()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

//severity codes
//3. Notify - Email/Slack message
//2. Warning - SMS
//1. Emergency - Call
func alert(severity int, message string) {

	switch severity {
	case 1: //For when calls actually work
		if callMute {
			sendCallAlert(message)
		}
	case 2:
		if smsMute {
			sendSMSAlert(message)
		}
	case 3: //When I can get mailgun to actually work
		if emailMute {
			sendEmailAlert(message)
		}
	}
}

func sendSMSAlert(alert string) {
	twilio := gotwilio.NewTwilioClient(s.ApiKeys.SMSAccount, s.ApiKeys.SMSToken)

	from := "+18647777941"
	//Send me a text message
	to := "+18643493157"
	message := alert
	twilio.SendSMS(from, to, message, "", "")
}

func sendEmailAlert(message string) {
	//mg := mailgun.NewMailgun(domain, privateAPIKey)
	mg := mailgun.NewMailgun(s.ApiKeys.EmailDomain, s.ApiKeys.EmailPrivateKey, s.ApiKeys.EmailPublicKey)

	from := "flood-Alert@alerts.com"
	subject := "Weather Alert"
	body := message
	recipient := "floodapptest@gmail.com"
	email := mg.NewMessage(from, subject, body, recipient)
	resp, id, err := mg.Send(email)
	if err != nil {
		fmt.Println("Unable to send message", err)
	}
	fmt.Println(resp, id)

}

func sendCallAlert(message string) {
	fmt.Println("Calling is not yet implemented. Message: ", message)
}

func getSecrets() error {
	file, err := ioutil.ReadFile("./secrets/api.yaml")
	if err != nil {
		return errors.Wrap(err, "Failed to open file")
	}
	err = yaml.Unmarshal(file, s)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal file")
	}

	return nil
}
