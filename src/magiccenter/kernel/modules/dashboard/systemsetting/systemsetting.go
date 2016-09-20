package systemsetting

import (
	"magiccenter/common"
	"magiccenter/kernel/auth"
	"magiccenter/kernel/modules/dashboard/systemsetting/ui"
	"magiccenter/module"
	"magiccenter/router"

	"muidea.com/util"
)

// ID 模块ID
const ID = "8ec33ffc-0d7c-4d1d-84e8-f1983197fc9f"

// Name 模块名称
const Name = "Magic SystemSetting"

// Description 描述信息
const Description = "Magic 系统设置模块"

// URL 模块Url
const URL string = "admin"

type systemsetting struct {
}

var instance *systemsetting

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &systemsetting{}
	}

	module.RegisterModule(instance)
}

func (instance *systemsetting) ID() string {
	return ID
}

func (instance *systemsetting) Name() string {
	return Name
}

func (instance *systemsetting) Description() string {
	return Description
}

func (instance *systemsetting) Group() string {
	return "admin systemsetting"
}

func (instance *systemsetting) Type() int {
	return common.KERNEL
}

func (instance *systemsetting) URL() string {
	return URL
}

func (instance *systemsetting) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *systemsetting) Routes() []common.Route {
	routes := []common.Route{
		// 管理视图
		router.NewRoute(common.GET, "systemsetting/", ui.SystemSettingViewHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *systemsetting) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *systemsetting) Cleanup() {

}

// Invoke 执行外部命令
func (instance *systemsetting) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
