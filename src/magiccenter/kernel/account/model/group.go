package model

import (

)

const (
	ADMIN_GROUP = iota
	COMMON_GROUP
)

type Group struct {
	Id int
	Name string
	Type int
	Creater User
}

type GroupInfo struct {
	Group

	UserCount int
}

func (this *Group)AdminGroup() bool {
	return this.Type == ADMIN_GROUP
}


