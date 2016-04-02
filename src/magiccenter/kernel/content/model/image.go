package model

import (
	"magiccenter/kernel/account/model"
)

type ImageDetail struct {
	Id int
	Name string
	Url string		
	Desc string
	Creater model.User	
	Catalog []Catalog
}

func (this *ImageDetail) RId() int {
	return this.Id
}

func (this *ImageDetail) RName() string {
	return this.Name
}

func (this *ImageDetail) RType() int {
	return IMAGE
}

func (this *ImageDetail) RRelative() []Resource {
	relatives := []Resource{}
	
	for _, c := range this.Catalog {
		relatives = append(relatives, &c)
	}
	
	return relatives
}
