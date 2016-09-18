package ui

import (
	"encoding/json"
	"html"
	"html/template"
	"log"
	"magiccenter/common"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/configuration"
	"magiccenter/kernel/modules/content/bll"
	"magiccenter/session"
	"net/http"
	"strconv"

	"muidea.com/util"
)

// ManageArticleView 文章管理视图
type ManageArticleView struct {
	Articles []model.ArticleSummary
	Catalogs []model.Catalog
	Users    []model.User
}

// AllArticleSummaryList 全部ArticleSummary列表
type AllArticleSummaryList struct {
	Articles []model.ArticleSummary
}

// SingleArticle 单篇Article信息
type SingleArticle struct {
	common.Result
	Article model.Article
}

// ManageArticleViewHandler 文章管理主界面
// 显示Article列表信息
// 返回html页面
//
func ManageArticleViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageArticleViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/content/article.html")
	if err != nil {
		panic("parse files failed, err:" + err.Error())
	}

	view := ManageArticleView{}
	view.Articles = bll.QueryAllArticleSummary()
	view.Catalogs = bll.QueryAllCatalogList()
	view.Users = commonbll.QueryAllUserList()

	t.Execute(w, view)
}

// QueryAllArticleSummaryHandler 查询所有Article
// 返回json
func QueryAllArticleSummaryHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryAllArticleSummaryHandler")

	result := AllArticleSummaryList{}
	result.Articles = bll.QueryAllArticleSummary()

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// QueryArticleHandler 查询指定Article内容
// 返回json
func QueryArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("QueryArticleHandler")

	result := SingleArticle{}

	for true {
		params := util.SplitParam(r.URL.RawQuery)
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		article, found := bll.QueryArticleByID(aid)
		if !found {
			result.ErrCode = 1
			result.Reason = "查询文章失败"
			break
		}

		article.Content = html.UnescapeString(article.Content)
		result.Article = article
		result.ErrCode = 0
		result.Reason = "查询成功"

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// DeleteArticleHandler 删除指定Article
// 返回json
func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("DeleteArticleHandler")

	result := common.Result{}

	for true {
		params := util.SplitParam(r.URL.RawQuery)
		id, found := params["id"]
		if !found {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		if !bll.DeleteArticle(aid) {
			result.ErrCode = 1
			result.Reason = "删除失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "删除成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// AjaxArticleHandler 保存Article
// 返回json
func AjaxArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("AjaxArticleHandler")

	authID, found := configuration.GetOption(configuration.AuthorithID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}

	session := session.GetSession(w, r)
	user, found := session.GetOption(authID)
	if !found {
		panic("unexpected, must login system first.")
	}

	result := common.Result{}
	for true {
		err := r.ParseMultipartForm(0)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		id := r.FormValue("article-id")
		title := r.FormValue("article-title")
		content := html.EscapeString(r.FormValue("article-content"))
		catalog := r.MultipartForm.Value["article-catalog"]

		aid, err := strconv.Atoi(id)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效请求数据"
			break
		}

		catalogs := []int{}
		for _, ca := range catalog {
			cid, err := strconv.Atoi(ca)
			if err != nil {
				result.ErrCode = 1
				result.Reason = "无效请求数据"
				break
			}

			catalogs = append(catalogs, cid)
		}

		if !bll.SaveArticle(aid, title, content, user.(model.UserDetail).ID, catalogs) {
			result.ErrCode = 1
			result.Reason = "保存失败"
			break
		}

		result.ErrCode = 0
		result.Reason = "保存成功"
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
