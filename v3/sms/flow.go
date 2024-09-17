package sms

import "fmt"

type SupportsPhone interface {
	GetMobilePhone() string
}

type SupportsKosher interface {
	IsKosher() bool
}

func Send(recipient SupportsPhone, from, message string) error {
	to := recipient.GetMobilePhone()

	var selectedProvider = twilioId

	if sk, ok := recipient.(SupportsKosher); ok && sk.IsKosher() {
		selectedProvider = "telemessage" //not a const.. left it just to show single resp.
	} else if to[:2] == "86" {
		selectedProvider = alibabaId
	}
	if err := providers[selectedProvider].SendSms(from, to, message); err != nil {
		return fmt.Errorf("provider %s was not able to send message to %s, error: %s",
			selectedProvider, to, err.Error())
	}
	return nil
}
