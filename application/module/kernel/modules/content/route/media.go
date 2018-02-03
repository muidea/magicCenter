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

// AppendMediaRoute 追加User Route
func AppendMediaRoute(routes []common.Route, contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateGetMediaByIDRoute(contentHandler)
	routes = append(routes, rt)

	rt = CreateGetMediaListRoute(contentHandler)
	routes = append(routes, rt)

	rt = CreateCreateMediaRoute(contentHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateMediaRoute(contentHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyMediaRoute(contentHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetMediaByIDRoute 新建GetMedia Route
func CreateGetMediaByIDRoute(contentHandler common.ContentHandler) common.Route {
	i := mediaGetByIDRoute{contentHandler: contentHandler}
	return &i
}

// CreateGetMediaListRoute 新建GetAllMedia Route
func CreateGetMediaListRoute(contentHandler common.ContentHandler) common.Route {
	i := mediaGetListRoute{contentHandler: contentHandler}
	return &i
}

// CreateCreateMediaRoute 新建CreateMediaRoute Route
func CreateCreateMediaRoute(contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaCreateRoute{contentHandler: contentHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateMediaRoute UpdateMediaRoute Route
func CreateUpdateMediaRoute(contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaUpdateRoute{contentHandler: contentHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyMediaRoute DestroyMediaRoute Route
func CreateDestroyMediaRoute(contentHandler common.ContentHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaDestroyRoute{contentHandler: contentHandler, sessionRegistry: sessionRegistry}
	return &i
}

type mediaGetByIDRoute struct {
	contentHandler common.ContentHandler
}

type mediaGetByIDResult struct {
	common.Result
	Media model.MediaDetail
}

func (i *mediaGetByIDRoute) Method() string {
	return common.GET
}

func (i *mediaGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetMediaDetail)
}

func (i *mediaGetByIDRoute) Handler() interface{} {
	return i.getMediaHandler
}

func (i *mediaGetByIDRoute) AuthGroup() int {
	return common.VisitorAuthGroup.ID
}

func (i *mediaGetByIDRoute) getMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getMediaHandler")

	result := mediaGetByIDResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
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

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type mediaGetListRoute struct {
	contentHandler common.ContentHandler
}

type mediaGetListResult struct {
	common.Result
	Media []model.Summary
}

func (i *mediaGetListRoute) Method() string {
	return common.GET
}

func (i *mediaGetListRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetMediaList)
}

func (i *mediaGetListRoute) Handler() interface{} {
	return i.getMediaListHandler
}

func (i *mediaGetListRoute) AuthGroup() int {
	return common.VisitorAuthGroup.ID
}

func (i *mediaGetListRoute) getMediaListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getMediaListHandler")

	result := mediaGetListResult{}
	for true {
		catalog := r.URL.Query()["catalog"]
		if len(catalog) < 1 {
			result.Media = i.contentHandler.GetAllMedia()
			result.ErrCode = 0
			break
		}

		id, err := strconv.Atoi(catalog[0])
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		result.Media = i.contentHandler.GetMediaByCatalog(id)
		result.ErrCode = 0
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

func (i *mediaCreateRoute) Method() string {
	return common.POST
}

func (i *mediaCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostMedia)
}

func (i *mediaCreateRoute) Handler() interface{} {
	return i.createMediaHandler
}

func (i *mediaCreateRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
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
		name := r.FormValue("name")
		url := r.FormValue("url")
		desc := r.FormValue("desc")
		var catalogs []int
		err := json.Unmarshal([]byte(r.FormValue("catalog")), &catalogs)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}
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

func (i *mediaUpdateRoute) Method() string {
	return common.PUT
}

func (i *mediaUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutMedia)
}

func (i *mediaUpdateRoute) Handler() interface{} {
	return i.updateMediaHandler
}

func (i *mediaUpdateRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *mediaUpdateRoute) updateMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
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
		media := model.MediaDetail{}
		media.ID = id
		media.Name = r.FormValue("name")
		media.URL = r.FormValue("url")
		media.Desc = r.FormValue("desc")
		var catalogs []int
		err = json.Unmarshal([]byte(r.FormValue("catalog")), &catalogs)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}
		media.Catalog = catalogs
		media.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		media.Creater = user.ID
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

func (i *mediaDestroyRoute) Method() string {
	return common.DELETE
}

func (i *mediaDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteMedia)
}

func (i *mediaDestroyRoute) Handler() interface{} {
	return i.deleteMediaHandler
}

func (i *mediaDestroyRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *mediaDestroyRoute) deleteMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
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
