package main

import (
	"go.uber.org/fx"
	"providers/customers"
	"providers/employees"
	"providers/sms"
	//alibaba "providers/sms/provider/alibaba"
	twilio "providers/sms/provider/twilio"
)

func main() {
	fx.New(
		fx.Provide(
			customers.NewHandler,
			employees.NewHandler,

			//sms.AsProvider(alibaba.NewProvider),
			sms.AsProvider(twilio.NewProvider),

			NewMyService,
		),
		fx.Invoke(doSomeServiceStuff),
	)
}

func doSomeServiceStuff(m MyService) {
	m.customersH.DoSomething()
	m.employeesH.DoSomething()
}
