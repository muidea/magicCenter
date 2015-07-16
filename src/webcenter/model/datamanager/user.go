package datamanager

import (

)

type User struct {
	Id int
	Name string
	Password string
	Email string
	Group int
}

func (this *User)Valid() bool {
	return this.Id > 0 && len(this.Name) > 0 && len(this.Password) > 0
}

func (this *User)ValidPassword(password string) bool {
	return this.Password == password
}