package sms

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioProvider struct {
	ApiKey     string
	Secret     string
	AccountSid string
}

func init() {
	providers["twilio"] = twilioProvider{AccountSid: "123", Secret: "abc", ApiKey: "key123"}
}

func (t twilioProvider) SendSms(from, to, message string) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{t.ApiKey, t.Secret, t.AccountSid, nil})

	params := &openapi.CreateMessageParams{}

	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)

	return err
}
