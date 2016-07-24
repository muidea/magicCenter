package dal

import (
	"fmt"
	"magiccenter/kernel/modules/account/model"
	"magiccenter/util/modelhelper"
)

// QueryAllGroup 查询所有的分组
func QueryAllGroup(helper modelhelper.Model) []model.Group {
	groupList := []model.Group{}
	sql := fmt.Sprintf("select id, name, catalog from `group`")
	helper.Query(sql)

	for helper.Next() {
		g := model.Group{}
		helper.GetValue(&g.Id, &g.Name, &g.Type)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryGroupByID 查询指定分组
func QueryGroupByID(helper modelhelper.Model, id int) (model.Group, bool) {
	group := model.Group{}
	sql := fmt.Sprintf("select id,name,catalog from `group` where id=%d", id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&group.Id, &group.Name, &group.Type)
		result = true
	}

	return group, result
}

// QueryGroupByName 查询指定分组
func QueryGroupByName(helper modelhelper.Model, name string) (model.Group, bool) {
	group := model.Group{}
	sql := fmt.Sprintf("select id,name,catalog from `group` where name='%s'", name)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&group.Id, &group.Name, &group.Type)
		result = true
	}

	return group, result
}

// DeleteGroup 删除分组
func DeleteGroup(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from `group` where id =%d", id)
	_, result := helper.Execute(sql)

	return result
}

// SaveGroup 保存分组
func SaveGroup(helper modelhelper.Model, group model.Group) bool {
	sql := fmt.Sprintf("select id from `group` where id=%d", group.Id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		group.Type = 1
		sql = fmt.Sprintf("insert into `group` (name, catalog) values ('%s',%d)", group.Name, group.Type)
	} else {
		// modify
		sql = fmt.Sprintf("update `group` set name ='%s' where id=%d", group.Name, group.Id)
	}

	_, ret := helper.Execute(sql)
	return ret
}
