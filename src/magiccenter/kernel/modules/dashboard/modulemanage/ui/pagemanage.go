package ui

import (
	"html/template"
	"log"
	"magiccenter/kernel/modules/dashboard/modulemanage/bll"
	"magiccenter/kernel/modules/dashboard/modulemanage/model"
	"net/http"
)

// PageManageView Page管理视图
type PageManageView struct {
	Modules []model.Module
}

// PageManageViewHandler Page管理视图处理器
func PageManageViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("PageManageViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/page.html")
	if err != nil {
		panic("parse files failed")
	}

	view := PageManageView{}
	view.Modules = bll.QueryAllModules()

	t.Execute(w, view)
}
