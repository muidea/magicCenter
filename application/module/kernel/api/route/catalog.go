package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/net"
	"muidea.com/magicCenter/foundation/util"
)

// CreateGetCatalogRoute 新建GetCatalog Route
func CreateGetCatalogRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := catalogGetRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetAllCatalogRoute 新建GetAllCatalog Route
func CreateGetAllCatalogRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := catalogGetAllRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetByCatalogCatalogRoute 新建GetByCatalogCatalogRoute Route
func CreateGetByCatalogCatalogRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := catalogGetByCatalogRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateCreateCatalogRoute 新建CreateCatalogRoute Route
func CreateCreateCatalogRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := catalogCreateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateUpdateCatalogRoute UpdateCatalogRoute Route
func CreateUpdateCatalogRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := catalogUpdateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateDestroyCatalogRoute DestroyCatalogRoute Route
func CreateDestroyCatalogRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := catalogDestroyRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

type catalogGetRoute struct {
	contentHandler common.ContentHandler
}

type catalogGetResult struct {
	common.Result
	Catalog model.CatalogDetail
}

func (i *catalogGetRoute) Type() string {
	return common.GET
}

func (i *catalogGetRoute) Pattern() string {
	return "content/catalog/[0-9]*/"
}

func (i *catalogGetRoute) Handler() interface{} {
	return i.getCatalogHandler
}

func (i *catalogGetRoute) getCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getCatalogHandler")

	result := catalogGetResult{}
	value, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if ok {
			id, err := strconv.Atoi(value)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			catalog, ok := i.contentHandler.GetCatalogByID(id)
			if ok {
				result.Catalog = catalog
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

type catalogGetAllRoute struct {
	contentHandler common.ContentHandler
}

type catalogGetAllResult struct {
	common.Result
	Catalog []model.Catalog
}

func (i *catalogGetAllRoute) Type() string {
	return common.GET
}

func (i *catalogGetAllRoute) Pattern() string {
	return "content/catalog/"
}

func (i *catalogGetAllRoute) Handler() interface{} {
	return i.getAllCatalogHandler
}

func (i *catalogGetAllRoute) getAllCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllCatalogHandler")

	result := catalogGetAllResult{}
	for true {
		result.Catalog = i.contentHandler.GetAllCatalog()
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type catalogGetByParentRoute struct {
	contentHandler common.ContentHandler
}

type catalogGetByParentResult struct {
	common.Result
	Catalog []model.Catalog
}

func (i *catalogGetByParentRoute) Type() string {
	return common.GET
}

func (i *catalogGetByParentRoute) Pattern() string {
	return "content/catalog/?parent=[0-9]*"
}

func (i *catalogGetByParentRoute) Handler() interface{} {
	return i.getByParentCatalogHandler
}

func (i *catalogGetByParentRoute) getByParentCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getByParentCatalogHandler")

	result := catalogGetByCatalogResult{}
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

			result.Catalog = i.contentHandler.GetCatalogByParent(id)
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

type catalogCreateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type catalogCreateResult struct {
	common.Result
	Catalog model.Catalog
}

func (i *catalogCreateRoute) Type() string {
	return common.POST
}

func (i *catalogCreateRoute) Pattern() string {
	return "content/catalog/"
}

func (i *catalogCreateRoute) Handler() interface{} {
	return i.createCatalogHandler
}

func (i *catalogCreateRoute) createCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createCatalogHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := catalogCreateResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrCode = 1
			result.Reason = "无效权限"
			break
		}

		r.ParseForm()

		title := r.FormValue("catalog-name")
		parents, _ := util.Str2IntArray(r.FormValue("catalog-parent"))
		catalog, ok := i.contentHandler.CreateCatalog(name, parents, user.ID)
		if !ok {
			result.ErrCode = 1
			result.Reason = "新建失败"
			break
		}
		result.ErrCode = 0
		result.Catalog = catalog
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
	sessionRegistry common.SessionRegistry
}

type catalogUpdateResult struct {
	common.Result
	Catalog model.Catalog
}

func (i *catalogUpdateRoute) Type() string {
	return common.PUT
}

func (i *catalogUpdateRoute) Pattern() string {
	return "content/catalog/[0-9]*/"
}

func (i *catalogUpdateRoute) Handler() interface{} {
	return i.updateCatalogHandler
}

func (i *catalogUpdateRoute) updateCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateCatalogHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := catalogCreateResult{}
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
		catalog := model.CatalogDetail{}
		catalog.ID = id
		catalog.Name = r.FormValue("catalog-name")
		catalog.Parent, _ = util.Str2IntArray(r.FormValue("catalog-parent"))
		catalog.Creater = user.ID
		summmary, ok := i.contentHandler.SaveCatalog(catalog)
		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}
		result.ErrCode = 0
		result.Catalog = summmary
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
	sessionRegistry common.SessionRegistry
}

type catalogDestroyResult struct {
	common.Result
}

func (i *catalogDestroyRoute) Type() string {
	return common.DELETE
}

func (i *catalogDestroyRoute) Pattern() string {
	return "content/catalog/[0-9]*/"
}

func (i *catalogDestroyRoute) Handler() interface{} {
	return i.deleteCatalogHandler
}

func (i *catalogDestroyRoute) deleteCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteCatalogHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := catalogCreateResult{}
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

		ok := i.contentHandler.DestroyCatalog(id)
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
