package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/foundation/util"
)

// QueryUserAuthGroup 获取指定用户的授权组
func QueryUserAuthGroup(helper dbhelper.DBHelper, user int) []int {
	retValue := []int{}

	groups := ""
	sql := fmt.Sprintf("select groups from authority_authgroup where user=%d", user)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&groups)

		retValue, _ = util.Str2IntArray(groups)
	}

	return retValue
}

// UpdateUserAuthGroup 更新指定用户的授权组
func UpdateUserAuthGroup(helper dbhelper.DBHelper, user int, authGroups []int) bool {
	retVal := false

	sql := fmt.Sprintf("select id from authority_authgroup where user=%d", user)
	helper.Query(sql)
	if helper.Next() {
		id := -1
		helper.GetValue(&id)

		groups := util.IntArray2Str(authGroups)
		sql = fmt.Sprintf("update authority_authgroup set groups='%s' where id=%d", groups, id)
		num, ok := helper.Execute(sql)
		retVal = (num == 1 && ok)
	}

	return retVal
}
