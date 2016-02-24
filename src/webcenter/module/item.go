package module

import (
	"webcenter/util/modelhelper"
)

type Item interface {
	ID() int
	Name() string	
	Url() string
	Owner() int
}

type item struct {
	id int
	name string
	url string
	owner int
}

func (this *item)ID() int {
	return this.id
}

func (this *item)Name() string {
	return this.name
}

func (this *item)Url() string {
	return this.url
}

func (this *item)Owner() int {
	return this.owner
}

func AddItem(name,url string, owner int) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()

	return addBlockItem(helper, name,url,owner)
}

func RemoveItem(id int) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()

	removeBlockItem(helper, id)	
}

func QueryBlockItems(owner int) []Item {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()

	return queryBlockItems(helper,owner)	
}

