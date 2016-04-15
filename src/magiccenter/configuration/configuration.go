package configuration

import (
	"log"
	"magiccenter/configuration/bll"
)

const (
	APP_NAME = "@application_name"
	APP_DOMAIN = "@application_domain"
	APP_LOGO = "@application_logo"
	
	MAIL_SERVER = "@system_mailServer"
	MAIL_ACCOUNT = "@system_mailAccount"
	MAIL_PASSWORD = "@system_mailPassword"	
	SYS_DEFULTMODULE = "@system_defaultModule"
	
	STATIC_PATH = "@system_staticPath"
	RESOURCE_PATH = "@system_staticPath"
	UPLOAD_PATH = "@system_uploadPath"
	
	AUTHORITH_ID = "@authorith_Id"
)

type SystemInfo struct {
	Name string
	Logo string
	Domain string
	MailServer string
	MailAccount string
	MailPassword string
}

var configInfoMap = map[string]string{}

func LoadConfig() {
	log.Println("configuration initialize ...")
	
	keys := [] string {APP_NAME, APP_DOMAIN, APP_LOGO, MAIL_SERVER, MAIL_ACCOUNT, MAIL_PASSWORD, SYS_DEFULTMODULE}
	
	configInfoMap = bll.GetConfiguration(keys)
	
	configInfoMap[STATIC_PATH] = "static"
	configInfoMap[RESOURCE_PATH] = "template"
	configInfoMap[UPLOAD_PATH] = "upload"
	configInfoMap[AUTHORITH_ID] = "@@@$$auth_Id@@@"
}

func UpdateSystemInfo(info SystemInfo) bool {
	configs := map[string]string{}
	configs[APP_NAME] = info.Name
	configs[APP_DOMAIN] = info.Domain
	configs[APP_LOGO] = info.Logo
	configs[MAIL_SERVER] = info.MailServer
	configs[MAIL_ACCOUNT] = info.MailAccount
	configs[MAIL_PASSWORD] = info.MailPassword
	
	configInfoMap[APP_NAME] = info.Name
	configInfoMap[APP_DOMAIN] = info.Domain
	configInfoMap[APP_LOGO] = info.Logo
	configInfoMap[MAIL_SERVER] = info.MailServer
	configInfoMap[MAIL_ACCOUNT] = info.MailAccount
	configInfoMap[MAIL_PASSWORD] = info.MailPassword
	
	return bll.UpdateConfigurations(configs)
}

func GetSystemInfo() SystemInfo {
	info := SystemInfo{}
	info.Name = configInfoMap[APP_NAME]
	info.Domain = configInfoMap[APP_DOMAIN] 
	info.Logo = configInfoMap[APP_LOGO]
	info.MailServer = configInfoMap[MAIL_SERVER]
	info.MailAccount = configInfoMap[MAIL_ACCOUNT]
	info.MailPassword = configInfoMap[MAIL_PASSWORD]
	
	return info
}

func GetOption(name string) (string, bool) {
	value, found := configInfoMap[name]
	
	return value,found	
}

func SetOption(name, value string) bool {
	// 如果值没有变化则直接返回成功
	oldValue, found := configInfoMap[name]
	if found && oldValue == value {
		return true
	} 
	
	if bll.UpdateConfiguration(name, value) {
		configInfoMap[name] = value
		return true
	}
	
	return false
}



