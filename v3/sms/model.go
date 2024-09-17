package sms

type smsSender interface {
	SendSms(from, to, message string) error
}

type smsProviders map[string]smsSender

var providers = make(smsProviders)
