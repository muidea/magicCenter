package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/util"
)

// InsertACL 新增ACL记录
func InsertACL(helper dbhelper.DBHelper, url, method, module string, status int) (model.ACL, bool) {
	acl := model.ACL{URL: url, Method: method, Module: module, Status: status, AuthGroup: []int{}}
	sql := fmt.Sprintf("insert into acl (url, method, module, status, authgroup) values ('%s','%s','%s',%d,'%s')", url, method, module, status, util.IntArray2Str(acl.AuthGroup))
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

// EnableACL 启用ACL
func EnableACL(helper dbhelper.DBHelper, ids []int) bool {
	str := util.IntArray2Str(ids)
	sql := fmt.Sprintf("update acl set status=1 where id in(%s)", str)
	num, ok := helper.Execute(sql)

	return ok && (int(num) == len(ids))
}

// DisableACL 禁用ACL
func DisableACL(helper dbhelper.DBHelper, ids []int) bool {
	str := util.IntArray2Str(ids)
	sql := fmt.Sprintf("update acl set status=0 where id in(%s)", str)
	num, ok := helper.Execute(sql)

	return ok && (int(num) == len(ids))
}

// UpateACL 更新ACL记录
func UpateACL(helper dbhelper.DBHelper, acl model.ACL) bool {
	authGroup := util.IntArray2Str(acl.AuthGroup)

	sql := fmt.Sprintf("update acl set authgroup='%s', status=%d where id=%d", authGroup, acl.Status, acl.ID)
	num, ok := helper.Execute(sql)
	return ok && num == 1
}

// QueryAllACL 查询全部ACL信息
func QueryAllACL(helper dbhelper.DBHelper, status int) []model.ACL {
	acls := []model.ACL{}
	sql := ""
	if status == -1 {
		sql = fmt.Sprint("select id, url, method, module, status, authgroup from acl order by module, method")
	} else {
		sql = fmt.Sprintf("select id, url, method, module, status, authgroup from acl where status=%d order by module, method", status)
	}

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

// QueryACL 查询指定Module的ACL信息
func QueryACL(helper dbhelper.DBHelper, module string, status int) []model.ACL {
	acls := []model.ACL{}
	sql := ""
	if status == -1 {
		sql = fmt.Sprintf("select id, url, method, module, status, authgroup from acl where module='%s'", module)
	} else {
		sql = fmt.Sprintf("select id, url, method, module, status, authgroup from acl where module='%s' and status=%d", module, status)
	}

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

// LoadACL 加载所有ACL
func LoadACL(helper dbhelper.DBHelper, method string) []model.ACL {
	acls := []model.ACL{}
	sql := fmt.Sprintf("select id, url, method, module, status, authgroup from acl where method='%s'", method)
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
