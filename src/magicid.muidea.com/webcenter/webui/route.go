package webui

import (
	"net/http"
	"html/template"
	"log"
	"strings"
	"reflect"
	"magicid.muidea.com/webcenter/session"
)

func InitRoute() {
    http.Handle("/resources/css/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/scripts/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/images/", http.FileServer(http.Dir("template")))
     
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/login/",loginHandler)
    http.HandleFunc("/logout/",logoutHandler)
    http.HandleFunc("/register/", registerHandler)
    http.HandleFunc("/ajax/",ajaxHandler)
    http.HandleFunc("/",notFoundHandler)	
}

func getSession(w http.ResponseWriter, r *http.Request) *session.Session {
	var userSession *session.Session
	
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == ""{
		userSession = session.CreateSession()
		log.Printf("can't find cookie,create new session")
	} else {
		cur, found := session.SessionManger().Find(cookie.Value)
		if !found {
			userSession = session.CreateSession()
			log.Printf("invalid cookie,create new session, cookieValue:%s", cookie.Value)
		} else {
			userSession = cur
			log.Print("find exist ession from cookie")
		}
	}
	
    // 存入cookie,使用cookie存储
    session_cookie := http.Cookie{Name: "session_id", Value: userSession.Id(), Path: "/magic.muidea.com/webcenter"}
    http.SetCookie(w, &session_cookie)
	
	return userSession
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(w,r)
	
	_, found := session.GetOption("account")		
    if !found {
        http.Redirect(w, r, "/login/", http.StatusFound)
        return
    }
    
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 0 {
        action = strings.Title(parts[0]) + "Action"
    }
    
    admin := &adminController{}
    controller := reflect.ValueOf(admin)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("admin") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    sessionValue := reflect.ValueOf(session)
    method.Call([]reflect.Value{responseValue, requestValue, sessionValue})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(w,r)
	
	_, found := session.GetOption("account")
	if found {
        http.Redirect(w, r, "/admin/", http.StatusFound)
        return		
	}
	
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 0 {
        action = strings.Title(parts[0]) + "Action"
    }
 
    login := &loginController{}
    controller := reflect.ValueOf(login)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("login") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    sessionValue := reflect.ValueOf(session)
    method.Call([]reflect.Value{responseValue, requestValue, sessionValue})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(w,r)
	
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
    }
    t.Execute(w, nil)	
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(w,r)
	
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 0 {
        action = strings.Title(parts[0]) + "Action"
    }
 
    login := &ajaxController{}
    controller := reflect.ValueOf(login)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("ajax") + "Action")
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    sessionValue := reflect.ValueOf(session)
    method.Call([]reflect.Value{responseValue, requestValue, sessionValue})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
    if r.URL.Path == "/" {
        //http.Redirect(w, r, "/login/", http.StatusFound)
	    t, err := template.ParseFiles("template/html/index.html")
    	if (err != nil) {
        	log.Println(err)
    	}
	    t.Execute(w, nil)
        return
    }
    
    t, err := template.ParseFiles("template/html/404.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}


