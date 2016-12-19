package ui

import (
	"html/template"
	"log"
	commonbll "magiccenter/common/bll"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/content/bll"
	"magiccenter/system"
	"net/http"
)

// ManageArticleView 文章管理视图
type ManageArticleView struct {
	Articles []model.ArticleSummary
	Catalogs []model.Catalog
	Users    []model.User
}

// ManageArticleViewHandler 文章管理主界面
// 显示Article列表信息
// 返回html页面
//
func ManageArticleViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageArticleViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/content/article.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageArticleView{}
	view.Articles = bll.QueryAllArticleSummary()
	view.Catalogs = bll.QueryAllCatalog()
	view.Users, _ = commonbll.QueryAllUser()

	t.Execute(w, view)
}
