package route

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/content/def"
	"muidea.com/magicCenter/foundation/net"
	"muidea.com/magicCenter/foundation/util"
)

// AppendLinkRoute 追加User Route
func AppendLinkRoute(routes []common.Route, contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateGetLinkByIDRoute(contentHandler)
	routes = append(routes, rt)

	rt = CreateGetLinkListRoute(contentHandler)
	routes = append(routes, rt)

	rt = CreateCreateLinkRoute(contentHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateLinkRoute(contentHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyLinkRoute(contentHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetLinkByIDRoute 新建GetLink Route
func CreateGetLinkByIDRoute(contentHandler common.ContentHandler) common.Route {
	i := linkGetByIDRoute{contentHandler: contentHandler}
	return &i
}

// CreateGetLinkListRoute 新建GetAllLink Route
func CreateGetLinkListRoute(contentHandler common.ContentHandler) common.Route {
	i := linkGetListRoute{contentHandler: contentHandler}
	return &i
}

// CreateCreateLinkRoute 新建CreateLinkRoute Route
func CreateCreateLinkRoute(contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := linkCreateRoute{contentHandler: contentHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateLinkRoute UpdateLinkRoute Route
func CreateUpdateLinkRoute(contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := linkUpdateRoute{contentHandler: contentHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyLinkRoute DestroyLinkRoute Route
func CreateDestroyLinkRoute(contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := linkDestroyRoute{contentHandler: contentHandler, sessionRegistry: sessionRegistry}
	return &i
}

type linkGetByIDRoute struct {
	contentHandler common.ContentHandler
}

type linkGetByIDResult struct {
	common.Result
	Link model.LinkDetail
}

func (i *linkGetByIDRoute) Method() string {
	return common.GET
}

func (i *linkGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetLinkDetail)
}

func (i *linkGetByIDRoute) Handler() interface{} {
	return i.getLinkHandler
}

func (i *linkGetByIDRoute) AuthGroup() int {
	return common.VisitorAuthGroup.ID
}

func (i *linkGetByIDRoute) getLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getLinkHandler")

	result := linkGetByIDResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		link, ok := i.contentHandler.GetLinkByID(id)
		if ok {
			result.Link = link
			result.ErrCode = 0
		} else {
			result.ErrCode = 1
			result.Reason = "对象不存在"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type linkGetListRoute struct {
	contentHandler common.ContentHandler
}

type linkGetListResult struct {
	common.Result
	Link []model.Summary
}

func (i *linkGetListRoute) Method() string {
	return common.GET
}

func (i *linkGetListRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetLinkList)
}

func (i *linkGetListRoute) Handler() interface{} {
	return i.getLinkListHandler
}

func (i *linkGetListRoute) AuthGroup() int {
	return common.VisitorAuthGroup.ID
}

func (i *linkGetListRoute) getLinkListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getLinkListHandler")

	result := linkGetListResult{}
	for true {
		catalog := r.URL.Query()["catalog"]
		if len(catalog) < 1 {
			result.Link = i.contentHandler.GetAllLink()
			result.ErrCode = 0
			break
		}

		id, err := strconv.Atoi(catalog[0])
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		result.Link = i.contentHandler.GetLinkByCatalog(id)
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type linkCreateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type linkCreateResult struct {
	common.Result
	Link model.Summary
}

func (i *linkCreateRoute) Method() string {
	return common.POST
}

func (i *linkCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostLink)
}

func (i *linkCreateRoute) Handler() interface{} {
	return i.createLinkHandler
}

func (i *linkCreateRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *linkCreateRoute) createLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrCode = 1
			result.Reason = "无效权限"
			break
		}

		r.ParseForm()

		name := r.FormValue("name")
		url := r.FormValue("url")
		logo := r.FormValue("logo")
		catalogs, _ := util.Str2IntArray(r.FormValue("catalog"))
		createDate := time.Now().Format("2006-01-02 15:04:05")
		link, ok := i.contentHandler.CreateLink(name, url, logo, createDate, catalogs, user.ID)
		if !ok {
			result.ErrCode = 1
			result.Reason = "新建失败"
			break
		}
		result.ErrCode = 0
		result.Link = link
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type linkUpdateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type linkUpdateResult struct {
	common.Result
	Link model.Summary
}

func (i *linkUpdateRoute) Method() string {
	return common.PUT
}

func (i *linkUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutLink)
}

func (i *linkUpdateRoute) Handler() interface{} {
	return i.updateLinkHandler
}

func (i *linkUpdateRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *linkUpdateRoute) updateLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		user, found := session.GetAccount()
		if !found {
			result.ErrCode = 1
			result.Reason = "无效权限"
			break
		}

		r.ParseForm()
		link := model.LinkDetail{}
		link.ID = id
		link.Name = r.FormValue("name")
		link.URL = r.FormValue("url")
		link.Logo = r.FormValue("logo")
		link.Catalog, _ = util.Str2IntArray(r.FormValue("catalog"))
		link.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		link.Creater = user.ID
		summmary, ok := i.contentHandler.SaveLink(link)
		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}
		result.ErrCode = 0
		result.Link = summmary
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type linkDestroyRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type linkDestroyResult struct {
	common.Result
}

func (i *linkDestroyRoute) Method() string {
	return common.DELETE
}

func (i *linkDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteLink)
}

func (i *linkDestroyRoute) Handler() interface{} {
	return i.deleteLinkHandler
}

func (i *linkDestroyRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *linkDestroyRoute) deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}
		_, found := session.GetAccount()
		if !found {
			result.ErrCode = 1
			result.Reason = "无效权限"
			break
		}

		ok := i.contentHandler.DestroyLink(id)
		if !ok {
			result.ErrCode = 1
			result.Reason = "删除失败"
			break
		}
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
