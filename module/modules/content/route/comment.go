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

// AppendCommentRoute 追加User Route
func AppendCommentRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {

	rt := CreateGetCommentListRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateCommentRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateCommentRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyCommentRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetCommentListRoute 新建GetAllComment Route
func CreateGetCommentListRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := commentGetListRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateCreateCommentRoute 新建CreateCommentRoute Route
func CreateCreateCommentRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := commentCreateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateCommentRoute UpdateCommentRoute Route
func CreateUpdateCommentRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := commentUpdateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyCommentRoute DestroyCommentRoute Route
func CreateDestroyCommentRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := commentDestroyRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

type commentGetListRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *commentGetListRoute) Method() string {
	return common.GET
}

func (i *commentGetListRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetCommentList)
}

func (i *commentGetListRoute) Handler() interface{} {
	return i.getCommentListHandler
}

func (i *commentGetListRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *commentGetListRoute) getCommentListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getCommentListHandler")

	result := common_def.QueryCommentListResult{}
	for true {
		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil || strictCatalog == nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			break
		}

		comments := i.contentHandler.GetCommentByCatalog(*strictCatalog)
		for _, val := range comments {
			comment := model.CommentDetailView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogSummarys := i.contentHandler.GetSummaryByIDs(val.Catalog)

			comment.CommentDetail = val
			comment.Creater = user.User
			comment.Catalog = catalogSummarys

			result.Comment = append(result.Comment, comment)
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

type commentCreateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *commentCreateRoute) Method() string {
	return common.POST
}

func (i *commentCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostComment)
}

func (i *commentCreateRoute) Handler() interface{} {
	return i.createCommentHandler
}

func (i *commentCreateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *commentCreateRoute) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createCommentHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.CreateCommentResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.CreateCommentParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil || strictCatalog == nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			break
		}

		catalogUnits := []model.CatalogUnit{}
		createDate := time.Now().Format("2006-01-02 15:04:05")
		catalogUnits = append(catalogUnits, model.CatalogUnit{ID: strictCatalog.ID, Type: strictCatalog.Type})

		comment, ok := i.contentHandler.CreateComment(param.Subject, param.Content, createDate, catalogUnits, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "新建失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(catalogUnits)

		result.Comment.Summary = comment
		result.Comment.Creater = user
		result.Comment.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type commentUpdateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *commentUpdateRoute) Method() string {
	return common.PUT
}

func (i *commentUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutComment)
}

func (i *commentUpdateRoute) Handler() interface{} {
	return i.updateCommentHandler
}

func (i *commentUpdateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *commentUpdateRoute) updateCommentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateCommentHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.UpdateCommentResult{}
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

		param := &common_def.UpdateCommentParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}
		strictCatalog, err := common_def.DecodeStrictCatalog(r)
		if err != nil || strictCatalog == nil {
			result.ErrorCode = common_def.IllegalParam
			result.Reason = "非法参数"
			break
		}

		catalogUnits := []model.CatalogUnit{}
		updateDate := time.Now().Format("2006-01-02 15:04:05")
		catalogUnits = append(catalogUnits, model.CatalogUnit{ID: strictCatalog.ID, Type: strictCatalog.Type})

		comment := model.CommentDetail{}
		comment.ID = id
		comment.Subject = param.Subject
		comment.Content = param.Content
		comment.Catalog = catalogUnits
		comment.CreateDate = updateDate
		comment.Creater = user.ID
		summmary, ok := i.contentHandler.SaveComment(comment)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新失败"
			break
		}
		catalogSummarys := i.contentHandler.GetSummaryByIDs(catalogUnits)

		result.Comment.Summary = summmary
		result.Comment.Creater = user
		result.Comment.Catalog = catalogSummarys
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type commentDestroyRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *commentDestroyRoute) Method() string {
	return common.DELETE
}

func (i *commentDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteComment)
}

func (i *commentDestroyRoute) Handler() interface{} {
	return i.deleteCommentHandler
}

func (i *commentDestroyRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *commentDestroyRoute) deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteCommentHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.DestroyCommentResult{}
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

		ok := i.contentHandler.DestroyComment(id)
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
