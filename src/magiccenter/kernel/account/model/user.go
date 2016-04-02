package model

import (

)

const (
	CREATE = iota
	DEACTIVE
	ACTIVE
)

type User struct {
	Id int
	Name string	
}

type UserDetail struct {
	User
	
	Account string
	Email string
	Status int
	Groups []Group
}
