package route

import (
	"encoding/json"
	"net/http"
	"time"

	"strconv"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/content/def"
	common_const "github.com/muidea/magicCommon/common"
	common_def "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/net"
	"github.com/muidea/magicCommon/model"
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
	return common_const.VisitorAuthGroup.ID
}

func (i *linkGetByIDRoute) getLinkHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.QueryLinkResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		link, ok := i.contentHandler.GetLinkByID(id)
		if ok {
			user, _ := i.accountHandler.FindUserByID(link.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(link.Catalog)

			result.Link.LinkDetail = link
			result.Link.Creater = user.User
			result.Link.Catalog = catalogSummarys
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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
	return common_const.VisitorAuthGroup.ID
}

func (i *linkGetListRoute) getLinkListHandler(w http.ResponseWriter, r *http.Request) {
	result := common_def.QueryLinkListResult{Link: []model.SummaryView{}}
	for true {
		filter := &common_def.Filter{}
		filter.Decode(r)

		catalog, err := common_def.DecodeStrictCatalog(r)
		if catalog == nil && err == nil {
			links, total := i.contentHandler.GetAllLink(filter)
			for _, val := range links {
				link := model.SummaryView{}
				user, _ := i.accountHandler.FindUserByID(val.Creater)
				catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

				link.Summary = val
				link.Creater = user.User
				link.Catalog = catalogSummarys

				result.Link = append(result.Link, link)
			}
			result.Total = total
			result.ErrorCode = common_def.Success
			break
		} else if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			break
		}

		links, total := i.contentHandler.GetLinkByCatalog(*catalog, filter)
		for _, val := range links {
			link := model.SummaryView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

			link.Summary = val
			link.Creater = user.User
			link.Catalog = catalogSummarys

			result.Link = append(result.Link, link)
		}
		result.Total = total
		result.ErrorCode = common_def.Success
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
	return common_const.UserAuthGroup.ID
}

func (i *linkCreateRoute) createLinkHandler(w http.ResponseWriter, r *http.Request) {
	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.CreateLinkResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.CreateLinkParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		createDate := time.Now().Format("2006-01-02 15:04:05")

		link, ok := i.contentHandler.CreateLink(param.Name, param.Description, param.URL, param.Logo, createDate, param.Catalog, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "新建失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)

		result.Link.Summary = link
		result.Link.Creater = user
		result.Link.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
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
	return common_const.UserAuthGroup.ID
}

func (i *linkUpdateRoute) updateLinkHandler(w http.ResponseWriter, r *http.Request) {
	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.UpdateLinkResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.UpdateLinkParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		updateDate := time.Now().Format("2006-01-02 15:04:05")
		link := model.LinkDetail{}
		link.ID = id
		link.Name = param.Name
		link.Description = param.Description
		link.URL = param.URL
		link.Logo = param.Logo
		link.Catalog = param.Catalog
		link.CreateDate = updateDate
		link.Creater = user.ID
		summmary, ok := i.contentHandler.SaveLink(link)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)

		result.Link.Summary = summmary
		result.Link.Creater = user
		result.Link.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
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
	return common_const.MaintainerAuthGroup.ID
}

func (i *linkDestroyRoute) deleteLinkHandler(w http.ResponseWriter, r *http.Request) {
	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.DestroyLinkResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		_, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		ok := i.contentHandler.DestroyLink(id)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "删除失败"
			break
		}
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
