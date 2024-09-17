package sms

import (
	"fmt"
)

type SupportsKosher interface {
	IsKosher() bool
}

type SupportsPhone interface {
	GetMobilePhone() string
}

func (h Handler) Send(caller SupportsPhone, from, message string) error {
	to := caller.GetMobilePhone()
	var selectedProvider = "twilio"

	if sk, ok := caller.(SupportsKosher); ok && sk.IsKosher() {
		selectedProvider = "telemessage"
	} else if to[:2] == "86" {
		selectedProvider = "alibaba"
	}
	if err := providers[selectedProvider].SendSms(from, to, message); err != nil {
		return fmt.Errorf("provider %s was not able to send message to %s, error: %s",
			selectedProvider, to, err.Error())
	}
	return nil
}
