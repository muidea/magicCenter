package route

import (
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/extern/modules/blog/def"
	"muidea.com/magicCommon/foundation/net"
	common_const "muidea.com/magicCommon/common"
)

// AppendBlogRoute 追加User Route
func AppendBlogRoute(routes []common.Route, contentHandler common.ContentHandler) []common.Route {

	rt := CreateMainPageRoute(contentHandler)
	routes = append(routes, rt)

	rt = CreateDetailPageRoute(contentHandler)
	routes = append(routes, rt)

	rt = CreateListPageRoute(contentHandler)
	routes = append(routes, rt)

	return routes
}

// CreateMainPageRoute 新建GetUserRoute
func CreateMainPageRoute(contentHandler common.ContentHandler) common.Route {
	i := mainPageRoute{contentHandler: contentHandler}
	return &i
}

// CreateDetailPageRoute 新建GetAllUser Route
func CreateDetailPageRoute(contentHandler common.ContentHandler) common.Route {
	i := detailPageRoute{contentHandler: contentHandler}
	return &i
}

// CreateListPageRoute 新建CreateUser Route
func CreateListPageRoute(contentHandler common.ContentHandler) common.Route {
	i := listPageRoute{contentHandler: contentHandler}
	return &i
}

type mainPageRoute struct {
	contentHandler common.ContentHandler
}

func (i *mainPageRoute) Method() string {
	return common.GET
}

func (i *mainPageRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetMainPage)
}

func (i *mainPageRoute) Handler() interface{} {
	return i.getMainPageHandler
}

func (i *mainPageRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *mainPageRoute) getMainPageHandler(res http.ResponseWriter, req *http.Request) {
	log.Print("getMainPageHandler")
}

type detailPageRoute struct {
	contentHandler common.ContentHandler
}

func (i *detailPageRoute) Method() string {
	return common.GET
}

func (i *detailPageRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetDetailPage)
}

func (i *detailPageRoute) Handler() interface{} {
	return i.getDetailPageHandler
}

func (i *detailPageRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *detailPageRoute) getDetailPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getDetailPageHandler")
}

type listPageRoute struct {
	contentHandler common.ContentHandler
}

func (i *listPageRoute) Method() string {
	return common.GET
}

func (i *listPageRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetListPage)
}

func (i *listPageRoute) Handler() interface{} {
	return i.getListPageHandler
}

func (i *listPageRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *listPageRoute) getListPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getListPageHandler")

}
