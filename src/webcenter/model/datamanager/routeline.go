package datamanager

import (

)

type Routeline struct {
	Id int
	Name string
	Description string
	Creater int
	
}

func (this *Routeline)Valid() bool {
	return this.Id > 0 && len(this.Name) > 0
}
