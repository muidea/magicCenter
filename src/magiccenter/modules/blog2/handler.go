package blog2

/*
import (
	"net/http"
	"html/template"
	"strings"
	"strconv"
	"log"
)

func viewArticleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/blog/view.html")
    if (err != nil) {
        log.Print(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
	var id = -1
	idInfo := r.URL.RawQuery
	if len(idInfo) > 0 {
		parts := strings.Split(idInfo,"=")
		if len(parts) == 2 {
			id, err = strconv.Atoi(parts[1])
			if err != nil {
				http.Redirect(w, r, "/404/", http.StatusNotFound)
			}
		}
	}
		    
    controller := &uiController{}
    view := controller.ViewArticleAction(id)
        
    t.Execute(w, view)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/blog/index.html")
    if (err != nil) {
        log.Print(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    controller := &uiController{}
    view := controller.IndexAction()
    
    t.Execute(w, view)
}
*/


