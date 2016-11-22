package blog

import (
	"magiccenter/auth"
	"magiccenter/common"
	"magiccenter/module"
	"magiccenter/router"

	"muidea.com/util"
)

// ID Blog Module ID
const ID = "f17133ec-63e9-4b46-8757-e6ca1af6fe3e"

// Name Blog Module Name
const Name = "Magic Blog"

// Description Blog Module Description
const Description = "Magic 博客"

// URL Blog Module URL
const URL = "blog"

type blog struct {
}

var instance *blog

// LoadModule 加载模块
func LoadModule() {
	if instance == nil {
		instance = &blog{}
	}

	module.RegisterModule(instance)
}

func (b *blog) ID() string {
	return ID
}

func (b *blog) Name() string {
	return Name
}

func (b *blog) Description() string {
	return Description
}

func (b *blog) Group() string {
	return "content"
}

func (b *blog) Type() int {
	return common.INTERNAL
}

func (b *blog) URL() string {
	return URL
}

func (b *blog) Status() int {
	return 0
}

func (b *blog) EndPoint() common.EndPoint {
	return nil
}

func (b *blog) Routes() []common.Route {
	routes := []common.Route{
		router.NewRoute(common.GET, "/", indexHandler, nil),
		router.NewRoute(common.GET, "/view/", viewContentHandler, nil),
		router.NewRoute(common.GET, "/catalog/", viewCatalogHandler, nil),
		router.NewRoute(common.GET, "/link/", viewLinkHandler, nil),
		router.NewRoute(common.GET, "/maintain/", MaintainViewHandler, auth.AdminAuthVerify()),
		router.NewRoute(common.POST, "/ajaxMaintain/", MaintainActionHandler, auth.AdminAuthVerify()),
	}

	return routes
}

func (b *blog) Startup() bool {
	return true
}

func (b *blog) Cleanup() {

}

func (b *blog) Invoke(param interface{}, result interface{}) bool {
	util.ValidataPtr(param)
	if result != nil {
		util.ValidataPtr(result)
	}

	return false
}
