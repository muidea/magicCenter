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

// CreateGetLinkRoute 新建GetLink Route
func CreateGetLinkRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkGetRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetAllLinkRoute 新建GetAllLink Route
func CreateGetAllLinkRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkGetAllRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetByCatalogLinkRoute 新建GetByCatalogLinkRoute Route
func CreateGetByCatalogLinkRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := linkGetByCatalogRoute{contentHandler: endPoint.(common.ContentHandler)}
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

type linkGetRoute struct {
	contentHandler common.ContentHandler
}

type linkGetResult struct {
	common.Result
	Link model.Link
}

func (i *linkGetRoute) Type() string {
	return common.GET
}

func (i *linkGetRoute) Pattern() string {
	return "content/link/[0-9]*/"
}

func (i *linkGetRoute) Handler() interface{} {
	return i.getLinkHandler
}

func (i *linkGetRoute) getLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getLinkHandler")

	result := linkGetResult{}
	value, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if ok {
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

type linkGetAllRoute struct {
	contentHandler common.ContentHandler
}

type linkGetAllResult struct {
	common.Result
	Link []model.Link
}

func (i *linkGetAllRoute) Type() string {
	return common.GET
}

func (i *linkGetAllRoute) Pattern() string {
	return "content/link/"
}

func (i *linkGetAllRoute) Handler() interface{} {
	return i.getAllLinkHandler
}

func (i *linkGetAllRoute) getAllLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllLinkHandler")

	result := linkGetAllResult{}
	for true {
		result.Link = i.contentHandler.GetAllLink()
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type linkGetByCatalogRoute struct {
	contentHandler common.ContentHandler
}

type linkGetByCatalogResult struct {
	common.Result
	Link []model.Link
}

func (i *linkGetByCatalogRoute) Type() string {
	return common.GET
}

func (i *linkGetByCatalogRoute) Pattern() string {
	return "content/link/?catalog=[0-9]*"
}

func (i *linkGetByCatalogRoute) Handler() interface{} {
	return i.getByCatalogLinkHandler
}

func (i *linkGetByCatalogRoute) getByCatalogLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getByCatalogLinkHandler")

	result := linkGetByCatalogResult{}
	_, params, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if ok {
			catalog, ok := params["catalog"]
			if !ok {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			id, err := strconv.Atoi(catalog)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			result.Link = i.contentHandler.GetLinkByCatalog(id)
			result.ErrCode = 0
			break
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

type linkCreateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type linkCreateResult struct {
	common.Result
	Link model.Link
}

func (i *linkCreateRoute) Type() string {
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

		title := r.FormValue("link-title")
		content := r.FormValue("link-content")
		catalogs, _ := util.Str2IntArray(r.FormValue("link-catalog"))
		createDate := time.Now().Format("2006-01-02 15:04:05")
		link, ok := i.contentHandler.CreateLink(title, content, createDate, catalogs, user.ID)
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
	Link model.LinkSummary
}

func (i *linkUpdateRoute) Type() string {
	return common.PUT
}

func (i *linkUpdateRoute) Pattern() string {
	return "content/link/[0-9]*/"
}

func (i *linkUpdateRoute) Handler() interface{} {
	return i.updateLinkHandler
}

func (i *linkUpdateRoute) updateLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	value, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if !ok {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}
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
		link := model.Link{}
		link.ID = id
		link.Title = r.FormValue("link-title")
		link.Content = r.FormValue("link-content")
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

func (i *linkDestroyRoute) Type() string {
	return common.DELETE
}

func (i *linkDestroyRoute) Pattern() string {
	return "content/link/[0-9]*/"
}

func (i *linkDestroyRoute) Handler() interface{} {
	return i.deleteLinkHandler
}

func (i *linkDestroyRoute) deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteLinkHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := linkCreateResult{}
	value, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if !ok {
			result.ErrCode = 1
			result.Reason = "无效参数"
		}
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
