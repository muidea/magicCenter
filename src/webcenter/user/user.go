package user

import (
	"net/http"
	"reflect"
	"webcenter/session"
)

func init() {
	registerRouter()
}

func registerRouter() {
    http.HandleFunc("/user/verify/",verifyUserHandler)
}

func verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	session := session.GetSession(w,r)
	
    action := "VerifyAction" 
    controller := &userController{}
    verifyController := reflect.ValueOf(controller)
    method := verifyController.MethodByName(action)
    
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    sessionValue := reflect.ValueOf(session)
    method.Call([]reflect.Value{responseValue, requestValue, sessionValue})
}


