package blog


import (
	"webcenter/application"
)

func registerRouter() {
    application.RegisterGetHandler("/blog/article/",viewArticleHandler)
    application.RegisterGetHandler("/blog/",indexHandler)
}

