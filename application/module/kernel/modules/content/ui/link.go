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

// ManageLinkView Link管理视图
type ManageLinkView struct {
	Links    []model.Link
	Catalogs []model.Catalog
	Users    []model.User
}

// ManageLinkViewHandler 链接管理视图处理器
// Link管理主界面
// 显示Link列表信息
// 返回html页面
//
func ManageLinkViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageLinkViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/content/link.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageLinkView{}
	view.Links = bll.QueryAllLink()
	view.Catalogs = bll.QueryAllCatalog()
	view.Users, _ = commonbll.QueryAllUser()

	t.Execute(w, view)
}
