package module

import (
	"webcenter/util/modelhelper"
)

type Block interface {
	ID() int
	Name() string
	Owner() string
	Items() []Item
}

type block struct {
	id int
	name string
	owner string
	items []Item
}

func (this *block)ID() int {
	return this.id
}

func (this *block)Name() string {
	return this.name
}

func (this *block)Owner() string {
	return this.owner
}

func (this *block)Items() []Item {
	return this.items
}

func InsertModuleBlock(name, owner string) (Block,bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	return insertModuleBlock(helper,name,owner)	
}

func DeleteModuleBlock(name string) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	deleteModuleBlock(helper, name)	
}

func QueryModuleBlocks(owner string) []Block {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()

	return queryModuleBlocks(helper, owner)	
}




