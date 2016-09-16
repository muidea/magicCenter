package ui

import (
	"html/template"
	"log"
	"magiccenter/kernel/modules/dashboard/modulemanage/bll"
	"magiccenter/kernel/modules/dashboard/modulemanage/model"
	"net/http"
)

// ContentManageView Page管理视图
type ContentManageView struct {
	Modules []model.Module
}

// ContentManageViewHandler Content管理视图处理器
func ContentManageViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ContentManageViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/content.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ContentManageView{}
	view.Modules = bll.QueryAllModules()

	t.Execute(w, view)
}
