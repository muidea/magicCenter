package system

import (
	"webcenter/util/modelhelper"
)

func UpdateSystemName(model modelhelper.Model, systemName string) bool {
	return SetOption(model, "@systemName", systemName)
}

func GetSystemName(model modelhelper.Model) (string, bool) {
	return GetOption(model, "@systemName")
}

func UpdateSystemLogo(helper modelhelper.Model, systemLogo string) bool {
	return SetOption(helper, "@systemLogo", systemLogo)
}

func GetSystemLogo(model modelhelper.Model) (string, bool) {
	return GetOption(helper, "@systemLogo")
}

func UpdateSystemDomain(helper modelhelper.Model, systemDomain string) bool {
	return SetOption(helper, "@systemDomain", systemDomain)
}

func GetSystemDomain(helper modelhelper.Model) (string, bool) {
	return GetOption(helper, "@systemDomain")
}

func UpdateSystemEMailServer(helper modelhelper.Model, systemEMailServer string) bool {
	return SetOption(helper, "@systemEMailServer", systemEMailServer)
}

func GetSystemEMailServer(helper modelhelper.Model) (string, bool) {
	return GetOption(helper, "@systemEMailServer")
}

func UpdateSystemEMailAccount(helper modelhelper.Model, systemEMail string) bool {
	return SetOption(helper, "@systemEMailAccount", systemEMail)
}

func GetSystemEMailAccount(helper modelhelper.Model) (string, bool) {
	return GetOption(helper, "@systemEMailAccount")
}

func UpdateSystemEMailPassword(helper modelhelper.Model, systemEMail string) bool {
	return SetOption(helper, "@systemEMailPassword", systemEMail)
}

func GetSystemEMailPassword(helper modelhelper.Model) (string, bool) {
	return GetOption(helper, "@systemEMailPassword")
}



