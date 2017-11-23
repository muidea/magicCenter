package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
)

// QueryUserAuthGroup 获取指定用户的授权组
func QueryUserAuthGroup(helper dbhelper.DBHelper, user int) int {
	retValue := common.VisitorAuthGroup.ID

	sql := fmt.Sprintf("select authgroup from authority_authgroup where user=%d", user)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&retValue)
	}

	return retValue
}

// UpdateUserAuthGroup 更新指定用户的授权组
func UpdateUserAuthGroup(helper dbhelper.DBHelper, user int, authGroup int) bool {
	retVal := false

	sql := fmt.Sprintf("select id from authority_authgroup where user=%d", user)
	helper.Query(sql)
	if helper.Next() {
		id := -1
		helper.GetValue(&id)

		sql = fmt.Sprintf("update authority_authgroup set authgroup=%d where id=%d", authGroup, id)
		num, ok := helper.Execute(sql)
		retVal = (num == 1 && ok)
	}

	return retVal
}
