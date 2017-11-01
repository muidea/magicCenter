package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
)

//SetOption 保存配置项
func SetOption(helper dbhelper.DBHelper, owner, key, value string) bool {
	sql := fmt.Sprintf("select id, value from common_option where `key`='%s' and owner='%s'", key, owner)
	helper.Query(sql)

	id := -1
	oldValue := ""
	found := false
	if helper.Next() {
		helper.GetValue(&id, &oldValue)
		found = true
	}

	if value == oldValue {
		return true
	}

	if found {
		sql = fmt.Sprintf("update common_option set value='%s' where id=%d and owner='%s'", value, id, owner)
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

	value := ""
	found := false
	if helper.Next() {
		helper.GetValue(&value)
		found = true
	}

	return value, found
}
