package static

import (
	"magiccenter/module"
)

// ID Static模块ID
const ID = "f17133ec-63f0-4b46-0000-e6ca1af6fe4e"

// Name Static模块名称
const Name = "Magic Static"

// Description Static模块描述信息
const Description = "Magic 静态管理模块"

// URI Static模块URI
const URI = "/static"

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

// Startup 启动Static模块
func (instance *static) Startup() {
}

// Cleanup 清除Static模块
func (instance *static) Cleanup() {

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

// Uri Static Uri
func (instance *static) Uri() string {
	return URI
}

// Route Static 路由信息
func (instance *static) Routes() []module.Route {
	routes := []module.Route{module.NewRoute(module.GET, "/static/**", viewArticleHandler)}

	return routes
}
