package main

import (
	"providers/customers"
	"providers/employees"
	"providers/sms"
)

func main() {

	smsHandler := sms.NewHandler()

	customerHandler := customers.NewHandler(smsHandler)
	customerHandler.DoSomething()

	employeesHandler := employees.NewHandler(smsHandler)
	employeesHandler.DoSomething()

}
