package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

// QueryUserModuleAuthGroup 获取指定用户拥有的模块
func QueryUserModuleAuthGroup(helper dbhelper.DBHelper, user int) model.UserModuleAuthGroup {
	retValue := model.UserModuleAuthGroup{User: user}

	sql := fmt.Sprintf("select module, authgroup from authority_module where user=%d", user)
	helper.Query(sql)
	for helper.Next() {
		moduleAuthGroup := model.ModuleAuthGroup{}
		helper.GetValue(&moduleAuthGroup.Module, &moduleAuthGroup.AuthGroup)

		retValue.ModuleAuthGroup = append(retValue.ModuleAuthGroup, moduleAuthGroup)
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

// QueryModuleUserAuthGroup 查询拥有指定Module的User
func QueryModuleUserAuthGroup(helper dbhelper.DBHelper, module string) model.ModuleUserAuthGroup {
	retValue := model.ModuleUserAuthGroup{Module: module}
	sql := fmt.Sprintf("select user, authgroup from authority_module where module='%s'", module)
	helper.Query(sql)
	for helper.Next() {
		userAuthGroup := model.UserAuthGroup{}
		helper.GetValue(&userAuthGroup.User, &userAuthGroup.AuthGroup)
		retValue.UserAuthGroup = append(retValue.UserAuthGroup, userAuthGroup)
	}

	return retValue
}

// UpdateModuleUserAuthGroup 更新指定Module的拥有着
func UpdateModuleUserAuthGroup(helper dbhelper.DBHelper, module string, userAuthGroup []model.UserAuthGroup) bool {
	retVal := false

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
