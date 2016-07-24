package blog

import (
	"html/template"
	"log"
	"magiccenter/kernel/dashboard/module/bll"
	"magiccenter/kernel/dashboard/module/model"
	contentbll "magiccenter/kernel/modules/content/bll"
	"net/http"
	"strconv"

	"muidea.com/util"
)

type PageView struct {
	View model.PageView
}

type PageContentView struct {
	View model.PageContentView
}

type PageCatalogView struct {
	View model.PageCatalogView
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("indexHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	view := PageView{}

	url := req.URL.Path
	view.View, _ = bll.QueryPageView(ID, url)

	t, err := template.ParseFiles("template/html/modules/blog/index.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

func viewContentHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewContentHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	params := util.SplitParam(req.URL.RawQuery)
	str, found := params["id"]
	if !found {
		panic("illegl param")
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		panic("illegl id, err:" + err.Error())
	}

	view := PageContentView{}

	url := req.URL.Path
	view.View, _ = bll.QueryContentView(ID, url, id)
	t, err := template.ParseFiles("template/html/modules/blog/view.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

func viewCatalogHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewCatalogHandler")

	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")

	params := util.SplitParam(req.URL.RawQuery)
	str, found := params["id"]
	if !found {
		panic("illegl param")
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		panic("illegl id, err:" + err.Error())
	}

	view := PageCatalogView{}

	url := req.URL.Path
	view.View, _ = bll.QueryCatalogView(ID, url, id)

	t, err := template.ParseFiles("template/html/modules/blog/catalog.html")
	if err != nil {
		panic("ParseFiles failed, err:" + err.Error())
	}

	t.Execute(res, view)
}

func viewLinkHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewLinkHandler")

	params := util.SplitParam(req.URL.RawQuery)
	str, found := params["id"]
	if !found {
		panic("illegl param")
	}

	id, err := strconv.Atoi(str)
	if err != nil {
		panic("illegl id, err:" + err.Error())
	}

	link, found := contentbll.QueryLinkById(id)
	if !found {
		http.Redirect(res, req, "/", http.StatusNotFound)
		return
	}

	http.Redirect(res, req, link.Url, http.StatusFound)
}
