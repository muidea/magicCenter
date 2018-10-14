package route

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"strconv"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/content/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// AppendMediaRoute 追加User Route
func AppendMediaRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, fileRegistryHandler common.FileRegistryHandler, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateGetMediaByIDRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetMediaListRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateMediaRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateBatchCreateMediaRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateMediaRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyMediaRoute(contentHandler, accountHandler, fileRegistryHandler, sessionRegistry)
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

// CreateBatchCreateMediaRoute 新建CreateBatchCreateMediaRoute Route
func CreateBatchCreateMediaRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaBatchCreateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateMediaRoute UpdateMediaRoute Route
func CreateUpdateMediaRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaUpdateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyMediaRoute DestroyMediaRoute Route
func CreateDestroyMediaRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, fileRegistryHandler common.FileRegistryHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := mediaDestroyRoute{contentHandler: contentHandler, accountHandler: accountHandler, fileRegistryHandler: fileRegistryHandler, sessionRegistry: sessionRegistry}
	return &i
}

type mediaGetByIDRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
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

	result := common_def.QueryMediaResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		media, ok := i.contentHandler.GetMediaByID(id)
		if ok {
			user, _ := i.accountHandler.FindUserByID(media.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(media.Catalog)

			result.Media.MediaDetail = media
			result.Media.Creater = user.User
			result.Media.Catalog = catalogSummarys
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

type mediaGetListRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
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

	result := common_def.QueryMediaListResult{Media: []model.SummaryView{}}
	for true {
		filter := &common_def.Filter{}
		filter.Parse(r)

		catalog, err := common_def.DecodeStrictCatalog(r)
		if catalog == nil && err == nil {
			medias, total := i.contentHandler.GetAllMedia(filter.PageFilter)
			for _, val := range medias {
				media := model.SummaryView{}
				user, _ := i.accountHandler.FindUserByID(val.Creater)
				catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

				media.Summary = val
				media.Creater = user.User
				media.Catalog = catalogSummarys

				result.Media = append(result.Media, media)
			}
			result.Total = total
			result.ErrorCode = common_def.Success
			break
		} else if err != nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			break
		}

		medias, total := i.contentHandler.GetMediaByCatalog(*catalog, filter.PageFilter)
		for _, val := range medias {
			media := model.SummaryView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

			media.Summary = val
			media.Creater = user.User
			media.Catalog = catalogSummarys

			result.Media = append(result.Media, media)
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

type mediaCreateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
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
	result := common_def.CreateMediaResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.CreateMediaParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		createDate := time.Now().Format("2006-01-02 15:04:05")
		media, ok := i.contentHandler.CreateMedia(param.Name, param.Description, param.FileToken, createDate, param.Catalog, param.Expiration, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "新建失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)

		result.Media.Summary = media
		result.Media.Creater = user
		result.Media.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type mediaBatchCreateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *mediaBatchCreateRoute) Method() string {
	return common.POST
}

func (i *mediaBatchCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostBatchMedia)
}

func (i *mediaBatchCreateRoute) Handler() interface{} {
	return i.createBatchMediaHandler
}

func (i *mediaBatchCreateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *mediaBatchCreateRoute) createBatchMediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createBatchMediaHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.BatchCreateMediaResult{Medias: []model.SummaryView{}}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.BatchCreateMediaParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		createDate := time.Now().Format("2006-01-02 15:04:05")
		mediaItems := []model.MediaItem{}
		for _, val := range param.Medias {
			item := model.MediaItem{Name: val.Name, FileToken: val.FileToken, Description: param.Description, Expiration: param.Expiration, Catalog: param.Catalog}

			mediaItems = append(mediaItems, item)
		}

		medias, ok := i.contentHandler.BatchCreateMedia(mediaItems, createDate, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "批量新建失败"
			break
		}

		for _, val := range medias {
			catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)
			view := model.SummaryView{Summary: val, Catalog: catalogSummarys, Creater: user}
			result.Medias = append(result.Medias, view)
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

type mediaUpdateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
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
	result := common_def.UpdateMediaResult{}
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

		param := &common_def.UpdateMediaParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		updateDate := time.Now().Format("2006-01-02 15:04:05")
		media := model.MediaDetail{}
		media.ID = id
		media.Name = param.Name
		media.FileToken = param.FileToken
		media.Description = param.Description
		media.Catalog = param.Catalog
		media.CreateDate = updateDate
		media.Expiration = param.Expiration
		media.Creater = user.ID
		summmary, ok := i.contentHandler.SaveMedia(media)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(param.Catalog)

		result.Media.Summary = summmary
		result.Media.Creater = user
		result.Media.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type mediaDestroyRoute struct {
	contentHandler      common.ContentHandler
	accountHandler      common.AccountHandler
	fileRegistryHandler common.FileRegistryHandler
	sessionRegistry     common.SessionRegistry
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
	result := common_def.DestroyMediaResult{}
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

		media, ok := i.contentHandler.GetMediaByID(id)
		if !ok {
			result.ErrorCode = common_def.NoExist
			result.Reason = "对象不存在"
			break
		}

		ok = i.contentHandler.DestroyMedia(id)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "删除失败"
			break
		}

		i.fileRegistryHandler.RemoveFile(media.FileToken)

		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
