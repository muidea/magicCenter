package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

// QueryUserModule 获取指定用户拥有的模块
func QueryUserModule(helper dbhelper.DBHelper, user int) model.UserModuleAuthGroupInfo {
	retValue := model.UserModuleAuthGroupInfo{User: user}

	sql := fmt.Sprintf("select module, authgroup from authority_module where user=%d", user)
	helper.Query(sql)
	for helper.Next() {
		moduleAuthGroup := model.ModuleAuthGroup{}
		helper.GetValue(&moduleAuthGroup.Module, &moduleAuthGroup.AuthGroup)

		retValue.ModuleAuthGroups = append(retValue.ModuleAuthGroups, moduleAuthGroup)
	}

	return retValue
}

// UpdateUserModule 更新指定用户拥有的模块
func UpdateUserModule(helper dbhelper.DBHelper, user int, moduleAuthGroups []model.ModuleAuthGroup) bool {
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

// QueryModuleUser 查询拥有指定Module的User
func QueryModuleUser(helper dbhelper.DBHelper, module string) model.ModuleUserAuthGroupInfo {
	retValue := model.ModuleUserAuthGroupInfo{Module: module}
	sql := fmt.Sprintf("select user, authgroup from authority_module where module='%s'", module)
	helper.Query(sql)
	for helper.Next() {
		userAuthGroup := model.UserAuthGroup{}
		helper.GetValue(&userAuthGroup.User, &userAuthGroup.AuthGroup)
		retValue.UserAuthGroups = append(retValue.UserAuthGroups, userAuthGroup)
	}

	return retValue
}

// UpdateModuleUser 更新指定Module的拥有着
func UpdateModuleUser(helper dbhelper.DBHelper, module string, userAuthGroup []model.UserAuthGroup) bool {
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
