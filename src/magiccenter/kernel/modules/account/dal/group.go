package dal

import (
	"fmt"
	"magiccenter/common/model"
	"magiccenter/util/dbhelper"
)

// QueryAllGroup 查询所有的分组
func QueryAllGroup(helper dbhelper.DBHelper) []model.Group {
	groupList := []model.Group{}
	sql := fmt.Sprintf("select id, name, description, catalog from `group`")
	helper.Query(sql)

	for helper.Next() {
		g := model.Group{}
		helper.GetValue(&g.ID, &g.Name, &g.Description, &g.Type)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryGroupByID 查询指定分组
func QueryGroupByID(helper dbhelper.DBHelper, id int) (model.Group, bool) {
	group := model.Group{}
	sql := fmt.Sprintf("select id, name, description,catalog from `group` where id=%d", id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&group.ID, &group.Name, &group.Description, &group.Type)
		result = true
	}

	return group, result
}

// QueryGroupByName 查询指定分组
func QueryGroupByName(helper dbhelper.DBHelper, name string) (model.Group, bool) {
	group := model.Group{}
	sql := fmt.Sprintf("select id, name, description, catalog from `group` where name='%s'", name)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&group.ID, &group.Name, &group.Description, &group.Type)
		result = true
	}

	return group, result
}

// DeleteGroup 删除分组
func DeleteGroup(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from `group` where id =%d", id)
	_, result := helper.Execute(sql)

	return result
}

// SaveGroup 保存分组
func SaveGroup(helper dbhelper.DBHelper, group model.Group) bool {
	sql := fmt.Sprintf("select id from `group` where id=%d", group.ID)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `group` (name, description, catalog) values ('%s','%s',%d)", group.Name, group.Description, group.Type)
	} else {
		// modify
		sql = fmt.Sprintf("update `group` set name ='%s', description='%s' where id=%d", group.Name, group.Description, group.ID)
	}

	_, ret := helper.Execute(sql)
	return ret
}
