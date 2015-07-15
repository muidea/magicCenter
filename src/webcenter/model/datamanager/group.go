package datamanager

import (

)

var ADMIN_GROUP = 1

type Group struct {
	id int
	name string
	catalog int
}

func (this *Group)Valid() bool {
	return this.id > 0 && len(this.name) > 0 && this.catalog > 0
}


