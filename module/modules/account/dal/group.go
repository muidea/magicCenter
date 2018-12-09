package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/common/dbhelper"
	common_const "muidea.com/magicCommon/common"
	common_util "muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

func loadGroupID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from account_group`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// QueryGroupCount 查询分组数量
func QueryGroupCount(helper dbhelper.DBHelper) int {
	sql := fmt.Sprintf("select count(id) from account_group")
	helper.Query(sql)
	defer helper.Finish()

	countVal := 0
	if helper.Next() {
		helper.GetValue(&countVal)
	}

	return countVal
}

// QueryAllGroupDetail 查询所有的分组
func QueryAllGroupDetail(helper dbhelper.DBHelper, filter *common_util.PageFilter) []model.GroupDetail {
	groupList := []model.GroupDetail{}
	sql := fmt.Sprintf("select id, name, description, catalog from account_group")
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		g := model.GroupDetail{}
		helper.GetValue(&g.ID, &g.Name, &g.Description, &g.Catalog)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryAllGroup 查询所有的分组
func QueryAllGroup(helper dbhelper.DBHelper) []model.Group {
	groupList := []model.Group{}
	sql := fmt.Sprintf("select id, name from account_group")
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		g := model.Group{}
		helper.GetValue(&g.ID, &g.Name)

		groupList = append(groupList, g)
	}

	return groupList
}

// QuerySubGroups 查询指定分组的子分组
func QuerySubGroups(helper dbhelper.DBHelper, id int) []model.Group {
	groupList := []model.Group{}

	sql := fmt.Sprintf("select id, name from account_group where catalog =%d", id)
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		g := model.Group{}
		helper.GetValue(&g.ID, &g.Name)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryGroups 查询分组信息
func QueryGroups(helper dbhelper.DBHelper, ids []int) []model.Group {
	groupList := []model.Group{}
	if len(ids) == 0 {
		return groupList
	}

	sql := fmt.Sprintf("select id, name from account_group where id in(%s)", common_util.IntArray2Str(ids))
	helper.Query(sql)

	for helper.Next() {
		g := model.Group{}
		helper.GetValue(&g.ID, &g.Name)

		groupList = append(groupList, g)
	}

	return groupList
}

// QueryGroupByID 查询指定分组
func QueryGroupByID(helper dbhelper.DBHelper, id int) (model.GroupDetail, bool) {
	group := model.GroupDetail{}
	sql := fmt.Sprintf("select id, name, description,catalog from account_group where id=%d", id)
	helper.Query(sql)
	defer helper.Finish()

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
	defer helper.Finish()

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
	sql := fmt.Sprintf("select id from account_group where name='%s' and catalog=%d", name, catalog)
	helper.Query(sql)

	if helper.Next() {
		helper.Finish()
		return group, false
	}
	helper.Finish()

	id := allocGroupID()
	sql = fmt.Sprintf("insert into account_group (id, name, description, catalog) values (%d, '%s','%s',%d)", id, name, description, catalog)
	_, result := helper.Execute(sql)
	if !result {
		return group, result
	}

	group.ID = id

	return group, result
}

// DeleteGroup 删除分组
func DeleteGroup(helper dbhelper.DBHelper, id int) bool {
	if id == common_const.SystemAccountGroup.ID {
		return false
	}

	sql := fmt.Sprintf("delete from account_group where id =%d and reserve != 1", id)
	_, result := helper.Execute(sql)

	return result
}

// SaveGroup 保存分组
func SaveGroup(helper dbhelper.DBHelper, group model.GroupDetail) (model.GroupDetail, bool) {
	// modify
	sql := fmt.Sprintf("update account_group set name ='%s', description='%s', catalog=%d where id=%d", group.Name, group.Description, group.Catalog, group.ID)

	_, ret := helper.Execute(sql)
	return group, ret
}
