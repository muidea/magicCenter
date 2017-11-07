package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
)

// QueryUserModule 获取指定用户拥有的模块
func QueryUserModule(helper dbhelper.DBHelper, user int) []string {
	retValue := []string{}

	sql := fmt.Sprintf("select module from authority_module where user=%d", user)
	helper.Query(sql)
	for helper.Next() {
		module := ""
		helper.GetValue(&module)

		retValue = append(retValue, module)
	}

	return retValue
}

// UpdateUserModule 更新指定用户拥有的模块
func UpdateUserModule(helper dbhelper.DBHelper, user int, modules []string) bool {
	retVal := false

	helper.BeginTransaction()
	sql := fmt.Sprintf("delete from authority_module where user=%d", user)
	_, retVal = helper.Execute(sql)
	if retVal {
		for _, v := range modules {
			sql := fmt.Sprintf("insert into authority_module (user, module) values (%d,'%s')", user, v)

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
func QueryModuleUser(helper dbhelper.DBHelper, module string) []int {
	ids := []int{}
	sql := fmt.Sprintf("select user from authority_module where module='%s'", module)
	helper.Query(sql)
	for helper.Next() {
		id := -1
		helper.GetValue(&id)
		ids = append(ids, id)
	}

	return ids
}

// UpdateModuleUser 更新指定Module的拥有着
func UpdateModuleUser(helper dbhelper.DBHelper, module string, users []int) bool {
	retVal := false

	helper.BeginTransaction()
	sql := fmt.Sprintf("delete from authority_module where module='%s'", module)
	_, retVal = helper.Execute(sql)
	if retVal {
		for _, v := range users {
			sql := fmt.Sprintf("insert into authority_module (user, module) values (%d,'%s')", v, module)

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
