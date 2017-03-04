package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/system"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// ContentMetaList 元数据列表
type ContentMetaList struct {
	common.Result
	ContentMetaList []model.ContentMeta
}

// GetContentMetadataListActionHandler 获取Content元数据列表
func GetContentMetadataListActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentMetadataListActionHandler")

	result := ContentMetaList{}
	found := false
	result.ContentMetaList, found = commonbll.QueryContentMetas()
	if found {
		result.ErrCode = 0
	} else {
		result.ErrCode = 1
		result.Reason = "查询失败"
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// ArticleList 文章列表
type ArticleList struct {
	common.Result
	ArticleList []model.ArticleSummary
}

// SingleArticle 单篇文章
type SingleArticle struct {
	common.Result
	Article model.Article
}

// GetContentArticleActionHandler 获取文章
func GetContentArticleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentArticleActionHandler")
	params := util.SplitParam(r.URL.RawQuery)
	aid, found := params["id"]
	if !found {
		result := ArticleList{}
		result.ArticleList, found = commonbll.QueryContentArticles()
		if found {
			result.ErrCode = 0
		} else {
			result.ErrCode = 1
			result.Reason = "查询失败"
		}

		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	} else {
		result := SingleArticle{}
		for true {
			id, err := strconv.Atoi(aid)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "参数非法"
				break
			}

			result.Article, found = commonbll.QuerySingleArticle(id)
			if !found {
				result.ErrCode = 1
				result.Reason = "无效参数"
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
}

// SingleArticleSummary 单篇文章
type SingleArticleSummary struct {
	common.Result
	Article model.ArticleSummary
}

// PostContentArticleActionHandler 新增文章
func PostContentArticleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostContentArticleActionHandler")

	session := system.GetSession(w, r)

	result := SingleArticleSummary{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		title := r.FormValue("article-title")
		content := r.FormValue("article-content")
		catalog := r.FormValue("article-catalog")
		catalogs, ret := util.Str2IntArray(catalog)
		if !ret || len(title) == 0 || len(content) == 0 {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Article, ret = commonbll.CreateArticle(title, content, catalogs, user.ID)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "新增文章失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// PutContentArticleActionHandler 更新文章
func PutContentArticleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentArticleActionHandler")

	session := system.GetSession(w, r)

	result := SingleArticleSummary{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		aid := r.FormValue("article-id")
		id, err := strconv.Atoi(aid)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		title := r.FormValue("article-title")
		content := r.FormValue("article-content")
		catalog := r.FormValue("article-catalog")
		catalogs, ret := util.Str2IntArray(catalog)
		if !ret || len(title) == 0 || len(content) == 0 {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		article := model.Article{}
		article.ID = id
		article.Title = title
		article.Content = content
		article.Catalog = catalogs
		article.Author = user.ID

		result.Article, ret = commonbll.UpdateArticle(article)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "更新文章失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteContentArticleActionHandler 删除文章
func DeleteContentArticleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteContentArticleActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	uid, found := params["id"]
	for true {
		if !found {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}
		id, err := strconv.Atoi(uid)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}

		found = commonbll.DeleteArticle(id)
		if !found {
			result.ErrCode = 1
			result.Reason = "删除文章失败"
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

// CatalogList 分类列表
type CatalogList struct {
	common.Result
	CatalogList []model.Catalog
}

// SingleCatalog 单个分类
type SingleCatalog struct {
	common.Result
	Catalog model.CatalogDetail
}

// GetContentCatalogActionHandler 获取文章列表
func GetContentCatalogActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentCatalogActionHandler")
	params := util.SplitParam(r.URL.RawQuery)
	cid, found := params["id"]
	if !found {
		result := CatalogList{}
		found := false
		result.CatalogList, found = commonbll.QueryContentCatalogs()
		if found {
			result.ErrCode = 0
		} else {
			result.ErrCode = 1
			result.Reason = "查询失败"
		}

		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	} else {
		result := SingleCatalog{}
		for true {
			id, err := strconv.Atoi(cid)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}

			found := false
			result.Catalog, found = commonbll.QuerySingleCatalog(id)
			if found {
				result.ErrCode = 0
			} else {
				result.ErrCode = 1
				result.Reason = "查询失败"
			}
			break
		}

		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	}
}

// PostContentCatalogActionHandler 新建分类
func PostContentCatalogActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostContentCatalogActionHandler")

	session := system.GetSession(w, r)

	result := SingleCatalog{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("catalog-name")
		parents := r.FormValue("catalog-parent")
		parent, ret := util.Str2IntArray(parents)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Catalog, ret = commonbll.CreateCatalog(name, parent, user.ID)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "新建分类失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// PutContentCatalogActionHandler 更新分类
func PutContentCatalogActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutContentCatalogActionHandler")

	session := system.GetSession(w, r)

	result := SingleCatalog{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		cid := r.FormValue("catalog-id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("catalog-name")
		parents := r.FormValue("catalog-parent")
		parent, ret := util.Str2IntArray(parents)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		catalog := model.CatalogDetail{}
		catalog.ID = id
		catalog.Name = name
		catalog.Parent = parent
		catalog.Creater = user.ID

		result.Catalog, ret = commonbll.UpdateCatalog(catalog)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "更新分类失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteContentCatalogActionHandler 删除分类
func DeleteContentCatalogActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteContentCatalogActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	cid, found := params["id"]
	for true {
		if !found {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		id, err := strconv.Atoi(cid)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		ret := commonbll.DeleteCatalog(id)
		if !ret {
			result.ErrCode = 1
			result.Reason = "删除分类失败"
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

// LinkList 链接列表
type LinkList struct {
	common.Result
	LinkList []model.Link
}

// SingleLink 单个Link
type SingleLink struct {
	common.Result
	Link model.Link
}

// GetContentLinkActionHandler 获取文章列表
func GetContentLinkActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentLinkActionHandler")
	params := util.SplitParam(r.URL.RawQuery)
	lid, found := params["id"]
	if !found {
		result := LinkList{}
		found := false
		result.LinkList, found = commonbll.QueryContentLinks()
		if found {
			result.ErrCode = 0
		} else {
			result.ErrCode = 1
			result.Reason = "查询失败"
		}
		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	} else {
		result := SingleLink{}
		for true {
			id, err := strconv.Atoi(lid)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}
			result.Link, found = commonbll.QuerySingleLink(id)
			if !found {
				result.ErrCode = 1
				result.Reason = "查询失败"
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
}

// PostContentLinkActionHandler 新建分类
func PostContentLinkActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostContentLinkActionHandler")

	session := system.GetSession(w, r)

	result := SingleLink{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("link-name")
		url := r.FormValue("link-url")
		logo := r.FormValue("link-logo")
		catalogs := r.FormValue("link-catalog")
		catalog, ret := util.Str2IntArray(catalogs)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Link, ret = commonbll.CreateLink(name, url, logo, catalog, user.ID)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "新建链接失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// PutContentLinkActionHandler 更新链接
func PutContentLinkActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutContentLinkActionHandler")

	session := system.GetSession(w, r)

	result := SingleLink{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		cid := r.FormValue("link-id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("link-name")
		url := r.FormValue("link-url")
		logo := r.FormValue("link-logo")
		catalogs := r.FormValue("link-catalog")
		catalog, ret := util.Str2IntArray(catalogs)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		lnk := model.Link{}
		lnk.ID = id
		lnk.Name = name
		lnk.URL = url
		lnk.Logo = logo
		lnk.Catalog = catalog
		lnk.Creater = user.ID

		result.Link, ret = commonbll.UpdateLink(lnk)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "更新链接失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteContentLinkActionHandler 删除链接
func DeleteContentLinkActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteContentLinkActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	cid, found := params["id"]
	for true {
		if !found {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		id, err := strconv.Atoi(cid)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		ret := commonbll.DeleteLink(id)
		if !ret {
			result.ErrCode = 1
			result.Reason = "删除链接失败"
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

// MediaList 分类列表
type MediaList struct {
	common.Result
	MediaList []model.MediaDetail
}

// SingleMedia 单个Media
type SingleMedia struct {
	common.Result
	Media model.MediaDetail
}

// GetContentMediaActionHandler 获取文章列表
func GetContentMediaActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentMediaActionHandler")
	params := util.SplitParam(r.URL.RawQuery)
	lid, found := params["id"]
	if !found {
		result := MediaList{}
		found := false
		result.MediaList, found = commonbll.QueryContentMedias()
		if found {
			result.ErrCode = 0
		} else {
			result.ErrCode = 1
			result.Reason = "查询失败"
		}
		b, err := json.Marshal(result)
		if err != nil {
			panic("json.Marshal, failed, err:" + err.Error())
		}

		w.Write(b)
	} else {
		result := SingleMedia{}
		for true {
			id, err := strconv.Atoi(lid)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效参数"
				break
			}
			result.Media, found = commonbll.QuerySingleMedia(id)
			if !found {
				result.ErrCode = 1
				result.Reason = "查询失败"
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
}

// PostContentMediaActionHandler 新建文件
func PostContentMediaActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PostContentMediaActionHandler")

	session := system.GetSession(w, r)

	result := SingleMedia{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("media-name")
		url := r.FormValue("media-url")
		typeName := r.FormValue("media-type")
		desc := r.FormValue("media-desc")
		catalogs := r.FormValue("media-catalog")
		catalog, ret := util.Str2IntArray(catalogs)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		result.Media, ret = commonbll.CreateMedia(name, url, typeName, desc, catalog, user.ID)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "新建文件失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// PutContentMediaActionHandler 更新文件
func PutContentMediaActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PutContentMediaActionHandler")

	session := system.GetSession(w, r)

	result := SingleMedia{}
	for true {
		user, found := session.GetAccount()
		if !found {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法请求"
			break
		}

		err := r.ParseForm()
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		cid := r.FormValue("media-id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			result.Result.ErrCode = 1
			result.Result.Reason = "非法参数"
			break
		}

		name := r.FormValue("media-name")
		url := r.FormValue("media-url")
		typeName := r.FormValue("media-type")
		desc := r.FormValue("media-desc")
		catalogs := r.FormValue("media-catalog")
		catalog, ret := util.Str2IntArray(catalogs)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "无效参数"
			break
		}

		mediaFile := model.MediaDetail{}
		mediaFile.ID = id
		mediaFile.Name = name
		mediaFile.URL = url
		mediaFile.Type = typeName
		mediaFile.Desc = desc
		mediaFile.Catalog = catalog
		mediaFile.Creater = user.ID

		result.Media, ret = commonbll.UpdateMedia(mediaFile)
		if !ret {
			result.Result.ErrCode = 1
			result.Result.Reason = "更新文件失败"
			break
		}

		result.Result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteContentMediaActionHandler 删除文件
func DeleteContentMediaActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteContentMediaActionHandler")

	result := common.Result{}
	params := util.SplitParam(r.URL.RawQuery)
	cid, found := params["id"]
	for true {
		if !found {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		id, err := strconv.Atoi(cid)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		ret := commonbll.DeleteMedia(id)
		if !ret {
			result.ErrCode = 1
			result.Reason = "删除文件失败"
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
