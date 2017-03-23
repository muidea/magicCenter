package route

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/net"
	"muidea.com/magicCenter/foundation/util"
)

// AppendLinkRoute 追加User Route
func AppendLinkRoute(routes []common.Route, modHub common.ModuleHub, sessionRegistry common.SessionRegistry) []common.Route {

	rt, _ := CreateGetLinkByIDRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateGetLinkListRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateCreateLinkRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateUpdateLinkRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateDestroyLinkRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetLinkByIDRoute 新建GetLink Route
func CreateGetLinkByIDRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkGetListRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetLinkListRoute 新建GetAllLink Route
func CreateGetLinkListRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkGetListRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateCreateLinkRoute 新建CreateLinkRoute Route
func CreateCreateLinkRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkCreateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateUpdateLinkRoute UpdateLinkRoute Route
func CreateUpdateLinkRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkUpdateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateDestroyLinkRoute DestroyLinkRoute Route
func CreateDestroyLinkRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkDestroyRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

type linkGetByIDRoute struct {
	contentHandler common.ContentHandler
}

type linkGetByIDResult struct {
	common.Result
	Link model.LinkDetail
}

func (i *linkGetByIDRoute) Action() string {
	return common.GET
}

func (i *linkGetByIDRoute) Pattern() string {
	return "content/link/[0-9]+/"
}

func (i *linkGetByIDRoute) Handler() interface{} {
	return i.getLinkHandler
}

func (i *linkGetByIDRoute) getLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getLinkHandler")

	result := linkGetByIDResult{}
	_, value := net.SplitResetAPI(r.URL.Path)
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

		result.ErrCode = 1
		result.Reason = "无效参数"
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

func (i *linkGetListRoute) Action() string {
	return common.GET
}

func (i *linkGetListRoute) Pattern() string {
	return "content/link/"
}

func (i *linkGetListRoute) Handler() interface{} {
	return i.getLinkListHandler
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

func (i *linkCreateRoute) Action() string {
	return common.POST
}

func (i *linkCreateRoute) Pattern() string {
	return "content/link/"
}

func (i *linkCreateRoute) Handler() interface{} {
	return i.createLinkHandler
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

		name := r.FormValue("link-name")
		url := r.FormValue("link-url")
		logo := r.FormValue("link-logo")
		catalogs, _ := util.Str2IntArray(r.FormValue("link-catalog"))
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

func (i *linkUpdateRoute) Action() string {
	return common.PUT
}

func (i *linkUpdateRoute) Pattern() string {
	return "content/link/[0-9]+/"
}

func (i *linkUpdateRoute) Handler() interface{} {
	return i.updateLinkHandler
}

func (i *linkUpdateRoute) updateLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	_, value := net.SplitResetAPI(r.URL.Path)
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
		link.Name = r.FormValue("link-name")
		link.URL = r.FormValue("link-url")
		link.Logo = r.FormValue("link-logo")
		link.Catalog, _ = util.Str2IntArray(r.FormValue("link-catalog"))
		link.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		link.Author = user.ID
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

func (i *linkDestroyRoute) Action() string {
	return common.DELETE
}

func (i *linkDestroyRoute) Pattern() string {
	return "content/link/[0-9]+/"
}

func (i *linkDestroyRoute) Handler() interface{} {
	return i.deleteLinkHandler
}

func (i *linkDestroyRoute) deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	_, value := net.SplitResetAPI(r.URL.Path)
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
