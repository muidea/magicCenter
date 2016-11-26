package ui

import (
	"encoding/json"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"net/http"
)

// ContentMetaList 元数据列表
type ContentMetaList struct {
	common.Result
	ContentMetaList []model.ContentMeta
}

// ArticleList 文章列表
type ArticleList struct {
	common.Result
	ArticleList []model.ArticleSummary
}

// CatalogList 分类列表
type CatalogList struct {
	common.Result
	CatalogList []model.Catalog
}

// LinkList 链接列表
type LinkList struct {
	common.Result
	LinkList []model.Link
}

// MediaList 分类列表
type MediaList struct {
	common.Result
	MediaList []model.MediaDetail
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

// GetContentArticleActionHandler 获取文章列表
func GetContentArticleActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("GetContentArticleActionHandler")

	result := ArticleList{}
	found := false
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
