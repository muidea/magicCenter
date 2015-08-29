package auth

import (
	"net/http"
	"html/template"
	"encoding/json"
	"log"
	"webcenter/session"
)

func init() {
	registerRouter()
}

func registerRouter() {
	http.HandleFunc("/auth/login/",loginHandler)
	http.HandleFunc("/auth/logout/",logoutHandler)
    http.HandleFunc("/auth/verify/",verifyHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("loginHandler");
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	session := session.GetSession(w,r)	
    t, err := template.ParseFiles("template/html/auth/login.html")
    if (err != nil) {
        log.Fatal(err)
        
        http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    _, found := session.GetOption(AccountSessionKey)
    if found {
        http.Redirect(w, r, "/admin/", http.StatusFound)
        return
    }
 
    controller := &loginController{}
    view := controller.Action(session)
        
    view.Accesscode = session.AccessToken()
    
    t.Execute(w, view)	
}


func logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("logoutHandler");
	
	session := session.GetSession(w,r)
	session.RemoveOption(AccountSessionKey)
	
    http.Redirect(w, r, "/", http.StatusFound)
}


func verifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("verifyHandler");
    err := r.ParseForm()
    if err != nil {
    	log.Fatal("paseform failed")
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }

	param := VerifyParam{}
	param.account = r.FormValue("login_account")
	param.password = r.FormValue("login_password")
	param.accesscode = r.FormValue("accesscode")
	
	log.Printf("account:%s,password:%s,accesscode:%s", param.account, param.password, param.accesscode)
		
	session := session.GetSession(w,r)
	
    controller := &verifyController{}
    result := controller.Action(&param, session)
    b, err := json.Marshal(result)
    if err != nil {
    	log.Fatal("json marshal failed, err:" + err.Error())
    	http.Redirect(w, r, "/404/", http.StatusNotFound)
        return
    }
    
    w.Write(b)
}


