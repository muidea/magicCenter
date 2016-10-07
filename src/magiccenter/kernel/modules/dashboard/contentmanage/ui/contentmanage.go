package ui

import (
	"html/template"
	"log"
	"magiccenter/common/model"
	"net/http"
)

// ContentManageView Content管理视图内容
// Blocks Module定义的功能块列表
// Pages Module定义的页面URL
type ContentManageView struct {
	Blocks []model.Block
	Pages  []string
}

// ItemList Item列表
type ItemList struct {
	Items []model.Item
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
	//view.Modules = bll.QueryAllModules()
	//view.DefaultModule, _ = configuration.GetOption(configuration.SysDefaultModule)

	t.Execute(w, view)
}

// BlockItemActionHandler Block拥有Item处理器
func BlockItemActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("BlockItemActionHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/module/content.html")
	if err != nil {
		panic("parse files failed")
	}

	view := ContentManageView{}
	//view.Modules = bll.QueryAllModules()
	//view.DefaultModule, _ = configuration.GetOption(configuration.SysDefaultModule)

	t.Execute(w, view)
}
