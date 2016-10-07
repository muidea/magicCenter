package bll

import (
	"magiccenter/configuration/dal"
	"magiccenter/util/dbhelper"
)

// UpdateConfigurations 更新配置集
func UpdateConfigurations(configs map[string]string) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	result := true
	helper.BeginTransaction()
	for k, v := range configs {
		if !dal.SetOption(helper, k, v) {
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
func UpdateConfiguration(key, value string) bool {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SetOption(helper, key, value)
}

// GetConfigurations 获取配置集
func GetConfigurations(keys []string) map[string]string {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ret := map[string]string{}
	for _, k := range keys {
		v, found := dal.GetOption(helper, k)
		if found {
			ret[k] = v
		} else {
			ret[k] = ""
		}
	}

	return ret
}

// GetConfiguration 获取指定配置项
func GetConfiguration(key string) (string, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.GetOption(helper, key)
}
