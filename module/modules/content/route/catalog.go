package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/content/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// AppendCatalogRoute 追加User Route
func AppendCatalogRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateGetCatalogByIDRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = QueryGetCatalogByNameRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetCatalogListRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateCatalogRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateCatalogRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyCatalogRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetCatalogByIDRoute 新建GetCatalog Route
func CreateGetCatalogByIDRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := catalogGetByIDRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// QueryGetCatalogByNameRoute 新建QueryCatalog Route
func QueryGetCatalogByNameRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := catalogQueryByNameRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetCatalogListRoute 新建GetAllCatalog Route
func CreateGetCatalogListRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := catalogGetListRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateCreateCatalogRoute 新建CreateCatalogRoute Route
func CreateCreateCatalogRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := catalogCreateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateCatalogRoute UpdateCatalogRoute Route
func CreateUpdateCatalogRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := catalogUpdateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyCatalogRoute DestroyCatalogRoute Route
func CreateDestroyCatalogRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := catalogDestroyRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

type catalogGetByIDRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *catalogGetByIDRoute) Method() string {
	return common.GET
}

func (i *catalogGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetCatalogDetail)
}

func (i *catalogGetByIDRoute) Handler() interface{} {
	return i.getCatalogHandler
}

func (i *catalogGetByIDRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *catalogGetByIDRoute) getCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getCatalogHandler")

	result := common_def.QueryCatalogResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		catalog, ok := i.contentHandler.GetCatalogByID(id)
		if ok {
			user, _ := i.accountHandler.FindUserByID(catalog.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(catalog.Catalog)

			result.Catalog.CatalogDetail = catalog
			result.Catalog.Creater = user.User
			result.Catalog.Catalog = catalogSummarys
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

type catalogQueryByNameRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *catalogQueryByNameRoute) Method() string {
	return common.GET
}

func (i *catalogQueryByNameRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QueryCatalogByName)
}

func (i *catalogQueryByNameRoute) Handler() interface{} {
	return i.queryCatalogByNameHandler
}

func (i *catalogQueryByNameRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *catalogQueryByNameRoute) queryCatalogByNameHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryCatalogByNameHandler")

	result := common_def.QueryCatalogResult{}
	for true {
		name := r.URL.Query().Get("name")
		if name == "" {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			break
		}

		if strictCatalog == nil {
			strictCatalog = common_const.SystemContentCatalog.CatalogUnit()
		}

		catalog, ok := i.contentHandler.QueryCatalogByName(name, *strictCatalog)
		if ok {
			user, _ := i.accountHandler.FindUserByID(catalog.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(catalog.Catalog)

			result.Catalog.CatalogDetail = catalog
			result.Catalog.Creater = user.User
			result.Catalog.Catalog = catalogSummarys
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.NoExist
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

type catalogGetListRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *catalogGetListRoute) Method() string {
	return common.GET
}

func (i *catalogGetListRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetCatalogList)
}

func (i *catalogGetListRoute) Handler() interface{} {
	return i.getCatalogListHandler
}

func (i *catalogGetListRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *catalogGetListRoute) getCatalogListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getCatalogListHandler")

	result := common_def.QueryCatalogListResult{}
	for true {
		filter := &common_def.Filter{}
		filter.Parse(r)

		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if strictCatalog == nil && err == nil {
			catalogs, total := i.contentHandler.GetAllCatalog(filter.PageFilter)
			for _, val := range catalogs {
				catalog := model.SummaryView{}
				user, _ := i.accountHandler.FindUserByID(val.Creater)
				catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

				catalog.Summary = val
				catalog.Creater = user.User
				catalog.Catalog = catalogSummarys

				result.Catalog = append(result.Catalog, catalog)
			}
			result.Total = total

			result.ErrorCode = common_def.Success
			break
		}

		if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "无效参数"
			break
		}

		catalogs, total := i.contentHandler.GetCatalogByCatalog(*strictCatalog, filter.PageFilter)
		for _, val := range catalogs {
			catalog := model.SummaryView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

			catalog.Summary = val
			catalog.Creater = user.User
			catalog.Catalog = catalogSummarys

			result.Catalog = append(result.Catalog, catalog)
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

type catalogCreateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *catalogCreateRoute) Method() string {
	return common.POST
}

func (i *catalogCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostCatalog)
}

func (i *catalogCreateRoute) Handler() interface{} {
	return i.createCatalogHandler
}

func (i *catalogCreateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *catalogCreateRoute) createCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createCatalogHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.CreateCatalogResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.CreateCatalogParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		createDate := time.Now().Format("2006-01-02 15:04:05")
		catalog, ok := i.contentHandler.CreateCatalog(param.Name, param.Description, createDate, param.Catalog, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "新建失败"
			break
		}

		catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)

		result.Catalog.Summary = catalog
		result.Catalog.Creater = user
		result.Catalog.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type catalogUpdateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *catalogUpdateRoute) Method() string {
	return common.PUT
}

func (i *catalogUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutCatalog)
}

func (i *catalogUpdateRoute) Handler() interface{} {
	return i.updateCatalogHandler
}

func (i *catalogUpdateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *catalogUpdateRoute) updateCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateCatalogHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.UpdateCatalogResult{}
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

		param := &common_def.UpdateCatalogParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		updateDate := time.Now().Format("2006-01-02 15:04:05")
		catalog := model.CatalogDetail{}
		catalog.ID = id
		catalog.Name = param.Name
		catalog.Description = param.Description
		catalog.CreateDate = updateDate
		catalog.Catalog = param.Catalog
		catalog.Creater = user.ID
		summmary, ok := i.contentHandler.SaveCatalog(catalog)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)

		result.Catalog.Summary = summmary
		result.Catalog.Creater = user
		result.Catalog.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type catalogDestroyRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *catalogDestroyRoute) Method() string {
	return common.DELETE
}

func (i *catalogDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteCatalog)
}

func (i *catalogDestroyRoute) Handler() interface{} {
	return i.deleteCatalogHandler
}

func (i *catalogDestroyRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *catalogDestroyRoute) deleteCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteCatalogHandler")

	result := common_def.DestroyCatalogResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		ok := i.contentHandler.DestroyCatalog(id)
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
