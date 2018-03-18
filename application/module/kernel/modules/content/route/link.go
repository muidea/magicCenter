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
)

// AppendLinkRoute 追加User Route
func AppendLinkRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateGetLinkByIDRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetLinkListRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateLinkRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateLinkRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyLinkRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetLinkByIDRoute 新建GetLink Route
func CreateGetLinkByIDRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := linkGetByIDRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetLinkListRoute 新建GetAllLink Route
func CreateGetLinkListRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := linkGetListRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateCreateLinkRoute 新建CreateLinkRoute Route
func CreateCreateLinkRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := linkCreateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateLinkRoute UpdateLinkRoute Route
func CreateUpdateLinkRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := linkUpdateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyLinkRoute DestroyLinkRoute Route
func CreateDestroyLinkRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := linkDestroyRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

type linkGetByIDRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

type linkGetByIDResult struct {
	common.Result
	Link model.LinkDetailView `json:"link"`
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
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		link, ok := i.contentHandler.GetLinkByID(id)
		if ok {
			user, _ := i.accountHandler.FindUserByID(link.Creater)
			catalogs := i.contentHandler.GetCatalogs(link.Catalog)

			result.Link.LinkDetail = link
			result.Link.Creater = user.User
			result.Link.Catalog = catalogs
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
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
	accountHandler common.AccountHandler
}

type linkGetListResult struct {
	common.Result
	Link []model.SummaryView `json:"link"`
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
		catalog := r.URL.Query().Get("catalog")
		if catalog == "" {
			links := i.contentHandler.GetAllLink()
			for _, val := range links {
				link := model.SummaryView{}
				user, _ := i.accountHandler.FindUserByID(val.Creater)
				catalogs := i.contentHandler.GetCatalogs(val.Catalog)

				link.Summary = val
				link.Creater = user.User
				link.Catalog = catalogs

				result.Link = append(result.Link, link)
			}
			result.ErrorCode = common.Success
			break
		}

		id, err := strconv.Atoi(catalog)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		links := i.contentHandler.GetLinkByCatalog(id)
		for _, val := range links {
			link := model.SummaryView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogs := i.contentHandler.GetCatalogs(val.Catalog)

			link.Summary = val
			link.Creater = user.User
			link.Catalog = catalogs

			result.Link = append(result.Link, link)
		}
		result.ErrorCode = common.Success
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
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

type linkCreateParam struct {
	Name    string          `json:"name"`
	URL     string          `json:"url"`
	Logo    string          `json:"logo"`
	Catalog []model.Catalog `json:"catalog"`
}

type linkCreateResult struct {
	common.Result
	Link model.SummaryView `json:"link"`
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
			result.ErrorCode = common.Failed
			result.Reason = "无效权限"
			break
		}

		param := &linkCreateParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		createDate := time.Now().Format("2006-01-02 15:04:05")
		catalogIds := []int{}
		catalogs, ok := i.contentHandler.UpdateCatalog(param.Catalog, createDate, user.ID)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "更新Catalog失败"
			break
		}
		for _, val := range catalogs {
			catalogIds = append(catalogIds, val.ID)
		}

		link, ok := i.contentHandler.CreateLink(param.Name, param.URL, param.Logo, createDate, catalogIds, user.ID)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "新建失败"
			break
		}

		result.Link.Summary = link
		result.Link.Creater = user
		result.Link.Catalog = catalogs
		result.ErrorCode = common.Success
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
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

type linkUpdateParam linkCreateParam

type linkUpdateResult struct {
	common.Result
	Link model.SummaryView `json:"link"`
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
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common.Failed
			result.Reason = "无效权限"
			break
		}

		param := &linkUpdateParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		updateDate := time.Now().Format("2006-01-02 15:04:05")
		catalogIds := []int{}
		catalogs, ok := i.contentHandler.UpdateCatalog(param.Catalog, updateDate, user.ID)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "更新Catalog失败"
			break
		}
		for _, val := range catalogs {
			catalogIds = append(catalogIds, val.ID)
		}

		link := model.LinkDetail{}
		link.ID = id
		link.Name = param.Name
		link.URL = param.URL
		link.Logo = param.Logo
		link.Catalog = catalogIds
		link.CreateDate = updateDate
		link.Creater = user.ID
		summmary, ok := i.contentHandler.SaveLink(link)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "更新失败"
			break
		}

		result.Link.Summary = summmary
		result.Link.Creater = user
		result.Link.Catalog = catalogs
		result.ErrorCode = common.Success
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
	accountHandler  common.AccountHandler
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
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}
		_, found := session.GetAccount()
		if !found {
			result.ErrorCode = common.Failed
			result.Reason = "无效权限"
			break
		}

		ok := i.contentHandler.DestroyLink(id)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "删除失败"
			break
		}
		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
