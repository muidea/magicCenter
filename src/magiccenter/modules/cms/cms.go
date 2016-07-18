package cms

import (
	"magiccenter/module"
)

const ID = "f17133ec-63e9-4b46-8758-e6ca1af6fe3f"
const Name = "Magic CMS"
const Description = "Magic 内容管理系统"
const URI = "/cms"

type cms struct {
}

var instance *cms = nil

func LoadModule() {
	if instance == nil {
		instance = &cms{}
	}

	module.RegisterModule(instance)
}

func (this *cms) ID() string {
	return ID
}

func (this *cms) Name() string {
	return Name
}

func (this *cms) Description() string {
	return Description
}

func (this *cms) Uri() string {
	return URI
}

func (this *cms) Routes() []module.Route {
	routes := []module.Route{module.NewRoute(module.GET, "/", indexHandler), module.NewRoute(module.GET, "/view/", viewContentHandler), module.NewRoute(module.GET, "/catalog/", viewCatalogHandler), module.NewRoute(module.GET, "/link/", viewLinkHandler)}

	return routes
}

func (this *cms) Startup() {
}

func (this *cms) Cleanup() {

}
