package route

import (
	"log"
	"net/http"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/static/def"
	common_const "github.com/muidea/magicCommon/common"
	"github.com/muidea/magicCommon/foundation/net"
)

// CreateStaticResRoute 新建静态资源路由
func CreateStaticResRoute(staticHandler common.StaticHandler) common.Route {
	i := &staticResRoute{staticHandler: staticHandler}

	return i
}

type staticResRoute struct {
	staticHandler common.StaticHandler
}

func (i *staticResRoute) Method() string {
	return common.GET
}

func (i *staticResRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetStatic)
}

func (i *staticResRoute) Handler() interface{} {
	return i.getStaticResHandler
}

func (i *staticResRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *staticResRoute) getStaticResHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("getStaticResHandler, path:%s", r.URL.Path)

	i.staticHandler.HandleResource(w, r)
}
