package ui

import (
	"html/template"
	"log"
	"magiccenter/common/model"
	"net/http"
)

// UserProfileView 用户Profile视图
type UserProfileView struct {
	Users  []model.UserDetail
	Groups []model.Group
}

// UserProfileViewHandler 个人空间页面处理器
func UserProfileViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("UserProfileViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/user/profile.html")
	if err != nil {
		panic("parse files failed")
	}

	view := UserProfileView{}

	t.Execute(w, view)
}
