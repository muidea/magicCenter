package blog

import (
	"log"
	"net/http"
	"html/template"
	"magiccenter/kernel/module/model"
	"magiccenter/kernel/module/bll"
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

func viewContentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("viewContentHandler")
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
    t, err := template.ParseFiles("template/html/blog/view.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(w, nil)
}


func viewCatalogHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("viewCatalogHandler")
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
    t, err := template.ParseFiles("template/html/blog/catalog.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(w, nil)
}

func viewLinkHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("viewLinkHandler")
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
    t, err := template.ParseFiles("template/html/blog/view.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(w, nil)
}


