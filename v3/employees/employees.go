package employees

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
