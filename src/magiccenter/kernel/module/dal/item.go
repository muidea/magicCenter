package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/module/model"
)

func AddItem(helper modelhelper.Model, name,url string, owner int) (model.Item, bool) {
	item := model.Item{}
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

func QueryItem(helper modelhelper.Model, id int) (model.Item, bool) {
	item := model.Item{}
	ret := false
	
	sql := fmt.Sprintf("select id,name,url,owner from item where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&item.Id, &item.Name, &item.Url, &item.Owner)
		ret = true
	}
	
	return item, ret
}

func ClearItems(helper modelhelper.Model, owner int) bool {
	sql := fmt.Sprintf("delete from item where owner=%d", owner)
	_, ok :=helper.Execute(sql)
	return ok
}

func QueryItems(helper modelhelper.Model, owner int) []model.Item {
	itemList := []model.Item{}
	
	sql := fmt.Sprintf("select id,name,url,owner from item where owner=%d", owner)
	helper.Query(sql)
	for helper.Next() {
		i := model.Item{}
		helper.GetValue(&i.Id, &i.Name, &i.Url, &i.Owner)
		
		itemList = append(itemList, i)
	}
	
	return itemList
}
