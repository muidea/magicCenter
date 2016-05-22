package blog

import (
	"log"
	"net/http"
	"html/template"
	"strconv"
	"muidea.com/util"
	"magiccenter/kernel/module/model"
	"magiccenter/kernel/module/bll"
	contentbll "magiccenter/kernel/content/bll"
)

type PageView struct {
	View model.PageView
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("indexHandler")
	
	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")
	
	view := PageView{}
		
	url := req.URL.Path	
	view.View, _ = bll.QueryPageView(ID, url)
	
    t, err := template.ParseFiles("template/html/blog/index.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(res, view)
}

func viewContentHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewContentHandler")
	
	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")
	
	/*
	params := util.SplitParam(req.URL.RawQuery)
	str, found := params["id"]
	if !found {
		panic("illegl param")
	}
	
	id, err := strconv.Atoi(str)
	if err!= nil {
		panic("illegl id, err:" + err.Error())
	}*/
		
	view := PageView{}
		
	url := req.URL.Path	
	view.View, _ = bll.QueryPageView(ID, url)
	
    t, err := template.ParseFiles("template/html/blog/view.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(res, view)
}


func viewCatalogHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("viewCatalogHandler")
	
	res.Header().Set("content-type", "text/html")
	res.Header().Set("charset", "utf-8")
	
	view := PageView{}
		
	url := req.URL.Path	
	view.View, _ = bll.QueryPageView(ID, url)
	
    t, err := template.ParseFiles("template/html/blog/catalog.html")
    if (err != nil) {
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
	if err!= nil {
		panic("illegl id, err:" + err.Error())
	}
		
	link, found := contentbll.QueryLinkById(id);
	if !found {
		http.Redirect(res, req, "/", http.StatusNotFound)
		return
	}
	
	http.Redirect(res, req, link.Url, http.StatusFound)
}


