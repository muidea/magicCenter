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

// AppendArticleRoute 追加User Route
func AppendArticleRoute(routes []common.Route, contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateGetArticleByIDRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateGetArticleListRoute(contentHandler, accountHandler)
	routes = append(routes, rt)

	rt = CreateCreateArticleRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateArticleRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDestroyArticleRoute(contentHandler, accountHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetArticleByIDRoute 新建GetArticle Route
func CreateGetArticleByIDRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := articleGetByIDRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetArticleListRoute 新建GetArticle Route
func CreateGetArticleListRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler) common.Route {
	i := articleGetListRoute{contentHandler: contentHandler, accountHandler: accountHandler}
	return &i
}

// CreateCreateArticleRoute 新建CreateArticleRoute Route
func CreateCreateArticleRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := articleCreateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateUpdateArticleRoute UpdateArticleRoute Route
func CreateUpdateArticleRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := articleUpdateRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

// CreateDestroyArticleRoute DestroyArticleRoute Route
func CreateDestroyArticleRoute(contentHandler common.ContentHandler, accountHandler common.AccountHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := articleDestroyRoute{contentHandler: contentHandler, accountHandler: accountHandler, sessionRegistry: sessionRegistry}
	return &i
}

type articleGetByIDRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *articleGetByIDRoute) Method() string {
	return common.GET
}

func (i *articleGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetArticleDetail)
}

func (i *articleGetByIDRoute) Handler() interface{} {
	return i.getArticleHandler
}

func (i *articleGetByIDRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *articleGetByIDRoute) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getArticleHandler")

	result := common_def.QueryArticleResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		article, ok := i.contentHandler.GetArticleByID(id)
		if ok {
			user, _ := i.accountHandler.FindUserByID(article.Creater)
			catalogs := i.contentHandler.GetCatalogs(article.Catalog)

			result.Article.ArticleDetail = article
			result.Article.Creater = user.User
			result.Article.Catalog = catalogs
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

type articleGetListRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

func (i *articleGetListRoute) Method() string {
	return common.GET
}

func (i *articleGetListRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetArticleList)
}

func (i *articleGetListRoute) Handler() interface{} {
	return i.getArticleListHandler
}

func (i *articleGetListRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *articleGetListRoute) getArticleListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getArticleListHandler")

	result := common_def.QueryArticleListResult{}
	for true {
		catalog := r.URL.Query().Get("strictCatalog")
		if catalog == "" {
			articles := i.contentHandler.GetAllArticle()
			for _, val := range articles {
				article := model.SummaryView{}
				user, _ := i.accountHandler.FindUserByID(val.Creater)
				catalogs := i.contentHandler.GetCatalogs(val.Catalog)

				article.Summary = val
				article.Creater = user.User
				article.Catalog = catalogs

				result.Article = append(result.Article, article)
			}

			result.ErrorCode = common_def.Success
			break
		}

		id, err := strconv.Atoi(catalog)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		articles := i.contentHandler.GetArticleByCatalog(id)
		for _, val := range articles {
			article := model.SummaryView{}
			user, _ := i.accountHandler.FindUserByID(val.Creater)
			catalogs := i.contentHandler.GetCatalogs(val.Catalog)

			article.Summary = val
			article.Creater = user.User
			article.Catalog = catalogs

			result.Article = append(result.Article, article)
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

type articleCreateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *articleCreateRoute) Method() string {
	return common.POST
}

func (i *articleCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostArticle)
}

func (i *articleCreateRoute) Handler() interface{} {
	return i.createArticleHandler
}

func (i *articleCreateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *articleCreateRoute) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.CreateArticleResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效权限"
			break
		}

		param := &common_def.CreateArticleParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		catalogID := common_const.SystemContentCatalog.ID
		catalog := r.URL.Query().Get("strictCatalog")
		if len(catalog) > 0 {
			catalogID, err = strconv.Atoi(catalog)
			if err != nil {
				result.ErrorCode = common_def.IllegalParam
				result.Reason = "非法参数"
				break
			}
		}

		description := "auto update catalog description"
		createDate := time.Now().Format("2006-01-02 15:04:05")
		catalogIds := []int{}
		catalogs, ok := i.contentHandler.UpdateCatalog(param.Catalog, catalogID, description, createDate, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新Catalog失败"
			break
		}
		for _, val := range catalogs {
			catalogIds = append(catalogIds, val.ID)
		}

		article, ok := i.contentHandler.CreateArticle(param.Name, param.Content, createDate, catalogIds, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "新建失败"
			break
		}

		result.Article.Summary = article
		result.Article.Creater = user
		result.Article.Catalog = catalogs
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type articleUpdateRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *articleUpdateRoute) Method() string {
	return common.PUT
}

func (i *articleUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutArticle)
}

func (i *articleUpdateRoute) Handler() interface{} {
	return i.updateArticleHandler
}

func (i *articleUpdateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *articleUpdateRoute) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.UpdateArticleResult{}
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

		param := &common_def.UpdateArticleParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		catalogID := common_const.SystemContentCatalog.ID
		catalog := r.URL.Query().Get("strictCatalog")
		if len(catalog) > 0 {
			catalogID, err = strconv.Atoi(catalog)
			if err != nil {
				result.ErrorCode = common_def.IllegalParam
				result.Reason = "非法参数"
				break
			}
		}

		description := "auto update catalog description"
		updateDate := time.Now().Format("2006-01-02 15:04:05")
		catalogIds := []int{}
		catalogs, ok := i.contentHandler.UpdateCatalog(param.Catalog, catalogID, description, updateDate, user.ID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新Catalog失败"
			break
		}

		for _, val := range catalogs {
			catalogIds = append(catalogIds, val.ID)
		}

		article := model.ArticleDetail{}
		article.ID = id
		article.Name = param.Name
		article.Content = param.Content
		article.Catalog = catalogIds
		article.CreateDate = updateDate
		article.Creater = user.ID
		summmary, ok := i.contentHandler.SaveArticle(article)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新失败"
			break
		}

		result.Article.Summary = summmary
		result.Article.Creater = user
		result.Article.Catalog = catalogs
		result.ErrorCode = common_def.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type articleDestroyRoute struct {
	contentHandler  common.ContentHandler
	accountHandler  common.AccountHandler
	sessionRegistry common.SessionRegistry
}

func (i *articleDestroyRoute) Method() string {
	return common.DELETE
}

func (i *articleDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteArticle)
}

func (i *articleDestroyRoute) Handler() interface{} {
	return i.deleteArticleHandler
}

func (i *articleDestroyRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *articleDestroyRoute) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler")

	result := common_def.DestoryArticleResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		ok := i.contentHandler.DestroyArticle(id)
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
