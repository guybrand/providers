package customers

import (
	"fmt"
	"providers/sms"
)

type handler struct {
	smsSender sms.Handler
}

func NewHandler(s sms.Handler) handler {
	return handler{s}
}

type customer struct {
	Name   string
	Phone  string
	kosher bool
}

func (c customer) IsKosher() bool {
	return c.kosher
}

func (c customer) GetMobilePhone() string {
	return c.Phone
}

func GetOne() customer {
	return customer{Name: "Guy", Phone: "972506666666", kosher: false}
}

func (h handler) DoSomething() {
	if err := h.smsSender.Send(GetOne(), "My hated shop", "and you keep buying here FCOL..."); err != nil {
		fmt.Println(err.Error())
	}
}
