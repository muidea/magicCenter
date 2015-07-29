package patrol

import (
    "net/http"
	"strings"
	"reflect"
	"log"
	"webcenter/session"	
)

func init() {
	registerRouter()
}

func registerRouter() {
    http.HandleFunc("/admin/patrol", adminPatrolHandler)	
    http.HandleFunc("/patrol/patrolline/", patrolLineHandler)	
}

func adminPatrolHandler(w http.ResponseWriter, r *http.Request) {
	session := session.GetSession(w,r)
	
	_, found := session.GetOption("account")		
    if !found {
        http.Redirect(w, r, "/login/", http.StatusFound)
        return
    }
    
    action := "AdminPatrolAction"
    patrolController := &patrolController{}
    controller := reflect.ValueOf(patrolController)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        http.Redirect(w, r, "/", http.StatusNotFound)
        return
    }
    
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    sessionValue := reflect.ValueOf(session)
    method.Call([]reflect.Value{responseValue, requestValue, sessionValue})
}


func patrolLineHandler(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	//pathInfo = strings.ToLower(pathInfo)
	parts := strings.Split(pathInfo, "/")
	
	var action string
	if len(parts) <= 2 {
		action = "GetallPatrolLineAction"
	} else {
		action = strings.Title(parts[len(parts) -1]) + "PatrolLineAction"
	}
	
	log.Print(action)
	
    patrolController := &patrolController{}
    controller := reflect.ValueOf(patrolController)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        http.Redirect(w, r, "/", http.StatusInternalServerError)
        return
    }
    
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

