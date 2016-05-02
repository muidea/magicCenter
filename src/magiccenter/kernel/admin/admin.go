package admin

import (
	"net/http"
	"html/template"
	"log"
	"magiccenter/configuration"
	"magiccenter/session"
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
	
	authId, found := configuration.GetOption(configuration.AUTHORITH_ID)
	if !found {
		panic("unexpected, can't fetch authorith id")
	}
	
	session := session.GetSession(w, r)
	user, found := session.GetOption(authId)
	if !found {
		panic("unexpected, must login system first.")
	}
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
    t, err := template.ParseFiles("template/html/admin/index.html")
    if (err != nil) {    	
    	panic("parse file failed")
    }
    
    view := AdminView{}
    view.User.Id = user.(model.UserDetailView).Id
    view.User.Name = user.(model.UserDetailView).Name
    t.Execute(w, view)	
}
