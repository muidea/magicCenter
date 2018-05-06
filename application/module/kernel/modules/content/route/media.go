package route

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/content/def"
	"muidea.com/magicCommon/foundation/net"
	common_const "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

// AppendMediaRoute 追加User Route
func AppendMediaRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateGetMediaByIDRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetMediaListRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateMediaRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateMediaRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyMediaRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetMediaByIDRoute 新建GetMedia Route
func CreateGetMediaByIDRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := mediaGetByIDRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetMediaListRoute 新建GetAllMedia Route
func CreateGetMediaListRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := mediaGetListRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateCreateMediaRoute 新建CreateMediaRoute Route
func CreateCreateMediaRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaCreateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateMediaRoute UpdateMediaRoute Route
func CreateUpdateMediaRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaUpdateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyMediaRoute DestroyMediaRoute Route
func CreateDestroyMediaRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaDestroyRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

type mediaGetByIDRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

type mediaGetByIDResult struct {
	common_result.Result
	Media model.MediaDetailView `json:"media"`
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
	return common_const.VisitorAuthGroup.ID
}

func (i *mediaGetByIDRoute) getMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getMediaHandler")

	result := mediaGetByIDResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		media, ok := i.contentHandler.GetMediaByID(id)
		if ok {
			user, _ := i.accountHandler.FindUserByID(media.Creater)
			catalogs := i.contentHandler.GetCatalogs(media.Catalog)

			result.Media.MediaDetail = media
			result.Media.Creater = user.User
			result.Media.Catalog = catalogs
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
	accountHandler common.AccountHandler
}

type mediaGetListResult struct {
	common_result.Result
	Media []model.SummaryView `json:"media"`
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
	return common_const.VisitorAuthGroup.ID
}

func (i *mediaGetListRoute) getMediaListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getMediaListHandler")

	result := mediaGetListResult{}
	for true {
		catalog := r.URL.Query().Get("catalog")
		if len(catalog) < 1 {
			medias := i.contentHandler.GetAllMedia()
			for _, val := range medias {
				media := model.SummaryView{}
				user, _ := i.accountHandler.FindUserByID(val.Creater)
				catalogs := i.contentHandler.GetCatalogs(val.Catalog)

				media.Summary = val
				media.Creater = user.User
				media.Catalog = catalogs

				result.Media = append(result.Media, media)
			}
			result.ErrorCode = common_result.Success
			break
		}

		id, err := strconv.Atoi(catalog)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		medias := i.contentHandler.GetMediaByCatalog(id)
		for _, val := range medias {
			media := model.SummaryView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogs := i.contentHandler.GetCatalogs(val.Catalog)

			media.Summary = val
			media.Creater = user.User
			media.Catalog = catalogs

			result.Media = append(result.Media, media)
		}
		result.ErrorCode = common_result.Success
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
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

type mediaCreateParam struct {
	Name        string          `json:"name"`
	URL         string          `json:"url"`
	Description string          `json:"description"`
	Catalog     []model.Catalog `json:"catalog"`
}

type mediaCreateResult struct {
	common_result.Result
	Media model.SummaryView `json:"media"`
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
	return common_const.UserAuthGroup.ID
}

func (i *mediaCreateRoute) createMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效权限"
			break
		}

		param := &mediaCreateParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		createDate := time.Now().Format("2006-01-02 15:04:05")
		catalogIds := []int{}
		catalogs, ok := i.contentHandler.UpdateCatalog(param.Catalog, createDate, user.ID)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "更新Catalog失败"
			break
		}
		for _, val := range catalogs {
			catalogIds = append(catalogIds, val.ID)
		}

		media, ok := i.contentHandler.CreateMedia(param.Name, param.URL, param.Description, createDate, catalogIds, user.ID)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "新建失败"
			break
		}

		result.Media.Summary = media
		result.Media.Creater = user
		result.Media.Catalog = catalogs
		result.ErrorCode = common_result.Success
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
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

type mediaUpdateParam mediaCreateParam

type mediaUpdateResult struct {
	common_result.Result
	Media model.SummaryView `json:"media"`
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
	return common_const.UserAuthGroup.ID
}

func (i *mediaUpdateRoute) updateMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效权限"
			break
		}

		param := &mediaUpdateParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		updateDate := time.Now().Format("2006-01-02 15:04:05")
		catalogIds := []int{}
		catalogs, ok := i.contentHandler.UpdateCatalog(param.Catalog, updateDate, user.ID)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "更新Catalog失败"
			break
		}
		for _, val := range catalogs {
			catalogIds = append(catalogIds, val.ID)
		}

		media := model.MediaDetail{}
		media.ID = id
		media.Name = param.Name
		media.URL = param.URL
		media.Description = param.Description
		media.Catalog = catalogIds
		media.CreateDate = updateDate
		media.Creater = user.ID
		summmary, ok := i.contentHandler.SaveMedia(media)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "更新失败"
			break
		}

		result.Media.Summary = summmary
		result.Media.Creater = user
		result.Media.Catalog = catalogs
		result.ErrorCode = common_result.Success
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
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

type mediaDestroyResult struct {
	common_result.Result
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
	return common_const.MaintainerAuthGroup.ID
}

func (i *mediaDestroyRoute) deleteMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := mediaCreateResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}
		_, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效权限"
			break
		}

		ok := i.contentHandler.DestroyMedia(id)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "删除失败"
			break
		}
		result.ErrorCode = common_result.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
