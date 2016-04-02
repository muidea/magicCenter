package model

import (
	"magiccenter/kernel/account/model"
)

type LinkDetail struct {
	Id int
	Name string
	Url string			
	Logo string
	Creater model.User
	
	Catalog []Catalog	
}

func (this *LinkDetail) RId() int {
	return this.Id
}

func (this *LinkDetail) RName() string {
	return this.Name
}

func (this *LinkDetail) RType() int {
	return LINK
}

func (this *LinkDetail) RRelative() []Resource {
	relatives := []Resource{}
	
	for _, c := range this.Catalog {
		relatives = append(relatives, &c)
	}
	
	return relatives
}
