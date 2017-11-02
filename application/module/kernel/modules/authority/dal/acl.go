package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/util"
)

// InsertACL 新增ACL记录
func InsertACL(helper dbhelper.DBHelper, url, method, module string, status int, authGroups []int) (model.ACL, bool) {
	acl := model.ACL{URL: url, Method: method, Module: module, Status: status, AuthGroup: authGroups}
	sql := fmt.Sprintf("insert into authority_acl (url, method, module, status, authgroup) values ('%s','%s','%s',%d,'%s')", url, method, module, status, util.IntArray2Str(acl.AuthGroup))
	num, ok := helper.Execute(sql)
	if !ok || num != 1 {
		return acl, false
	}

	ok = false
	sql = fmt.Sprintf("select id from authority_acl where url='%s'", url)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&acl.ID)
		ok = true
	}

	return acl, ok
}

// DeleteACL 删除ACL记录
func DeleteACL(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf("delete from authority_acl where id=%d", id)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// EnableACL 启用ACL
func EnableACL(helper dbhelper.DBHelper, ids []int) bool {
	if len(ids) == 0 {
		return true
	}

	str := util.IntArray2Str(ids)
	sql := fmt.Sprintf("update authority_acl set status=1 where id in(%s)", str)
	num, ok := helper.Execute(sql)

	return ok && (int(num) == len(ids))
}

// DisableACL 禁用ACL
func DisableACL(helper dbhelper.DBHelper, ids []int) bool {
	if len(ids) == 0 {
		return true
	}

	str := util.IntArray2Str(ids)
	sql := fmt.Sprintf("update authority_acl set status=0 where id in(%s)", str)
	num, ok := helper.Execute(sql)

	return ok && (int(num) == len(ids))
}

// QueryACLByID 查询指定的ACL
func QueryACLByID(helper dbhelper.DBHelper, id int) (model.ACL, bool) {
	acl := model.ACL{}
	retVal := false

	sql := fmt.Sprintf("select id, url, method, module, status, authgroup from authority_acl where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		authGroups := ""
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &authGroups)
		acl.AuthGroup, _ = util.Str2IntArray(authGroups)
		retVal = true
	}

	return acl, retVal
}

// QueryACL 查询指定的ACL
func QueryACL(helper dbhelper.DBHelper, url, method string) (model.ACL, bool) {
	acl := model.ACL{}
	retVal := false

	sql := fmt.Sprintf("select id, url, method, module, status, authgroup from authority_acl where url='%s' and method='%s'", url, method)
	helper.Query(sql)
	if helper.Next() {
		authGroups := ""
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &authGroups)
		acl.AuthGroup, _ = util.Str2IntArray(authGroups)
		retVal = true
	}

	return acl, retVal
}

// UpateACL 更新ACL记录
func UpateACL(helper dbhelper.DBHelper, acl model.ACL) bool {
	authGroup := util.IntArray2Str(acl.AuthGroup)

	sql := fmt.Sprintf("update authority_acl set authgroup='%s', status=%d where id=%d", authGroup, acl.Status, acl.ID)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// QueryModuleACL 查询指定Module的ACL信息
func QueryModuleACL(helper dbhelper.DBHelper, module string) []model.ACL {
	acls := []model.ACL{}
	sql := fmt.Sprintf("select id, url, method, module, status, authgroup from authority_acl where module='%s'", module)

	helper.Query(sql)
	for helper.Next() {
		acl := model.ACL{}
		authGroups := ""
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &authGroups)
		acl.AuthGroup, _ = util.Str2IntArray(authGroups)
		acls = append(acls, acl)
	}

	return acls
}
