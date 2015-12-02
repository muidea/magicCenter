package admin

import (
	"webcenter/application"
)

func init() {
	registerRouter()
}

func registerRouter() {
	application.RegisterGetHandler("/admin/", AdminHandler)
}

