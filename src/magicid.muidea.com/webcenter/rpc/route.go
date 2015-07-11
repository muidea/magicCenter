package rpc

import (
    "net/http"
	"strings"
	"reflect"
	"log"    
)

func InitRoute() {
    http.HandleFunc("/rpc/", rpcHandler)	
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
    pathInfo := strings.Trim(r.URL.Path, "/")
    log.Print(pathInfo)
    parts := strings.Split(pathInfo, "/")
    var action = ""
    if len(parts) > 0 {
        action = strings.Title(parts[0]) + "Action"        
    } else {
    	log.Print("illegal path");
    }

	//log.Panicf("action:%s", action);    
    rpc := &rpcController{}
    controller := reflect.ValueOf(rpc)
    method := controller.MethodByName(action)
    if !method.IsValid() {
        method = controller.MethodByName(strings.Title("rpc") + "Action")
    }
    
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
}

