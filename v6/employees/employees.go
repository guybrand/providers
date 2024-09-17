package employees

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
func (h Handler) DoSomething() {
	if err := h.smsSender.SendSms(GetOne().GetMobilePhone(), "NewJob", "Now you really have to work hard!"); err != nil {
		fmt.Println(err.Error())
	}
}
