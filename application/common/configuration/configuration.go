package configuration

/*
实现common.Configuration接口

应用端通过System获取接口对象
*/

import (
	"log"

	"muidea.com/magicCenter/application/common/configuration/bll"
	"muidea.com/magicCenter/application/common/model"
)

// Configuration 配置信息
type Configuration interface {
	GetOption(name string) (string, bool)
	SetOption(name, value string) bool
}

// CreateConfiguration 创建Configuration
func CreateConfiguration() Configuration {
	impl := &impl{}
	impl.configInfoMap = map[string]string{}

	impl.loadConfig()
	return impl
}

type impl struct {
	configInfoMap map[string]string
}

// LoadConfig 加载系统配置信息
func (instance *impl) loadConfig() {
	log.Println("configuration initialize ...")

	keys := []string{model.AppName, model.AppDescription, model.AppDomain, model.AppLogo, model.MailServer, model.MailAccount, model.MailPassword, model.SysDefaultModule}

	instance.configInfoMap = bll.GetConfigurations(keys)
}

// UpdateSystemInfo 更新系统信息
func (instance *impl) UpdateSystemInfo(info model.SystemInfo) bool {
	configs := map[string]string{}
	configs[model.AppName] = info.Name
	configs[model.AppDescription] = info.Description
	configs[model.AppDomain] = info.Domain
	configs[model.AppLogo] = info.Logo
	configs[model.MailServer] = info.MailServer
	configs[model.MailAccount] = info.MailAccount
	configs[model.MailPassword] = info.MailPassword

	instance.configInfoMap[model.AppName] = info.Name
	instance.configInfoMap[model.AppDescription] = info.Description
	instance.configInfoMap[model.AppDomain] = info.Domain
	instance.configInfoMap[model.AppLogo] = info.Logo
	instance.configInfoMap[model.MailServer] = info.MailServer
	instance.configInfoMap[model.MailAccount] = info.MailAccount
	instance.configInfoMap[model.MailPassword] = info.MailPassword

	return bll.UpdateConfigurations(configs)
}

// GetSystemInfo 获取系统信息
func (instance *impl) GetSystemInfo() model.SystemInfo {
	info := model.SystemInfo{}
	info.Name = instance.configInfoMap[model.AppName]
	info.Description = instance.configInfoMap[model.AppDescription]
	info.Domain = instance.configInfoMap[model.AppDomain]
	info.Logo = instance.configInfoMap[model.AppLogo]
	info.MailServer = instance.configInfoMap[model.MailServer]
	info.MailAccount = instance.configInfoMap[model.MailAccount]
	info.MailPassword = instance.configInfoMap[model.MailPassword]

	return info
}

// GetOption 获取指定的配置项
func (instance *impl) GetOption(name string) (string, bool) {
	value, found := instance.configInfoMap[name]
	if !found {
		return bll.GetConfiguration(name)
	}

	return value, found
}

// SetOption 设置指定配置项
func (instance *impl) SetOption(name, value string) bool {
	// 如果值没有变化则直接返回成功
	oldValue, found := instance.configInfoMap[name]
	if found && oldValue == value {
		return true
	}

	if bll.UpdateConfiguration(name, value) {
		if found {
			// 如果之前已经在内存中Load过了，这里也需要把内存中得信息更新一下
			instance.configInfoMap[name] = value
		}
		return true
	}

	return false
}
