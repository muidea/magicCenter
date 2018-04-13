package configuration

/*
实现common.Configuration接口

应用端通过System获取接口对象
*/

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCommon/model"
)

var systemConfig common.Configuration

// GetSystemConfiguration 获取SystemConfiguration
func GetSystemConfiguration() common.Configuration {
	if systemConfig == nil {
		LoadSystemConfiguration()
	}

	return systemConfig
}

// LoadSystemConfiguration 加载SystemConfig
func LoadSystemConfiguration() {
	if systemConfig != nil {
		return
	}

	cfg := GetConfiguration("SystemInternalConfig")
	keys := []string{model.AppName, model.AppDescription, model.AppDomain, model.AppLogo, model.MailServer, model.MailAccount, model.MailPassword, model.SysDefaultModule, model.UploadPath}
	cfg.LoadConfig(keys)

	systemConfig = cfg
}

// UpdateSystemProperty 更新系统信息
func UpdateSystemProperty(info model.SystemProperty) bool {
	configs := map[string]string{}
	configs[model.AppName] = info.Name
	configs[model.AppDescription] = info.Description
	configs[model.AppDomain] = info.Domain
	configs[model.AppLogo] = info.Logo
	configs[model.MailServer] = info.MailServer
	configs[model.MailAccount] = info.MailAccount
	configs[model.MailPassword] = info.MailPassword

	return systemConfig.UpdateOptions(configs)
}

// GetSystemProperty 获取系统信息
func GetSystemProperty() model.SystemProperty {
	info := model.SystemProperty{}
	val, ok := systemConfig.GetOption(model.AppName)
	if ok {
		info.Name = val
	}
	val, ok = systemConfig.GetOption(model.AppDescription)
	if ok {
		info.Description = val
	}
	val, ok = systemConfig.GetOption(model.AppDomain)
	if ok {
		info.Domain = val
	}
	val, ok = systemConfig.GetOption(model.AppLogo)
	if ok {
		info.Logo = val
	}
	val, ok = systemConfig.GetOption(model.MailServer)
	if ok {
		info.MailServer = val
	}
	val, ok = systemConfig.GetOption(model.MailAccount)
	if ok {
		info.MailAccount = val
	}
	val, ok = systemConfig.GetOption(model.MailPassword)
	if ok {
		info.MailPassword = val
	}

	return info
}
