package admin

import (
	"net/http"
	"html/template"
	"log"
	"webcenter/session"
)

type AdminView struct {
	Accesscode string
	NickName string
}


func AdminHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("adminHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	param := &ManageParam{}	
	session := session.GetSession(w,r)
	param.session = session
	
	controller := &manageController{}
	
	result := controller.ManageAction(param)
	if result.Fail() {
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return		
	}
		
    t, err := template.ParseFiles("template/html/admin/index.html")
    if (err != nil) {    	
    	panic("parse file failed")
    }
    
    view := AdminView{}
    view.Accesscode = session.AccessToken()
    view.NickName = result.user.NickName
    
    t.Execute(w, view)	
}
