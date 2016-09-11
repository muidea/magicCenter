package ui

import (
	"html/template"
	"log"
	"magiccenter/configuration"
	"net/http"
)

const passwordMark = "********"

// SystemSettingView 系统设置视图
type SystemSettingView struct {
	SystemInfo configuration.SystemInfo
}

// SystemSettingViewHandler 系统设置视图处理器
func SystemSettingViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("SystemSettingViewHandler")

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	t, err := template.ParseFiles("template/html/admin/system/system.html")
	if err != nil {
		panic("parse files failed")
	}

	view := SystemSettingView{}
	view.SystemInfo = configuration.GetSystemInfo()
	view.SystemInfo.MailPassword = passwordMark

	t.Execute(w, view)
}
