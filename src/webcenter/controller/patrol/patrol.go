package patrol

import (
    "net/http"
	"strings"
	"reflect"
)

func InitRoute() {
    http.HandleFunc("/patrol/patrolline/", patrolLineHandler)	
}

func patrolLineHandler(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	pathInfo = strings.ToLower(pathInfo)
	parts := strings.Split(pathInfo, "/")
	
	var action string
	if len(parts) == 0 {
		action = "PatrolLineAction"
	} else {
		action = strings.Title(parts[len(parts) -1]) + "Action"
	}
	
    patrolLineController := &patrolLineController{}
    controller := reflect.ValueOf(patrolLineController)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName("PatrolLineAction")
    }
    
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

