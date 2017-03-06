package configuration

/*
实现service.Configuration接口

应用端通过System获取接口对象
*/

import (
	"log"

	"muidea.com/magicCenter/application/common/service"
	"muidea.com/magicCenter/application/configuration/bll"
)

type impl struct {
	configInfoMap map[string]string
}

// CreateConfiguration 创建Configuration
func CreateConfiguration() service.Configuration {
	impl := &impl{}
	impl.configInfoMap = map[string]string{}

	return impl
}

// LoadConfig 加载系统配置信息
func (instance *impl) LoadConfig() {
	log.Println("configuration initialize ...")

	keys := []string{service.AppName, service.AppDomain, service.AppLogo, service.MailServer, service.MailAccount, service.MailPassword, service.SysDefaultModule}

	instance.configInfoMap = bll.GetConfigurations(keys)
}

// UpdateSystemInfo 更新系统信息
func (instance *impl) UpdateSystemInfo(info service.SystemInfo) bool {
	configs := map[string]string{}
	configs[service.AppName] = info.Name
	configs[service.AppDomain] = info.Domain
	configs[service.AppLogo] = info.Logo
	configs[service.MailServer] = info.MailServer
	configs[service.MailAccount] = info.MailAccount
	configs[service.MailPassword] = info.MailPassword

	instance.configInfoMap[service.AppName] = info.Name
	instance.configInfoMap[service.AppDomain] = info.Domain
	instance.configInfoMap[service.AppLogo] = info.Logo
	instance.configInfoMap[service.MailServer] = info.MailServer
	instance.configInfoMap[service.MailAccount] = info.MailAccount
	instance.configInfoMap[service.MailPassword] = info.MailPassword

	return bll.UpdateConfigurations(configs)
}

// GetSystemInfo 获取系统信息
func (instance *impl) GetSystemInfo() service.SystemInfo {
	info := service.SystemInfo{}
	info.Name = instance.configInfoMap[service.AppName]
	info.Domain = instance.configInfoMap[service.AppDomain]
	info.Logo = instance.configInfoMap[service.AppLogo]
	info.MailServer = instance.configInfoMap[service.MailServer]
	info.MailAccount = instance.configInfoMap[service.MailAccount]
	info.MailPassword = instance.configInfoMap[service.MailPassword]

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
