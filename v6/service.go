package main

import (
	"providers/customers"
	"providers/employees"
)

type MyService struct {
	customersH *customers.Handler
	employeesH *employees.Handler
}

func NewMyService(c *customers.Handler, e *employees.Handler) MyService {
	return MyService{customersH: c, employeesH: e}
}
