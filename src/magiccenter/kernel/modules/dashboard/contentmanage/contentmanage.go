package contentmanage

import (
	"magiccenter/auth"
	"magiccenter/common"
	"magiccenter/kernel/modules/dashboard/contentmanage/ui"
	"magiccenter/module"
	"magiccenter/router"

	"muidea.com/util"
)

// ID 模块ID
const ID = "ffe1c37f-4fea-4a03-a7c3-55a331f5995f"

// Name 模块名称
const Name = "Magic ContentManage"

// Description 模块描述信息
const Description = "Magic 内容管理模块"

// URL 模块Url
const URL string = "admin"

type contentmanage struct {
}

var instance *contentmanage

// LoadModule 加载ModuleManage模块
func LoadModule() {
	if instance == nil {
		instance = &contentmanage{}
	}

	module.RegisterModule(instance)
}

func (instance *contentmanage) ID() string {
	return ID
}

func (instance *contentmanage) Name() string {
	return Name
}

func (instance *contentmanage) Description() string {
	return Description
}

func (instance *contentmanage) Group() string {
	return "admin contentmanage"
}

func (instance *contentmanage) Type() int {
	return common.KERNEL
}

func (instance *contentmanage) URL() string {
	return URL
}

func (instance *contentmanage) EndPoint() common.EndPoint {
	return nil
}

// Route 路由信息
func (instance *contentmanage) Routes() []common.Route {
	routes := []common.Route{
		// 模块设置视图
		router.NewRoute(common.GET, "contentview/", ui.ContentManageViewHandler, auth.AdminAuthVerify()),
		// 获取指定Block对应的Items
		router.NewRoute(common.GET, "blockitem/", ui.BlockItemActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动模块
func (instance *contentmanage) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *contentmanage) Cleanup() {

}

// Invoke 执行外部命令
func (instance *contentmanage) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
