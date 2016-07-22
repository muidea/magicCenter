package blog

import (
	"magiccenter/module"
	"magiccenter/router"
)

// ID Blog Module ID
const ID = "f17133ec-63e9-4b46-8757-e6ca1af6fe3e"

// Name Blog Module Name
const Name = "Magic Blog"

// Description Blog Module Description
const Description = "Magic 博客系统"

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
	return module.INTERNAL
}

func (b *blog) URL() string {
	return URL
}

func (b *blog) Resource() module.Resource {
	return nil
}

func (b *blog) Routes() []router.Route {
	routes := []router.Route{
		router.NewRoute(router.GET, "/", indexHandler, nil),
		router.NewRoute(router.GET, "/view/", viewContentHandler, nil),
		router.NewRoute(router.GET, "/catalog/", viewCatalogHandler, nil),
		router.NewRoute(router.GET, "/link/", viewLinkHandler, nil),
	}

	return routes
}

func (b *blog) Startup() bool {
	return true
}

func (b *blog) Cleanup() {

}

func (b *blog) Invoke(param interface{}) bool {
	return false
}
