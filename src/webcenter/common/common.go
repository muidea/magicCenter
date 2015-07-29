package common

import (
	"net/http"
	"html/template"
	"log"
	"webcenter/session"
)

func init() {
	registerRouter()
}


func registerRouter() {
    http.Handle("/resources/css/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/scripts/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/images/", http.FileServer(http.Dir("template")))
     
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/login/",loginHandler)
    http.HandleFunc("/logout/",logoutHandler)
    http.HandleFunc("/register/", registerHandler)
    http.HandleFunc("/",notFoundHandler)		
}


func adminHandler(w http.ResponseWriter, r *http.Request) {
	session := session.GetSession(w,r)
	_, found := session.GetOption("account")		
    if !found {
        http.Redirect(w, r, "/login/", http.StatusFound)
        return
    }
    
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")	
	access_token := session.AccessToken()
	account, _ := session.GetOption("account")
	
    t, err := template.ParseFiles("template/html/admin.html")
    if (err != nil) {
    	log.Println(err)
        http.Redirect(w, r, "/", http.StatusNotFound)
        return
    }
 
 	pageInfo := AdminPage{Account:account.(string), AccessToken:access_token}    
    t.Execute(w, pageInfo)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session := session.GetSession(w,r)
	
	_, found := session.GetOption("account")
	if found {
        http.Redirect(w, r, "/admin/", http.StatusFound)
        return		
	}
	
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")	
	access_token := session.AccessToken()
    t, err := template.ParseFiles("template/html/login.html")
    if (err != nil) {
        log.Println(err)
        http.Redirect(w, r, "/", http.StatusNotFound)
        return
    }
    t.Execute(w, access_token)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session := session.GetSession(w,r)
	
	_, found := session.GetOption("account")
	if !found {
        http.Redirect(w, r, "/", http.StatusFound)
        return
	}

	session.RemoveOption("account")
	http.Redirect(w, r, "/", http.StatusFound)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template/html/register.html")
    if (err != nil) {
        log.Println(err)
        return
    }
    t.Execute(w, nil)	
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
    if r.URL.Path == "/" {
	    t, err := template.ParseFiles("template/html/index.html")
    	if (err != nil) {
        	log.Println(err)
        	return
    	}
	    t.Execute(w, nil)
	    return;
    }
    
    t, err := template.ParseFiles("template/html/404.html")
    if (err != nil) {
        log.Println(err)
        return
    }
    t.Execute(w, nil)
}