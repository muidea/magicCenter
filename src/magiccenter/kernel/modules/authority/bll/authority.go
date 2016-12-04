package bll

import (
	"magiccenter/kernel/modules/authority/dal"
	"magiccenter/system"
)

// QueryAuthGroups 查询aid所应用的gid
func QueryAuthGroups(aid int) []int {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.QueryAuthGroups(helper, aid)
}

// InsertAuthGroup 新增权限
func InsertAuthGroup(aid, gid int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.InsertAuthGroup(helper, aid, gid)
}

// DeleteAuthGroup 删除权限
func DeleteAuthGroup(aid, gid int) bool {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.DeleteAuthGroup(helper, aid, gid)
}
