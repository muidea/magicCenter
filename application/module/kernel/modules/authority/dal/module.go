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
