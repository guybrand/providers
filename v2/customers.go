package main

type customer struct {
	Name     string
	Phone    string
	IsKosher bool
}

func (c customer) GetProvider() string {
	if c.IsKosher {
		return "telemessage"
	} else if c.Phone[:2] == "86" {
		return "alibaba"
	}
	return "twilio"
}

func getCustomerFromSomeSource() customer {
	return customer{Name: "Guy", Phone: "972506666666", IsKosher: false}
}
