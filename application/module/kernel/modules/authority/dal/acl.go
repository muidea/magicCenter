package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/util"
)

// InsertACL 新增ACL记录
func InsertACL(helper dbhelper.DBHelper, url, method, module string, status int, authGroups int) (model.ACL, bool) {
	acl := model.ACL{URL: url, Method: method, Module: module, Status: status, AuthGroup: authGroups}
	sql := fmt.Sprintf("insert into authority_acl (url, method, module, status, authgroup) values ('%s','%s','%s',%d,%d)", url, method, module, status, acl.AuthGroup)
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

// UpdateACLStatus 更新ACL状态
func UpdateACLStatus(helper dbhelper.DBHelper, enableList []int, disableList []int) bool {
	if len(enableList) == 0 && len(disableList) == 0 {
		return true
	}

	enableOK := true
	disableOK := true
	helper.BeginTransaction()
	if len(enableList) > 0 {
		str := util.IntArray2Str(enableList)
		sql := fmt.Sprintf("update authority_acl set status=1 where id in(%s)", str)
		_, enableOK = helper.Execute(sql)
	}
	if len(disableList) > 0 {
		str := util.IntArray2Str(disableList)
		sql := fmt.Sprintf("update authority_acl set status=0 where id in(%s)", str)
		_, disableOK = helper.Execute(sql)
	}

	if enableOK && disableOK {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return enableOK && disableOK
}

// QueryACLByID 查询指定的ACL
func QueryACLByID(helper dbhelper.DBHelper, id int) (model.ACL, bool) {
	acl := model.ACL{}
	retVal := false

	sql := fmt.Sprintf("select id, url, method, module, status, authgroup from authority_acl where id=%d", id)
	helper.Query(sql)
	if helper.Next() {
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &acl.AuthGroup)
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
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &acl.AuthGroup)
		retVal = true
	}

	return acl, retVal
}

// QueryAvalibleACLByAuthGroup 查询指定AuthGroup可管理的ACL
func QueryAvalibleACLByAuthGroup(helper dbhelper.DBHelper, authGroup int) []model.ACL {
	acls := []model.ACL{}
	sql := fmt.Sprintf("select id, url, method, module, status, authgroup from authority_acl where authgroup <= %d", authGroup)
	helper.Query(sql)
	for helper.Next() {
		acl := model.ACL{}
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &acl.AuthGroup)
		acls = append(acls, acl)
	}

	return acls
}

// UpateACL 更新ACL记录
func UpateACL(helper dbhelper.DBHelper, acl model.ACL) bool {
	sql := fmt.Sprintf("update authority_acl set authgroup=%d, status=%d where id=%d", acl.AuthGroup, acl.Status, acl.ID)
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
		helper.GetValue(&acl.ID, &acl.URL, &acl.Method, &acl.Module, &acl.Status, &acl.AuthGroup)
		acls = append(acls, acl)
	}

	return acls
}
