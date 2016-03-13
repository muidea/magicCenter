package dal

import (
	"fmt"
	"webcenter/util/modelhelper"
)


func SetOption(model modelhelper.Model, key, value string) bool {
	sql := fmt.Sprintf("select id from option where `key`='%s'", key)
	model.Query(sql)

	id := -1
	found := false
	if model.Next() {
		model.GetValue(&id)
		found = true
	}
	
	if found {
		sql = fmt.Sprintf("update option set value='%s' where id=%d", value, id)		
	} else {
		sql = fmt.Sprintf("insert into option(`key`,value) values ('%s','%s')", key, value)
	}
	
	num, ret := model.Execute(sql)
	return num == 1 && ret
}

func GetOption(model modelhelper.Model, key string) (string, bool) {
	sql := fmt.Sprintf("select value from option where `key`='%s'", key)
	model.Query(sql)

	value := ""
	found := false
	if model.Next() {
		model.GetValue(&value)
		found = true
	}
	
	return value, found
}

