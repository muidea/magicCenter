package dal

import (
	"fmt"
	"webcenter/util/modelhelper"
)

type Item struct {
	Id int
	Name string
	Url string
	Owner int
}


func AddItem(helper modelhelper.Model, name,url string, owner int) (Item, bool) {
	item := Item{}
	ret := false
	
	sql := fmt.Sprintf("insert into item (name,url,owner) values('%s','%s',%d)", name, url, owner)
	_, ret = helper.Execute(sql)
	if ret {
		ret = false
		sql = fmt.Sprintf("select id from item where name='%s' and url='%s' and owner=%d", name, url, owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&item.Id)
			item.Name = name
			item.Url = url
			item.Owner = owner
			ret = true
		}
	}
	
	return item, ret
}

func RemoveItem(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete item where id=%d", id)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

func QueryItem(helper modelhelper.Model, id int) (Item, bool) {
	item := Item{}
	ret := false
	
	sql := fmt.Sprintf("select id,name,url,owner from item where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&item.Id, &item.Name, &item.Url, &item.Owner)
		ret = true
	}
	
	return item, ret
}

func QueryItems(helper modelhelper.Model, owner int) []Item {
	itemList := []Item{}
	
	sql := fmt.Sprintf("select id,name,url,owner from item where owner=%d", owner)
	helper.Query(sql)
	for helper.Next() {
		i := Item{}
		helper.GetValue(&i.Id, &i.Name, &i.Url, &i.Owner)
		
		itemList = append(itemList, i)
	}
	
	return itemList
}
