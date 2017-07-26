package bll

import (
	"muidea.com/magicCenter/application/common/configuration/dal"
	"muidea.com/magicCenter/application/common/dbhelper"
)

// UpdateConfigurations 更新配置集
func UpdateConfigurations(owner string, configs map[string]string) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	result := true
	helper.BeginTransaction()
	for k, v := range configs {
		if !dal.SetOption(helper, owner, k, v) {
			result = false
			break
		}
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return result
}

// UpdateConfiguration 更新配置项
func UpdateConfiguration(owner, key, value string) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SetOption(helper, owner, key, value)
}

// GetConfigurations 获取配置集
func GetConfigurations(owner string, keys []string) map[string]string {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ret := map[string]string{}
	for _, k := range keys {
		v, found := dal.GetOption(helper, owner, k)
		if found {
			ret[k] = v
		}
	}

	return ret
}

// GetConfiguration 获取指定配置项
func GetConfiguration(owner, key string) (string, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.GetOption(helper, owner, key)
}
