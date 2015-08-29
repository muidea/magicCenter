package common

import (
	"net/http"
	"html/template"
	"log"
)

func get404Content() string {
	content := "<html><head><title>404</title></head><body>"
	content += "<p>Author:muidea@gmail.com</p>"
	content += "</body></html>"
	return content
}

func init() {
	registerRouter()
}

func registerRouter() {
    http.Handle("/resources/css/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/scripts/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/images/", http.FileServer(http.Dir("template")))
     
    http.HandleFunc("/404/",notFoundHandler)
    http.HandleFunc("/",indexHandler)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/404.html")
    if (err != nil) {
        log.Println(err)
        //content := get404Content()
        
        //w.Write(content.to)
        return
    }
    
    t.Execute(w, nil)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/index.html")
    if (err != nil) {
        log.Fatal(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    t.Execute(w, nil)
}