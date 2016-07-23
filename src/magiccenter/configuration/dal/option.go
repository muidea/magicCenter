package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
)

//SetOption 保存配置项
func SetOption(helper modelhelper.Model, key, value string) bool {
	sql := fmt.Sprintf("select id, value from `option` where `key`='%s'", key)
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
		sql = fmt.Sprintf("update `option` set value='%s' where id=%d", value, id)
	} else {
		sql = fmt.Sprintf("insert into `option`(`key`,value) values ('%s','%s')", key, value)
	}

	num, ret := helper.Execute(sql)
	return num == 1 && ret
}

// GetOption 获取配置项
func GetOption(helper modelhelper.Model, key string) (string, bool) {
	sql := fmt.Sprintf("select value from `option` where `key`='%s'", key)
	helper.Query(sql)

	value := ""
	found := false
	if helper.Next() {
		helper.GetValue(&value)
		found = true
	}

	return value, found
}
