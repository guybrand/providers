package sms

type smsSender interface {
	SendSms(from, to, message string) error
}

var providers = make(smsProviders)

type smsProviders map[string]smsSender

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}
