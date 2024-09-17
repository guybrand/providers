package main

import "fmt"

type smsSender interface {
	SendSms(from, to, message string) error
}

var providers = make(map[string]smsSender)

func main() {
	cs := getCustomerFromSomeSource()
	if err := providers[cs.GetProvider()].SendSms("Go-meetup", cs.Phone, "How are you ?"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
