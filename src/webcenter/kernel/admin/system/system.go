package system

import (
    "webcenter/util/modelhelper"
    "webcenter/kernel"
    "webcenter/kernel/bll"
)

func init() {
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	info := bll.GetSystemInfo()
	kernel.UpdateSystemInfo(info)
}