package route

import (
	"net/http"

	"log"

	"muidea.com/magicCenter/application/common"
)

// CreateStaticViewRoute 新建静态视图路由
func CreateStaticViewRoute(staticHandler common.StaticHandler) common.Route {
	i := &staticViewRoute{staticHandler: staticHandler}

	return i
}

type staticViewRoute struct {
	staticHandler common.StaticHandler
}

func (i *staticViewRoute) Type() string {
	return common.GET
}

func (i *staticViewRoute) Pattern() string {
	return "**"
}

func (i *staticViewRoute) Handler() interface{} {
	return i.getStaticViewHandler
}

func (i *staticViewRoute) getStaticViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("getStaticViewHandler, path:%s", r.URL.Path)

	i.staticHandler.HandleView("", w, r)
}
