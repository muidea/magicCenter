package configuration

/*
实现common.Configuration接口

应用端通过System获取接口对象
*/

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration/bll"
)

func init() {
	localConfigurationMap = make(map[string]common.Configuration)
}

// GetConfiguration 获取指定的Configuration
func GetConfiguration(id string) common.Configuration {
	cfg, found := localConfigurationMap[id]
	if !found {
		cfg = createConfiguration(id)
		localConfigurationMap[id] = cfg
	}

	return cfg
}

var localConfigurationMap map[string]common.Configuration

// CreateConfiguration 创建Configuration
func createConfiguration(id string) common.Configuration {
	s := &impl{id: id}
	s.configProperityMap = map[string]string{}

	return s
}

type impl struct {
	id                 string
	configProperityMap map[string]string
}

func (instance *impl) ID() string {
	return instance.id
}

func (instance *impl) LoadConfig(items []string) {
	instance.configProperityMap = bll.GetConfigurations(instance.id, items)
}

// GetOption 获取指定的配置项
func (instance *impl) GetOption(name string) (string, bool) {
	value, found := instance.configProperityMap[name]
	if !found {
		return bll.GetConfiguration(instance.id, name)
	}

	return value, found
}

// SetOption 设置指定配置项
func (instance *impl) SetOption(name, value string) bool {
	// 如果值没有变化则直接返回成功
	oldValue, found := instance.configProperityMap[name]
	if found && oldValue == value {
		return true
	}

	if bll.UpdateConfiguration(instance.id, name, value) {
		if found {
			// 如果之前已经在内存中Load过了，这里也需要把内存中得信息更新一下
			instance.configProperityMap[name] = value
		}
		return true
	}

	return false
}

func (instance *impl) UpdateOptions(opts map[string]string) bool {
	ret := bll.UpdateConfigurations(instance.id, opts)
	if ret {
		for k, v := range opts {
			instance.configProperityMap[k] = v
		}
	}

	return ret
}
