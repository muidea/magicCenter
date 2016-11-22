package static

import (
	"magiccenter/auth"
	"magiccenter/common"
	"magiccenter/module"
	"magiccenter/router"

	"muidea.com/util"
)

// ID Static模块ID
const ID = "f17133ec-63f0-4b46-0000-e6ca1af6fe4e"

// Name Static模块名称
const Name = "Magic Static"

// Description Static模块描述信息
const Description = "Magic 静态文件管理"

// URL Static模块Url
const URL string = "static"

type static struct {
}

var instance *static

// LoadModule 加载Static模块
func LoadModule() {
	if instance == nil {
		instance = &static{}
	}

	module.RegisterModule(instance)
}

// ID Static ID
func (instance *static) ID() string {
	return ID
}

// Name Static 名称
func (instance *static) Name() string {
	return Name
}

// Description Static名称
func (instance *static) Description() string {
	return Description
}

func (instance *static) Group() string {
	return "resource"
}

func (instance *static) Type() int {
	return common.INTERNAL
}

// URL Static url
func (instance *static) URL() string {
	return URL
}

func (instance *static) Status() int {
	return 0
}

func (instance *static) EndPoint() common.EndPoint {
	return nil
}

// Route Static 路由信息
func (instance *static) Routes() []common.Route {
	routes := []common.Route{
		router.NewRoute(common.GET, "/static/**", viewArticleHandler, nil),
		router.NewRoute(common.GET, "/maintain/", MaintainViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.POST, "/ajaxMaintain/", MaintainActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

// Startup 启动Static模块
func (instance *static) Startup() bool {
	return true
}

// Cleanup 清除Static模块
func (instance *static) Cleanup() {

}

// Invoke 执行外部命令
func (instance *static) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
