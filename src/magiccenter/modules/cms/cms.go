package cms

import (
	"magiccenter/common"
	"magiccenter/module"
	"magiccenter/router"
)

// ID CMS Module ID
const ID = "f17133ec-63e9-4b46-8758-e6ca1af6fe3f"

// Name CMS Module Name
const Name = "Magic CMS"

// Description CMS Module Description
const Description = "Magic 内容管理系统"

// URL CMS Module URL
const URL = "cms"

type cms struct {
}

var instance *cms

// LoadModule 加载CMS模块
func LoadModule() {
	if instance == nil {
		instance = &cms{}
	}

	module.RegisterModule(instance)
}

func (c *cms) ID() string {
	return ID
}

func (c *cms) Name() string {
	return Name
}

func (c *cms) Description() string {
	return Description
}

func (c *cms) Group() string {
	return "content"
}

func (c *cms) Type() int {
	return common.INTERNAL
}

func (c *cms) URL() string {
	return URL
}

func (c *cms) EndPoint() common.EndPoint {
	return nil
}

func (c *cms) Routes() []router.Route {
	routes := []router.Route{
		router.NewRoute(router.GET, "/", indexHandler, nil),
		router.NewRoute(router.GET, "/view/", viewContentHandler, nil),
		router.NewRoute(router.GET, "/catalog/", viewCatalogHandler, nil),
		router.NewRoute(router.GET, "/link/", viewLinkHandler, nil),
	}

	return routes
}

func (c *cms) Startup() bool {
	return true
}

func (c *cms) Cleanup() {

}

func (c *cms) Invoke(param interface{}) bool {
	return false
}
