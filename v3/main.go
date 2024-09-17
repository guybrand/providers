package main

import (
	"fmt"
	"providers/customers"
	"providers/employees"
	"providers/sms"
)

func main() { //when an online order is completed
	cs := customers.GetOne()
	fmt.Println("sending to customer... ", sms.Send(cs, "My online shop", "We got your order !"))

	em := employees.GetOne()
	fmt.Println("sending to employee... ", sms.Send(em, "My workplace name", "Please deliver this order"))
}
