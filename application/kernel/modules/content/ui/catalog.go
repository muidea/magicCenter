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

// ManageCatalogView 分类视图
type ManageCatalogView struct {
	Catalogs []model.CatalogDetail
	Users    []model.User
}

// ManageCatalogViewHandler 分类管理主界面
// 显示Catalog列表信息
// 返回html页面
//
func ManageCatalogViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageCatalogViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/content/catalog.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageCatalogView{}
	view.Catalogs = bll.QueryAllCatalogDetail()
	view.Users, _ = commonbll.QueryAllUser()

	t.Execute(w, view)
}
