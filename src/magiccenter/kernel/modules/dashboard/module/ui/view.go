package ui

import (
	"html/template"
	"log"
	"magiccenter/kernel/dashboard/module/bll"
	"magiccenter/kernel/dashboard/module/model"
	"net/http"
)

// ModuleManageView Module管理视图内容
type ModuleManageView struct {
	Modules       []model.Module
	DefaultModule string
}

// ModuleManageViewHandler Module管理视图处理器
func ModuleManageViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModuleManageViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/manage.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ModuleManageView{}
	view.Modules = bll.QueryAllModules()
	view.DefaultModule, _ = configuration.GetOption(configuration.SYS_DEFULTMODULE)

	t.Execute(w, view)
}
