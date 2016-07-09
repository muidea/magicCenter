package ui

import (
	"html/template"
	"log"
	"magiccenter/kernel/module/bll"
	"magiccenter/kernel/module/model"
	"net/http"
)

type ModulePageView struct {
	Modules []model.Module
}

func ModulePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ModulePageHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/page.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ModulePageView{}
	view.Modules = bll.QueryAllModules()

	t.Execute(w, view)
}
