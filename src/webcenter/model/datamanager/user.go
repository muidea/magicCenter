package datamanager

import (

)

type User struct {
	id int
	name string
	password string
	email string
	group int
}

func (this *User)Valid() bool {
	return this.id > 0 && len(this.name) > 0 && len(this.password) > 0
}

func (this *User)ValidPassword(password string) bool {
	return this.password == password
}