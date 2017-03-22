package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

// InsertACL 新怎ACL记录
func InsertACL(helper dbhelper.DBHelper, url string, action int) (model.ACL, bool) {
	acl := model.ACL{URL: url, Action: action}
	sql := fmt.Sprintf("insert into acl (url, action) values ('%s', %d)", url, action)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return acl, false
	}

	ok = false
	sql = fmt.Sprintf("select id from acl where url='%s' and action=%d", url, action)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&acl.ID)
		ok = true
	}

	return acl, ok
}

// DeleteACL 删除ACL记录
func DeleteACL(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from acl where id=%d", id)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}
