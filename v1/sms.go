package main

import (
	"fmt"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	if err := sendSms("Meetup", "0509876543", "Hi Gophers !"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}

func sendSms(from, to, message string) error {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	return err
}
