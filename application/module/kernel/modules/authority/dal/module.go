package dal

import (
	"fmt"
	"log"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCommon/model"
)

// QueryAllModuleUser 获取全部模块用户信息
func QueryAllModuleUser(helper dbhelper.DBHelper) []model.ModuleUserInfo {
	retValue := []model.ModuleUserInfo{}

	sql := fmt.Sprintf("select distinct(module) from authority_module")
	helper.Query(sql)
	for helper.Next() {
		val := model.ModuleUserInfo{}
		helper.GetValue(&val.Module)
		retValue = append(retValue, val)
	}
	helper.Finish()

	for idx := range retValue {
		val := &retValue[idx]
		sql = fmt.Sprintf("select user from authority_module where module ='%s'", val.Module)
		helper.Query(sql)
		for helper.Next() {
			user := -1
			helper.GetValue(&user)
			val.User = append(val.User, user)
		}
		helper.Finish()
	}

	return retValue
}

// QueryModuleUser 查询拥有指定Module的User
func QueryModuleUser(helper dbhelper.DBHelper, module string) []int {
	retValue := []int{}
	sql := fmt.Sprintf("select distinct(user) from authority_module where module='%s'", module)
	helper.Query(sql)
	defer helper.Finish()
	for helper.Next() {
		val := -1
		helper.GetValue(&val)
		retValue = append(retValue, val)
	}

	return retValue
}

// QueryModuleUserAuthGroup 查询拥有指定Module的User授权组
func QueryModuleUserAuthGroup(helper dbhelper.DBHelper, module string) []model.UserAuthGroup {
	retValue := []model.UserAuthGroup{}
	sql := fmt.Sprintf("select user, authgroup from authority_module where module='%s'", module)
	helper.Query(sql)
	defer helper.Finish()
	for helper.Next() {
		val := model.UserAuthGroup{}
		helper.GetValue(&val.User, &val.AuthGroup)
		retValue = append(retValue, val)
	}

	return retValue
}

// UpdateModuleUserAuthGroup 更新指定Module的用户的授权组
func UpdateModuleUserAuthGroup(helper dbhelper.DBHelper, module string, userAuthGroup []model.UserAuthGroup) bool {
	retVal := false

	log.Printf("module:%s, authGroup size:%d", module, len(userAuthGroup))
	helper.BeginTransaction()
	sql := fmt.Sprintf("delete from authority_module where module='%s'", module)
	_, retVal = helper.Execute(sql)
	if retVal {
		for _, v := range userAuthGroup {
			sql := fmt.Sprintf("insert into authority_module (user, module, authgroup) values (%d,'%s', %d)", v.User, module, v.AuthGroup)

			num, ok := helper.Execute(sql)
			retVal = (num == 1 && ok)
			if !retVal {
				break
			}
		}
	}

	if retVal {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return retVal
}

// QueryAllUserModule 获取全部用户模块信息
func QueryAllUserModule(helper dbhelper.DBHelper) []model.UserModuleInfo {
	retValue := []model.UserModuleInfo{}

	sql := fmt.Sprintf("select distinct(user) from authority_module")
	helper.Query(sql)
	for helper.Next() {
		val := model.UserModuleInfo{}
		helper.GetValue(&val.User)
		retValue = append(retValue, val)
	}
	helper.Finish()

	for idx := range retValue {
		val := &retValue[idx]
		sql = fmt.Sprintf("select module from authority_module where user =%d", val.User)
		helper.Query(sql)
		for helper.Next() {
			mod := ""
			helper.GetValue(&mod)
			val.Module = append(val.Module, mod)
		}
		helper.Finish()
	}

	return retValue
}

// QueryUserModuleAuthGroup 获取指定用户拥有的模块
func QueryUserModuleAuthGroup(helper dbhelper.DBHelper, user int) []model.ModuleAuthGroup {
	retValue := []model.ModuleAuthGroup{}

	sql := fmt.Sprintf("select module, authgroup from authority_module where user=%d", user)
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		val := model.ModuleAuthGroup{}
		helper.GetValue(&val.Module, &val.AuthGroup)

		retValue = append(retValue, val)
	}

	return retValue
}

// UpdateUserModuleAuthGroup 更新指定用户拥有的模块
func UpdateUserModuleAuthGroup(helper dbhelper.DBHelper, user int, moduleAuthGroups []model.ModuleAuthGroup) bool {
	retVal := false

	helper.BeginTransaction()
	sql := fmt.Sprintf("delete from authority_module where user=%d", user)
	_, retVal = helper.Execute(sql)
	if retVal {
		for _, v := range moduleAuthGroups {
			sql := fmt.Sprintf("insert into authority_module (user, module, authgroup) values (%d,'%s', %d)", user, v.Module, v.AuthGroup)

			num, ok := helper.Execute(sql)
			retVal = (num == 1 && ok)
			if !retVal {
				break
			}
		}
	}

	if retVal {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return retVal
}
