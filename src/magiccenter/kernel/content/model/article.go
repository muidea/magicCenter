package model

import (
	"magiccenter/kernel/account/model"
)

type ArticleSummary struct {
	Id int
	Title string
	CreateDate string
	Catalog []Catalog
	Author model.User
}

type Article struct {
	ArticleSummary
	
	Content string
}

func (this *Article) RId() int {
	return this.Id
}

func (this *Article) RName() string {
	return this.Title
}

func (this *Article) RType() int {
	return ARTICLE
}

func (this *Article) RRelative() []Resource {
	relatives := []Resource{}
	
	for i, _ := range this.Catalog {
		c := &this.Catalog[i]
		relatives = append(relatives, c)
	}
	
	return relatives
}



