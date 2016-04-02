package bll

import (
	"magiccenter/util/modelhelper"
	"magiccenter/configuration/dal"
)

func UpdateConfigurations(configs map[string]string) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	result := true
	helper.BeginTransaction()
	for k,v := range configs {
		if !dal.SetOption(helper, k, v) {
			result = false
			break;
		} 
	}
	
	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}
	
	return result	
}

func UpdateConfiguration(key, value string) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return dal.SetOption(helper, key, value)
}

func GetConfiguration(keys []string) map[string]string {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	ret := map[string]string{}
	for _, k := range keys {
		v,found := dal.GetOption(helper,k)
		if found {
			ret[k] = v
		} else {
			ret[k] = ""
		}
	}
	
	return ret
}
