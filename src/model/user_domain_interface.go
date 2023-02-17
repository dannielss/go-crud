package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string

	SetID(string)

	EncryptPassword()
}

func NewUserDomain(
	email string, 
	password string, 
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		id: "",
		email: email,
		password: password,
		name: name,
		age: age,
	}
}