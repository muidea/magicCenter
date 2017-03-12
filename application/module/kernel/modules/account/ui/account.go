package ui

import (
	"html/template"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/account/bll"
)

// ManageUserView 用户管理视图数据
type ManageUserView struct {
	Users []model.UserDetail
}

// ManageUserViewHandler 用户管理视图处理器
func ManageUserViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("ManageUserViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	htmlFile := system.GetHTMLPath("kernel/modules/account/account.html")
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	view := ManageUserView{}
	view.Users = bll.QueryAllUserDetail()

	t.Execute(w, view)
}
