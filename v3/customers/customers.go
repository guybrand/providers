package customers

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
