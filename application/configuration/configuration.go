package configuration

/*
实现common.Configuration接口

应用端通过System获取接口对象
*/

import (
	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/configuration/bll"
)

type impl struct {
	configInfoMap map[string]string
}

// CreateConfiguration 创建Configuration
func CreateConfiguration() common.Configuration {
	impl := &impl{}
	impl.configInfoMap = map[string]string{}

	return impl
}

// LoadConfig 加载系统配置信息
func (instance *impl) LoadConfig() {
	log.Println("configuration initialize ...")

	keys := []string{common.AppName, common.AppDomain, common.AppLogo, common.MailServer, common.MailAccount, common.MailPassword, common.SysDefaultModule}

	instance.configInfoMap = bll.GetConfigurations(keys)
}

// UpdateSystemInfo 更新系统信息
func (instance *impl) UpdateSystemInfo(info common.SystemInfo) bool {
	configs := map[string]string{}
	configs[common.AppName] = info.Name
	configs[common.AppDomain] = info.Domain
	configs[common.AppLogo] = info.Logo
	configs[common.MailServer] = info.MailServer
	configs[common.MailAccount] = info.MailAccount
	configs[common.MailPassword] = info.MailPassword

	instance.configInfoMap[common.AppName] = info.Name
	instance.configInfoMap[common.AppDomain] = info.Domain
	instance.configInfoMap[common.AppLogo] = info.Logo
	instance.configInfoMap[common.MailServer] = info.MailServer
	instance.configInfoMap[common.MailAccount] = info.MailAccount
	instance.configInfoMap[common.MailPassword] = info.MailPassword

	return bll.UpdateConfigurations(configs)
}

// GetSystemInfo 获取系统信息
func (instance *impl) GetSystemInfo() common.SystemInfo {
	info := common.SystemInfo{}
	info.Name = instance.configInfoMap[common.AppName]
	info.Domain = instance.configInfoMap[common.AppDomain]
	info.Logo = instance.configInfoMap[common.AppLogo]
	info.MailServer = instance.configInfoMap[common.MailServer]
	info.MailAccount = instance.configInfoMap[common.MailAccount]
	info.MailPassword = instance.configInfoMap[common.MailPassword]

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
