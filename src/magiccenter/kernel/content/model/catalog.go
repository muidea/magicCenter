package model

import (
	"magiccenter/kernel/account/model"
)

type Catalog struct {
	Id int
	Name string
}

type CatalogDetail struct {
	Catalog

	Creater model.User
	Parent []Catalog
}


func (this *Catalog) RId() int {
	return this.Id
}

func (this *Catalog) RName() string {
	return this.Name
}

func (this *Catalog) RType() int {
	return CATALOG
}

func (this *Catalog) RRelative() []Resource {
	relatives := []Resource{}
	return relatives
}

func (this *CatalogDetail) RId() int {
	return this.Id
}

func (this *CatalogDetail) RName() string {
	return this.Name
}

func (this *CatalogDetail) RType() int {
	return CATALOG
}

func (this *CatalogDetail) RRelative() []Resource {
	relatives := []Resource{}
	
	for i, _ := range this.Parent {
		p := &this.Parent[i]
		relatives = append(relatives, p)
	}
	
	return relatives
}




