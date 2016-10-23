package dal

import (
	"fmt"
	"magiccenter/common/model"
	"magiccenter/util/dbhelper"
)

// AddItem 添加一条Item记录
func AddItem(helper dbhelper.DBHelper, item model.Item) (model.Item, bool) {
	ret := false

	sql := fmt.Sprintf("insert into item (rid,rtype,owner) values(%d,'%s',%d)", item.Rid, item.Rtype, item.Owner)
	_, ret = helper.Execute(sql)
	if ret {
		ret = false
		sql = fmt.Sprintf("select id from item where rid=%d and rtype='%s' and owner=%d", item.Rid, item.Rtype, item.Owner)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&item.ID)
			ret = true
		}
	}

	return item, ret
}

// RemoveItem 删除一条Item记录
func RemoveItem(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete item where id=%d", id)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

// QueryItem 查询Item记录
func QueryItem(helper dbhelper.DBHelper, id int) (model.Item, bool) {
	item := model.Item{}
	ret := false

	sql := fmt.Sprintf("select id,rid,rtype,owner from item where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&item.ID, &item.Rid, &item.Rtype, &item.Owner)
		ret = true
	}

	return item, ret
}

// ClearItems 清除指定owner的Items记录
func ClearItems(helper dbhelper.DBHelper, owner int) bool {
	sql := fmt.Sprintf("delete from item where owner=%d", owner)
	_, ok := helper.Execute(sql)
	return ok
}

// QueryItems 查询指定owner的Items记录
func QueryItems(helper dbhelper.DBHelper, owner int) []model.Item {
	itemList := []model.Item{}

	sql := fmt.Sprintf("select id,rid,rtype,owner from item where owner=%d", owner)
	helper.Query(sql)
	for helper.Next() {
		i := model.Item{}
		helper.GetValue(&i.ID, &i.Rid, &i.Rtype, &i.Owner)

		itemList = append(itemList, i)
	}

	return itemList
}
