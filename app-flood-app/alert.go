package main

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
	"gopkg.in/mailgun/mailgun-go.v1"
)

var callMute = false
var smsMute = false
var emailMute = false

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
	twilio := gotwilio.NewTwilioClient(secrets.ApiKeys.SMSAccount, secrets.ApiKeys.SMSToken)

	from := secrets.ApiKeys.ServerPhone
	//Send me a text message
	to := secrets.ApiKeys.ClientPhone
	message := alert
	twilio.SendSMS(from, to, message, "", "")
}

func sendEmailAlert(message string) {
	//mg := mailgun.NewMailgun(domain, privateAPIKey)
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

func sendCallAlert(message string) {
	fmt.Println("Calling is not yet implemented. Message: ", message)
}
