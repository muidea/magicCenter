package blog

import (
	"log"
	"net/http"
	"html/template"
)

func viewArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("viewArticleHandler")
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
    t, err := template.ParseFiles("template/html/blog/view.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("indexHandler")
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/blog/index.html")
    if (err != nil) {
    	panic("ParseFiles failed, err:" + err.Error())
    }
    
    t.Execute(w, nil)
}
