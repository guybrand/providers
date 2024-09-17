package sms

type SupportsKosher interface {
	IsKosher() bool
}

type SupportsPhone interface {
	GetMobilePhone() string
}
