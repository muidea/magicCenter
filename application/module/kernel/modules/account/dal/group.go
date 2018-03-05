package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/util"
)

// QueryAllGroup 查询所有的分组
func QueryAllGroup(helper dbhelper.DBHelper) []model.GroupDetail {
	groupList := []model.GroupDetail{}
	sql := fmt.Sprintf("select id, name, description, catalog from account_group")
	helper.Query(sql)

	for helper.Next() {
		g := model.GroupDetail{}
		helper.GetValue(&g.ID, &g.Name, &g.Description, &g.Catalog)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryGroups 查询分组信息
func QueryGroups(helper dbhelper.DBHelper, ids []int) []model.GroupDetail {
	groupList := []model.GroupDetail{}
	sql := fmt.Sprintf("select id, name, description, catalog from account_group where id in(%s)", util.IntArray2Str(ids))
	helper.Query(sql)

	for helper.Next() {
		g := model.GroupDetail{}
		helper.GetValue(&g.ID, &g.Name, &g.Description, &g.Catalog)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryGroupByID 查询指定分组
func QueryGroupByID(helper dbhelper.DBHelper, id int) (model.GroupDetail, bool) {
	group := model.GroupDetail{}
	sql := fmt.Sprintf("select id, name, description,catalog from account_group where id=%d", id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&group.ID, &group.Name, &group.Description, &group.Catalog)
		result = true
	}

	return group, result
}

// QueryGroupByName 查询指定分组
func QueryGroupByName(helper dbhelper.DBHelper, name string) (model.GroupDetail, bool) {
	group := model.GroupDetail{}
	sql := fmt.Sprintf("select id, name, description, catalog from account_group where name='%s'", name)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&group.ID, &group.Name, &group.Description, &group.Catalog)
		result = true
	}

	return group, result
}

// CreateGroup 新建分组
func CreateGroup(helper dbhelper.DBHelper, name, description string, catalog int) (model.GroupDetail, bool) {
	group := model.NewGroup(name, description, catalog)
	sql := fmt.Sprintf("select id from account_group where name='%s' and catalog=%d", name, 0)
	helper.Query(sql)
	if helper.Next() {
		return group, false
	}

	sql = fmt.Sprintf("insert into account_group (name, description, catalog) values ('%s','%s',%d)", name, description, catalog)
	_, result := helper.Execute(sql)
	if !result {
		return group, result
	}

	sql = fmt.Sprintf("select id from account_group where name='%s' and description='%s' and catalog=%d", name, description, catalog)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&group.ID)
		group.Name = name
		group.Description = description
		group.Catalog = 0
		result = true
	} else {
		result = false
	}

	return group, result
}

// DeleteGroup 删除分组
func DeleteGroup(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from account_group where id =%d", id)
	_, result := helper.Execute(sql)

	return result
}

// SaveGroup 保存分组
func SaveGroup(helper dbhelper.DBHelper, group model.GroupDetail) (model.GroupDetail, bool) {
	sql := fmt.Sprintf("select id from account_group where id=%d", group.ID)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into account_group (name, description, catalog) values ('%s','%s', %d)", group.Name, group.Description, group.Catalog)
	} else {
		// modify
		sql = fmt.Sprintf("update account_group set name ='%s', description='%s', catalog=%d where id=%d", group.Name, group.Description, group.ID, group.Catalog)
	}

	_, ret := helper.Execute(sql)
	return group, ret
}
