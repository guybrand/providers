package customers

import (
	"fmt"
	"providers/sms"
)

type Handler struct {
	smsSender sms.SmsSender
}

func NewHandler(s sms.SmsSender) *Handler {
	return &Handler{s}
}

type customer struct {
	Name   string
	Phone  string
	kosher bool
}

func (c customer) GetMobilePhone() string {
	return c.Phone
}

func GetOne() customer {
	return customer{Name: "Guy", Phone: "972506666666", kosher: false}
}

func (h Handler) DoSomething() {
	if err := h.smsSender.SendSms(GetOne().GetMobilePhone(), "Shop who?", "Vamos a casa, adiós!"); err != nil {
		fmt.Println(err.Error())
	}
}
