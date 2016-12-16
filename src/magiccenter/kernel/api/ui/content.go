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

// GetContentCatalogActionHandler 获取文章列表
func GetContentCatalogActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentCatalogActionHandler")

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
}

// LinkList 链接列表
type LinkList struct {
	common.Result
	LinkList []model.Link
}

// GetContentLinkActionHandler 获取文章列表
func GetContentLinkActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentLinkActionHandler")

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
}

// MediaList 分类列表
type MediaList struct {
	common.Result
	MediaList []model.MediaDetail
}

// GetContentMediaActionHandler 获取文章列表
func GetContentMediaActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentMediaActionHandler")

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
}
