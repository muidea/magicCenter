package ui

import (
	"html/template"
	"log"
	"magiccenter/common/model"
	"magiccenter/kernel/modules/account/bll"
	"magiccenter/system"
	"net/http"
)

// ManageGroupView 分组管理视图
type ManageGroupView struct {
	Groups []model.Group
}

// ManageGroupViewHandler 分组管理视图处理器
func ManageGroupViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageGroupViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/account/group.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageGroupView{}
	view.Groups = bll.QueryAllGroup()

	t.Execute(w, view)
}
