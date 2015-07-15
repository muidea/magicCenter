package rpc

import (
	"log"
	"encoding/json"
	"net/http"
	"webcenter/model/datamanager"
)

type rpcController struct {
}

type Result struct {
	Route []datamanager.Routeline
}

func (this *rpcController) RpcAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("charset", "utf-8")

	routelineManager := datamanager.GetRoutelineManager()
	routelist := routelineManager.GetAll()
	
	log.Printf("routelist size:%d", len(routelist))
	result := Result{Route:routelist}	
	b, err := json.Marshal(&result)
    if err != nil {
        return
    }
    w.Write(b)
}
