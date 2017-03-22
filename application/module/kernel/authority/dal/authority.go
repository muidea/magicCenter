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
