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

// AppendArticleRoute 追加User Route
func AppendArticleRoute(routes []common.Route, modHub common.ModuleHub, sessionRegistry common.SessionRegistry) []common.Route {
	rt, _ := CreateGetArticleRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateGetAllArticleRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateGetByCatalogArticleRoute(modHub)
	routes = append(routes, rt)

	rt, _ = CreateCreateArticleRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateUpdateArticleRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	rt, _ = CreateDestroyArticleRoute(modHub, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateGetArticleRoute 新建GetArticle Route
func CreateGetArticleRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleGetRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetAllArticleRoute 新建GetAllArticle Route
func CreateGetAllArticleRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleGetAllRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateGetByCatalogArticleRoute 新建GetByCatalogArticleRoute Route
func CreateGetByCatalogArticleRoute(modHub common.ModuleHub) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleGetByCatalogRoute{contentHandler: endPoint.(common.ContentHandler)}
		return &i, true
	}

	return nil, false
}

// CreateCreateArticleRoute 新建CreateArticleRoute Route
func CreateCreateArticleRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleCreateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateUpdateArticleRoute UpdateArticleRoute Route
func CreateUpdateArticleRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleUpdateRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

// CreateDestroyArticleRoute DestroyArticleRoute Route
func CreateDestroyArticleRoute(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) (common.Route, bool) {
	mod, found := modHub.FindModule(common.CotentModuleID)
	if !found {
		return nil, false
	}

	endPoint := mod.EndPoint()
	switch endPoint.(type) {
	case common.ContentHandler:
		i := articleDestroyRoute{contentHandler: endPoint.(common.ContentHandler), sessionRegistry: sessionRegistry}
		return &i, true
	}

	return nil, false
}

type articleGetRoute struct {
	contentHandler common.ContentHandler
}

type articleGetResult struct {
	common.Result
	Article model.ArticleDetail
}

func (i *articleGetRoute) Type() string {
	return common.GET
}

func (i *articleGetRoute) Pattern() string {
	return "content/article/[0-9]*/"
}

func (i *articleGetRoute) Handler() interface{} {
	return i.getArticleHandler
}

func (i *articleGetRoute) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getArticleHandler")

	result := articleGetResult{}
	value, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	for true {
		if ok {
			id, err := strconv.Atoi(value)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			article, ok := i.contentHandler.GetArticleByID(id)
			if ok {
				result.Article = article
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

type articleGetAllRoute struct {
	contentHandler common.ContentHandler
}

type articleGetAllResult struct {
	common.Result
	Article []model.Summary
}

func (i *articleGetAllRoute) Type() string {
	return common.GET
}

func (i *articleGetAllRoute) Pattern() string {
	return "content/article/"
}

func (i *articleGetAllRoute) Handler() interface{} {
	return i.getAllArticleHandler
}

func (i *articleGetAllRoute) getAllArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAllArticleHandler")

	result := articleGetAllResult{}
	for true {
		result.Article = i.contentHandler.GetAllArticle()
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type articleGetByCatalogRoute struct {
	contentHandler common.ContentHandler
}

type articleGetByCatalogResult struct {
	common.Result
	Article []model.Summary
}

func (i *articleGetByCatalogRoute) Type() string {
	return common.GET
}

func (i *articleGetByCatalogRoute) Pattern() string {
	return "content/article/?catalog=[0-9]*"
}

func (i *articleGetByCatalogRoute) Handler() interface{} {
	return i.getByCatalogArticleHandler
}

func (i *articleGetByCatalogRoute) getByCatalogArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getByCatalogArticleHandler")

	result := articleGetByCatalogResult{}
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

			result.Article = i.contentHandler.GetArticleByCatalog(id)
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

type articleCreateRoute struct {
	contentHandler  common.ContentHandler
	sessionRegistry common.SessionRegistry
}

type articleCreateResult struct {
	common.Result
	Article model.Summary
}

func (i *articleCreateRoute) Type() string {
	return common.POST
}

func (i *articleCreateRoute) Pattern() string {
	return "content/article/"
}

func (i *articleCreateRoute) Handler() interface{} {
	return i.createArticleHandler
}

func (i *articleCreateRoute) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("createArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := articleCreateResult{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.ErrCode = 1
			result.Reason = "无效权限"
			break
		}

		r.ParseForm()

		title := r.FormValue("article-title")
		content := r.FormValue("article-content")
		catalogs, _ := util.Str2IntArray(r.FormValue("article-catalog"))
		createDate := time.Now().Format("2006-01-02 15:04:05")
		article, ok := i.contentHandler.CreateArticle(title, content, createDate, catalogs, user.ID)
		if !ok {
			result.ErrCode = 1
			result.Reason = "新建失败"
			break
		}
		result.ErrCode = 0
		result.Article = article
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
	sessionRegistry common.SessionRegistry
}

type articleUpdateResult struct {
	common.Result
	Article model.Summary
}

func (i *articleUpdateRoute) Type() string {
	return common.PUT
}

func (i *articleUpdateRoute) Pattern() string {
	return "content/article/[0-9]*/"
}

func (i *articleUpdateRoute) Handler() interface{} {
	return i.updateArticleHandler
}

func (i *articleUpdateRoute) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := articleCreateResult{}
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
		article := model.ArticleDetail{}
		article.ID = id
		article.Name = r.FormValue("article-title")
		article.Content = r.FormValue("article-content")
		article.Catalog, _ = util.Str2IntArray(r.FormValue("article-catalog"))
		article.CreateDate = time.Now().Format("2006-01-02 15:04:05")
		article.Author = user.ID
		summmary, ok := i.contentHandler.SaveArticle(article)
		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}
		result.ErrCode = 0
		result.Article = summmary
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
	sessionRegistry common.SessionRegistry
}

type articleDestroyResult struct {
	common.Result
}

func (i *articleDestroyRoute) Type() string {
	return common.DELETE
}

func (i *articleDestroyRoute) Pattern() string {
	return "content/article/[0-9]*/"
}

func (i *articleDestroyRoute) Handler() interface{} {
	return i.deleteArticleHandler
}

func (i *articleDestroyRoute) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteArticleHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := articleCreateResult{}
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

		ok := i.contentHandler.DestroyArticle(id)
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
