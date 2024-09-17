package sms

import (
	"go.uber.org/fx"
)

type SmsSender interface {
	SendSms(from, to, message string) error
}

func AsProvider(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(SmsSender)),
	)
}
