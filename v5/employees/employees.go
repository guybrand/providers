package employees

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

type contactDetails struct {
	CellPhone string
	HomePhone string //So 90's
	Extension string //So 80's ...
}

type employee struct {
	Id   int
	Name string
	contactDetails
}

func (e employee) GetMobilePhone() string {
	return e.contactDetails.CellPhone
}

func GetOne() employee {
	return employee{Id: 1, Name: "Number 1", contactDetails: contactDetails{CellPhone: "860545555555"}}
}
func (h handler) DoSomething() {
	if err := h.smsSender.Send(GetOne(), "My prev workplace", "Are you still here ?"); err != nil {
		fmt.Println(err.Error())
	}
}
