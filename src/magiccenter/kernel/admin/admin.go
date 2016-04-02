package admin

import (
	"net/http"
	"html/template"
	"log"
	"magiccenter/router"
	"magiccenter/kernel/account/model"
)

type AdminView struct {
	User model.User
}

func RegisterRouter() {
	router.AddGetRoute("/admin/", AdminHandler)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("adminHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/index.html")
    if (err != nil) {    	
    	panic("parse file failed")
    }
    
    view := AdminView{}
    view.User.Id = 100
    view.User.Name = "rangh"
    t.Execute(w, view)	
}
