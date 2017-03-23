package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/util"
)

// InsertACL 新怎ACL记录
func InsertACL(helper dbhelper.DBHelper, url string) (model.ACL, bool) {
	acl := model.ACL{URL: url, AuthGroup: []int{}}
	sql := fmt.Sprintf("insert into acl (url) values ('%s')", url)
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return acl, false
	}

	ok = false
	sql = fmt.Sprintf("select id from acl where url='%s'", url)
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

// UpateACL 更新ACL记录
func UpateACL(helper dbhelper.DBHelper, acl model.ACL) bool {
	authGroup := util.IntArray2Str(acl.AuthGroup)
	sql := fmt.Sprintf("update acl set authgroup='%s' where id=%d", authGroup, acl.ID)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// LoadACL 加载所有ACL
func LoadACL(helper dbhelper.DBHelper) []model.ACL {
	acls := []model.ACL{}
	sql := fmt.Sprint("select id, url, authgroup from acl")
	helper.Query(sql)
	for helper.Next() {
		acl := model.ACL{}
		helper.GetValue(&acl.ID, &acl.URL, &acl.AuthGroup)
		acls = append(acls, acl)
	}

	return acls
}
