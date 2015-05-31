package rpc

import (
    "net/http"
)

func InitRoute() {
    http.HandleFunc("/rpc/", rpcHandler)	
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
}

