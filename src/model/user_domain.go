package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string

	Setid(string)

	Encryptpassword()
}

func NewUserDomain(
	email, 
	password, 
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

type userDomain struct {
	id	string
	email string
	password string 
	name string 
	age int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) Setid(id string) {
	ud.id = id
}

func (ud *userDomain) Encryptpassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}