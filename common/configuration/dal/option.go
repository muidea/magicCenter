package dal

import (
	"fmt"

	"github.com/muidea/magicCenter/common/dbhelper"
)

//SetOption 保存配置项
func SetOption(helper dbhelper.DBHelper, owner, key, value string) bool {
	oldValue, found := GetOption(helper, owner, key)

	sql := ""
	if found {
		if oldValue == value {
			return true
		}

		sql = fmt.Sprintf("update common_option set value='%s' where key='%s' and owner='%s'", value, key, owner)
	} else {
		sql = fmt.Sprintf("insert into common_option(`key`,value, owner) values ('%s','%s','%s')", key, value, owner)
	}

	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

// GetOption 获取配置项
func GetOption(helper dbhelper.DBHelper, owner, key string) (string, bool) {
	sql := fmt.Sprintf("select value from common_option where `key`='%s' and owner='%s'", key, owner)
	helper.Query(sql)
	defer helper.Finish()

	value := ""
	found := false
	if helper.Next() {
		helper.GetValue(&value)
		found = true
	}

	return value, found
}
