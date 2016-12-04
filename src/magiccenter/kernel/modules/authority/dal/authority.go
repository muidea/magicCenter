package dal

import (
	"fmt"
	"magiccenter/common"
)

// QueryAuthGroups 查询aid所应用的gid
func QueryAuthGroups(helper common.DBHelper, aid int) []int {
	gids := []int{}
	sql := fmt.Sprintf("select distinct gid from `authority` where uid=%d", aid)
	helper.Query(sql)

	for helper.Next() {
		g := 0
		helper.GetValue(&g)

		gids = append(gids, g)
	}

	return gids
}

// InsertAuthGroup 新建权限
func InsertAuthGroup(helper common.DBHelper, aid, gid int) bool {
	sql := fmt.Sprintf("insert into `authority`(aid, gid) values(%d, %d) ", aid, gid)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

// DeleteAuthGroup 删除权限
func DeleteAuthGroup(helper common.DBHelper, aid, gid int) bool {
	sql := fmt.Sprintf("delete `authority` where aid = %d and gid = %d", aid, gid)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}
