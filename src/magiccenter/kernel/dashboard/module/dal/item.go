package dal

import (
	"fmt"
	"magiccenter/kernel/dashboard/module/model"
	contentmodel "magiccenter/kernel/modules/content/model"
	"magiccenter/util/modelhelper"
)

func AddItem(helper modelhelper.Model, rid, rtype, owner int) (model.Item, bool) {
	item := model.Item{}
	ret := false

	sql := fmt.Sprintf("insert into item (rid,rtype,owner) values(%d,%d,%d)", rid, rtype, owner)
	_, ret = helper.Execute(sql)
	if ret {
		ret = false
		sql = fmt.Sprintf("select id from item where rid=%d and rtype=%d and owner=%d", rid, rtype, owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&item.Id)
			item.Rid = rid
			item.Rtype = rtype
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

	sql := fmt.Sprintf("select id,rid,rtype,owner from item where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&item.Id, &item.Rid, &item.Rtype, &item.Owner)
		ret = true
	}

	return item, ret
}

func ClearItems(helper modelhelper.Model, owner int) bool {
	sql := fmt.Sprintf("delete from item where owner=%d", owner)
	_, ok := helper.Execute(sql)
	return ok
}

func QueryItems(helper modelhelper.Model, rtype, owner int) []model.Item {
	itemList := []model.Item{}

	sql := fmt.Sprintf("select id,rid,rtype,owner from item where rtype=%d and owner=%d", rtype, owner)
	helper.Query(sql)
	for helper.Next() {
		i := model.Item{}
		helper.GetValue(&i.Id, &i.Rid, &i.Rtype, &i.Owner)

		itemList = append(itemList, i)
	}

	return itemList
}

func QueryItemViews(helper modelhelper.Model, owner int, uri string) []model.ItemView {
	itemList := []model.ItemView{}

	sql := fmt.Sprintf("select i.rid, r.`name`, i.rtype from item i, resource r where i.rid = r.id and i.rtype = r.type and i.`owner` = %d", owner)
	helper.Query(sql)
	for helper.Next() {
		item := model.ItemView{}
		otype := 0
		helper.GetValue(&item.Id, &item.Name, &otype)
		switch otype {
		case contentmodel.ARTICLE:
			item.Url = fmt.Sprintf("%sview/?id=%d", uri, item.Id)
		case contentmodel.CATALOG:
			item.Url = fmt.Sprintf("%scatalog/?id=%d", uri, item.Id)
		case contentmodel.LINK:
			item.Url = fmt.Sprintf("%slink/?id=%d", uri, item.Id)
		default:
			item.Url = fmt.Sprintf("%s404/?id=%d", uri, item.Id)
		}
		itemList = append(itemList, item)
	}

	return itemList
}

func QuerySubItemViews(helper modelhelper.Model, owner int, uri string) []model.ItemView {
	itemList := []model.ItemView{}
	sql := fmt.Sprintf("select r.id, r.`name`, r.type from resource r, resource_relative rr where r.id = rr.src and r.type = rr.srcType and rr.dst = %d and rr.dstType = %d", owner, contentmodel.CATALOG)
	helper.Query(sql)
	for helper.Next() {
		item := model.ItemView{}
		otype := 0
		helper.GetValue(&item.Id, &item.Name, &otype)
		switch otype {
		case contentmodel.ARTICLE:
			item.Url = fmt.Sprintf("%sview/?id=%d", uri, item.Id)
		case contentmodel.CATALOG:
			item.Url = fmt.Sprintf("%scatalog/?id=%d", uri, item.Id)
		case contentmodel.LINK:
			item.Url = fmt.Sprintf("%slink/?id=%d", uri, item.Id)
		default:
			item.Url = fmt.Sprintf("%s404/?id=%d", uri, item.Id)
		}
		itemList = append(itemList, item)
	}

	return itemList
}
