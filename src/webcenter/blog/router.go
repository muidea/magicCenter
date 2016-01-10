package blog


import (
	"webcenter/application"
)

func init() {
	registerRouter()
}

func registerRouter() {
    application.RegisterGetHandler("/blog/article/",viewArticleHandler)
    application.RegisterGetHandler("/blog/",indexHandler)
}

