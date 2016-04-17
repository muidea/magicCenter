package model

import (
	"magiccenter/kernel/account/model"
)

type Link struct {
	Id int
	Name string
	Url string			
	Logo string
	Creater model.User
	
	Catalog []Catalog	
}

func (this *Link) RId() int {
	return this.Id
}

func (this *Link) RName() string {
	return this.Name
}

func (this *Link) RType() int {
	return LINK
}

func (this *Link) RRelative() []Resource {
	relatives := []Resource{}
	
	for i, _ := range this.Catalog {
		c := &this.Catalog[i]
		relatives = append(relatives, c)
	}
	
	return relatives	
}
