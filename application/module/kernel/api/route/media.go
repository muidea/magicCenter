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

// AppendMediaRoute 追加User Route
func AppendMediaRoute(routes []common.Route, modHub common.ModuleHub, sessionRegistry common.SessionRegistry) []common.Route {

	rt, _ := CreateGetMediaRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateGetAllMediaRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateGetByCatalogMediaRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateCreateMediaRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateUpdateMediaRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateDestroyMediaRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetMediaRoute 新建GetMedia Route
func CreateGetMediaRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := mediaGetRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetAllMediaRoute 新建GetAllMedia Route
func CreateGetAllMediaRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := mediaGetAllRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetByCatalogMediaRoute 新建GetByCatalogMediaRoute Route
func CreateGetByCatalogMediaRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := mediaGetByCatalogRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateCreateMediaRoute 新建CreateMediaRoute Route
func CreateCreateMediaRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := mediaCreateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateUpdateMediaRoute UpdateMediaRoute Route
func CreateUpdateMediaRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := mediaUpdateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateDestroyMediaRoute DestroyMediaRoute Route
func CreateDestroyMediaRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := mediaDestroyRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

type mediaGetRoute struct {
	contentHandler common.ContentHandler
}

type mediaGetResult struct {
	common.Result
	Media model.MediaDetail
}

func (i *mediaGetRoute) Type() string {
	return common.GET
}

func (i *mediaGetRoute) Pattern() string {
	return "content/media/[0-9]*/"
}

func (i *mediaGetRoute) Handler() interface{} {
	return i.getMediaHandler
}

func (i *mediaGetRoute) getMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getMediaHandler")

	result := mediaGetResult{}
	value, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if ok {
			id, err := strconv.Atoi(value)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			media, ok := i.contentHandler.GetMediaByID(id)
			if ok {
				result.Media = media
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

type mediaGetAllRoute struct {
	contentHandler common.ContentHandler
}

type mediaGetAllResult struct {
	common.Result
	Media []model.Summary
}

func (i *mediaGetAllRoute) Type() string {
	return common.GET
}

func (i *mediaGetAllRoute) Pattern() string {
	return "content/media/"
}

func (i *mediaGetAllRoute) Handler() interface{} {
	return i.getAllMediaHandler
}

func (i *mediaGetAllRoute) getAllMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllMediaHandler")

	result := mediaGetAllResult{}
	for true {
		result.Media = i.contentHandler.GetAllMedia()
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type mediaGetByCatalogRoute struct {
	contentHandler common.ContentHandler
}

type mediaGetByCatalogResult struct {
	common.Result
	Media []model.Summary
}

func (i *mediaGetByCatalogRoute) Type() string {
	return common.GET
}

func (i *mediaGetByCatalogRoute) Pattern() string {
	return "content/media/?catalog=[0-9]*"
}

func (i *mediaGetByCatalogRoute) Handler() interface{} {
	return i.getByCatalogMediaHandler
}

func (i *mediaGetByCatalogRoute) getByCatalogMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getByCatalogMediaHandler")

	result := mediaGetByCatalogResult{}
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

			result.Media = i.contentHandler.GetMediaByCatalog(id)
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

type mediaCreateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type mediaCreateResult struct {
	common.Result
	Media model.Summary
}

func (i *mediaCreateRoute) Type() string {
	return common.POST
}

func (i *mediaCreateRoute) Pattern() string {
	return "content/media/"
}

func (i *mediaCreateRoute) Handler() interface{} {
	return i.createMediaHandler
}

func (i *mediaCreateRoute) createMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrCode = 1
			result.Reason = "无效权限"
			break
		}

		r.ParseForm()
		name := r.FormValue("media-name")
		url := r.FormValue("media-url")
		desc := r.FormValue("media-desc")
		catalogs, _ := util.Str2IntArray(r.FormValue("media-catalog"))
		createDate := time.Now().Format("2006-01-02 15:04:05")
		media, ok := i.contentHandler.CreateMedia(name, url, desc, createDate, catalogs, user.ID)
		if !ok {
			result.ErrCode = 1
			result.Reason = "新建失败"
			break
		}
		result.ErrCode = 0
		result.Media = media
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type mediaUpdateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type mediaUpdateResult struct {
	common.Result
	Media model.Summary
}

func (i *mediaUpdateRoute) Type() string {
	return common.PUT
}

func (i *mediaUpdateRoute) Pattern() string {
	return "content/media/[0-9]*/"
}

func (i *mediaUpdateRoute) Handler() interface{} {
	return i.updateMediaHandler
}

func (i *mediaUpdateRoute) updateMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
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
		media := model.MediaDetail{}
		media.ID = id
		media.Name = r.FormValue("media-name")
		media.URL = r.FormValue("media-url")
		media.Desc = r.FormValue("media-desc")
		media.Catalog, _ = util.Str2IntArray(r.FormValue("media-catalog"))
		media.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		media.Author = user.ID
		summmary, ok := i.contentHandler.SaveMedia(media)
		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}
		result.ErrCode = 0
		result.Media = summmary
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type mediaDestroyRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type mediaDestroyResult struct {
	common.Result
}

func (i *mediaDestroyRoute) Type() string {
	return common.DELETE
}

func (i *mediaDestroyRoute) Pattern() string {
	return "content/media/[0-9]*/"
}

func (i *mediaDestroyRoute) Handler() interface{} {
	return i.deleteMediaHandler
}

func (i *mediaDestroyRoute) deleteMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
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

		ok := i.contentHandler.DestroyMedia(id)
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
