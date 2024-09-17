package sms

type smsSender interface {
	SendSms(from, to, message string) error
	Name() string
}

var providers = make(smsProviders)

type smsProviders map[string]smsSender

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}
