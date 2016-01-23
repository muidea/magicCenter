package system

import (
	"fmt"
	"webcenter/util/modelhelper"
)

func SetOption(model modelhelper.Model, key, value string) bool {
	sql := fmt.Sprintf("select id from system_config where `key`='%s'", key)
	if !model.Query(sql) {
		panic("query failed")
	}

	id := -1
	found := false
	for model.Next() {
		model.GetValue(&id)
		found = true
		break
	}
	
	if found {
		sql = fmt.Sprintf("update system_config set value='%s' where id=%d", value, id)		
	} else {
		sql = fmt.Sprintf("insert into system_config(`key`,value) values ('%s','%s')", key, value)
	}
	
	if !model.Execute(sql) {
		return false
	}

	return true
}

func GetOption(model modelhelper.Model, key string) (string, bool) {
	sql := fmt.Sprintf("select value from system_config where `key`='%s'", key)
	if !model.Query(sql) {
		panic("query failed")
	}

	value := ""
	found := false
	for model.Next() {
		model.GetValue(&value)
		found = true
		break
	}
	
	return value, found
}


func UpdateSystemName(model modelhelper.Model, systemName string) bool {
	return SetOption(model, "@systemName", systemName)
}

func GetSystemName(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemName")
}

func UpdateSystemLogo(model modelhelper.Model, systemLogo string) bool {
	return SetOption(model, "@systemLogo", systemLogo)
}

func GetSystemLogo(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemLogo")
}

func UpdateSystemDomain(model modelhelper.Model, systemDomain string) bool {
	return SetOption(model, "@systemDomain", systemDomain)
}

func GetSystemDomain(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemDomain")
}

func UpdateSystemEMailServer(model modelhelper.Model, systemEMailServer string) bool {
	return SetOption(model, "@systemEMailServer", systemEMailServer)
}

func GetSystemEMailServer(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemEMailServer")
}

func UpdateSystemEMailAccount(model modelhelper.Model, systemEMail string) bool {
	return SetOption(model, "@systemEMailAccount", systemEMail)
}

func GetSystemEMailAccount(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemEMailAccount")
}

func UpdateSystemEMailPassword(model modelhelper.Model, systemEMail string) bool {
	return SetOption(model, "@systemEMailPassword", systemEMail)
}

func GetSystemEMailPassword(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemEMailPassword")
}



