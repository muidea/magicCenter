package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
)

// UpateUserAuthorityGroup 更新用户权限组
func UpateUserAuthorityGroup(helper dbhelper.DBHelper, userID int, authGroup []int) bool {
	helper.BeginTransaction()
	sql := fmt.Sprintf("delete from authority where uid=%d", userID)
	_, ok := helper.Execute(sql)
	if !ok {
		return false
	}

	for _, val := range authGroup {
		sql = fmt.Sprintf("insert into authority (uid, gid) values (%d,%d)", userID, val)
		_, ok = helper.Execute(sql)
		if !ok {
			break
		}
	}

	if ok {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return ok
}

// GetUserAuthorityGroup 获取指定用户的授权组
func GetUserAuthorityGroup(helper dbhelper.DBHelper, userID int) []int {
	authGroup := []int{}
	sql := fmt.Sprintf("select gid from authority where uid=%d", userID)
	helper.Query(sql)
	for helper.Next() {
		id := 0
		helper.GetValue(&id)
		authGroup = append(authGroup, id)
	}

	return authGroup
}
