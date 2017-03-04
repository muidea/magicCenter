package authority

import (
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/kernel/modules/authority/ui"
	"magiccenter/system"

	"muidea.com/util"
)

// ID 模块ID
const ID = "759a2ee4-147a-4169-ba89-15c0c692bc16"

// Name 模块名称
const Name = "Magic Authority"

// Description 模块描述信息
const Description = "Magic 权限管理模块"

// URL 模块Url
const URL string = "/authority"

type authority struct {
}

var instance *authority

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &authority{}
	}

	modulehub := system.GetModuleHub()
	modulehub.RegisterModule(instance)
}

func (instance *authority) ID() string {
	return ID
}

func (instance *authority) Name() string {
	return Name
}

func (instance *authority) Description() string {
	return Description
}

func (instance *authority) Group() string {
	return "kernel"
}

func (instance *authority) Type() int {
	return common.KERNEL
}

func (instance *authority) URL() string {
	return URL
}

func (instance *authority) Status() int {
	return 0
}

func (instance *authority) EndPoint() common.EndPoint {
	return nil
}

func (instance *authority) AuthGroups() []common.AuthGroup {
	groups := []common.AuthGroup{}

	return groups
}

// Route 路由信息
func (instance *authority) Routes() []common.Route {
	routes := []common.Route{}

	return routes
}

// Startup 启动模块
func (instance *authority) Startup() bool {
	return true
}

// Cleanup 清除模块
func (instance *authority) Cleanup() {

}

// Invoke 执行外部命令
func (instance *authority) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}
	switch param.(type) {
	case *commonbll.QueryAuthGroupRequest:
		{
			request := param.(*commonbll.QueryAuthGroupRequest)
			response := result.(*commonbll.QueryAuthGroupResponse)
			if request != nil && response != nil {
				return ui.QueryAuthGroup(request, response)
			}
		}
	case *commonbll.InsertAuthGroupRequest:
		{
			request := param.(*commonbll.InsertAuthGroupRequest)
			response := result.(*commonbll.InsertAuthGroupResponse)
			if request != nil && response != nil {
				return ui.InsertAuthGroup(request, response)
			}
		}
	case *commonbll.DeleteAuthGroupRequest:
		{
			request := param.(*commonbll.DeleteAuthGroupRequest)
			response := result.(*commonbll.DeleteAuthGroupResponse)
			if request != nil && response != nil {
				return ui.DeleteAuthGroup(request, response)
			}
		}
	}
	return false
}
