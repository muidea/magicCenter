package webui

import (
	"net/http"
	"html/template"
	"log"
	"strings"
	"reflect"
)

func InitRoute() {
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
     
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/login/",loginHandler)
    http.HandleFunc("/register/", registerHandler)
    http.HandleFunc("/ajax/",ajaxHandler)
    http.HandleFunc("/",notFoundHandler)	
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    // 获取cookie
    cookie, err := r.Cookie("loginAccount")
    if err != nil || cookie.Value == ""{
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
    userValue := reflect.ValueOf(cookie.Value)
    method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
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
    method.Call([]reflect.Value{responseValue, requestValue})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template/html/register.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)	
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
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
    method.Call([]reflect.Value{responseValue, requestValue})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        http.Redirect(w, r, "/login/", http.StatusFound)
    }
    
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	     
    t, err := template.ParseFiles("template/html/404.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}


