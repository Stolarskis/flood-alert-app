package main

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

var smsMute = false
var emailMute = false

//Sends SMS notification via twilio
func sendSMSAlert(alert string) {
	if smsMute == false {
		twilio := gotwilio.NewTwilioClient(secrets.ApiKeys.SMSAccount, secrets.ApiKeys.SMSToken)
		from := secrets.ApiKeys.ServerPhone
		//Send me a text message
		to := secrets.ApiKeys.ClientPhone
		twilio.SendSMS(from, to, alert, "", "")
	}
}

//Sends Email notification via mailgun
func sendEmailAlert(message string) {
	if emailMute == false {

		mg := mailgun.NewMailgun(secrets.ApiKeys.EmailDomain, secrets.ApiKeys.EmailPrivateKey, secrets.ApiKeys.EmailPublicKey)

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

}
