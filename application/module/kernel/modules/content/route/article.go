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

type articleGetByIDResult struct {
	common.Result
	Article model.ArticleDetailView `json:"article"`
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
	return common.VisitorAuthGroup.ID
}

func (i *articleGetByIDRoute) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getArticleHandler")

	result := articleGetByIDResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrorCode = common.Failed
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

type articleGetListRoute struct {
	contentHandler common.ContentHandler
	accountHandler common.AccountHandler
}

type articleGetListResult struct {
	common.Result
	Article []model.SummaryView `json:"article"`
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
	return common.VisitorAuthGroup.ID
}

func (i *articleGetListRoute) getArticleListHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getArticleListHandler")

	result := articleGetListResult{}
	for true {
		catalog := r.URL.Query().Get("catalog")
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

			result.ErrorCode = common.Success
			break
		}

		id, err := strconv.Atoi(catalog)
		if err != nil {
			result.ErrorCode = common.Failed
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

		result.ErrorCode = common.Success
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

type articleCreateParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Catalog []int  `json:"catalog"`
}

type articleCreateResult struct {
	common.Result
	Article model.SummaryView `json:"article"`
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
	return common.UserAuthGroup.ID
}

func (i *articleCreateRoute) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := articleCreateResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrorCode = common.Failed
			result.Reason = "无效权限"
			break
		}

		param := &articleCreateParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		createDate := time.Now().Format("2006-01-02 15:04:05")
		article, ok := i.contentHandler.CreateArticle(param.Title, param.Content, createDate, param.Catalog, user.ID)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "新建失败"
			break
		}

		catalogs := i.contentHandler.GetCatalogs(article.Catalog)

		result.Article.Summary = article
		result.Article.Creater = user
		result.Article.Catalog = catalogs
		result.ErrorCode = common.Success
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

type articleUpdateParam articleCreateParam

type articleUpdateResult struct {
	common.Result
	Article model.SummaryView `json:"article"`
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
	return common.UserAuthGroup.ID
}

func (i *articleUpdateRoute) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := articleCreateResult{}
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

		param := &articleUpdateParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}
		article := model.ArticleDetail{}
		article.ID = id
		article.Name = param.Title
		article.Content = param.Content
		article.Catalog = param.Catalog
		article.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		article.Creater = user.ID
		summmary, ok := i.contentHandler.SaveArticle(article)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "更新失败"
			break
		}

		catalogs := i.contentHandler.GetCatalogs(summmary.Catalog)

		result.Article.Summary = summmary
		result.Article.Creater = user
		result.Article.Catalog = catalogs
		result.ErrorCode = common.Success
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

type articleDestroyResult struct {
	common.Result
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
	return common.MaintainerAuthGroup.ID
}

func (i *articleDestroyRoute) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := articleCreateResult{}
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

		ok := i.contentHandler.DestroyArticle(id)
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
