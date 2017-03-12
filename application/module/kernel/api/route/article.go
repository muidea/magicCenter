package route

import (
	"net/http"

	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/kernel/modulehub"
)

// CreateGetArticleRoute 新建GetArticle Route
func CreateGetArticleRoute(modHub modulehub.ModuleHub) common.Route {
	i := &getArticleRoute{moduleHub: modHub}

	return i
}

type getArticleRoute struct {
	moduleHub modulehub.ModuleHub
}

func (i *getArticleRoute) Type() string {
	return common.GET
}

func (i *getArticleRoute) Pattern() string {
	return "content/article/[0-9]*/"
}

func (i *getArticleRoute) Handler() interface{} {
	return i.getArticleHandler
}

func (i *getArticleRoute) getArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getArticleHandler")

	//id, _, ok := net.ParseRestAPIUrl(r.URL.Path)
}
