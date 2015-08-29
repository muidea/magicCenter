package admin

import (
	"net/http"
	"html/template"
	"log"
	"webcenter/session"
	"webcenter/auth"
)

func init() {
	registerRouter()
}

func registerRouter() {
	http.HandleFunc("/admin/", adminHandler)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("adminHandler");
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)
	account, found := session.GetOption(auth.AccountSessionKey)
	if !found {
		log.Print("can't get account")
		
		http.Redirect(w, r, "/auth/login/", http.StatusFound)
		return
	}

	userModel, err := auth.NewModel()
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}
	defer userModel.Release()
	
	user, found := userModel.FindUserByAccount(account.(string))
	if !found || !user.IsAdmin() {
		http.Redirect(w, r, "/", http.StatusFound)
		return		
	}
	
    t, err := template.ParseFiles("template/html/admin/index.html")
    if (err != nil) {
        log.Fatal(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
 
    controller := &adminController{}
    view := controller.Action(session)
    view.Accesscode = session.AccessToken()
    view.NickName = user.NickName
    
    t.Execute(w, view)	
}
