package ui

import (
	"webcenter/application"
)

func init() {
	registerRouter()
}

func registerRouter() {
    application.RegisterGetHandler("/view/article/",viewArticleHandler)
    application.RegisterGetHandler("/",indexHandler)
}




