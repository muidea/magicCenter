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

// ManageMediaView Media管理视图
type ManageMediaView struct {
	Medias   []model.MediaDetail
	Catalogs []model.Catalog
	Users    []model.User
}

// ManageMediaViewHandler Media管理主界面处理器
// 显示Media列表信息
// 返回html页面
//
func ManageMediaViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageMediaViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/content/media.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}
	view := ManageMediaView{}
	view.Medias = bll.QueryAllMedia()
	view.Catalogs = bll.QueryAllCatalog()
	view.Users, _ = commonbll.QueryAllUser()

	t.Execute(w, view)
}
