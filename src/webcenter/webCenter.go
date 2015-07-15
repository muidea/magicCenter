package main
 
import (
    "net/http"
    "log"
    "webcenter/controller/webui"
    "webcenter/controller/rpc"
    "webcenter/controller/session"
    "webcenter/model/datamanager"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0");
	
	datamanager.InitDataManager()
	
	defer datamanager.UninitDataManager()
	
	session.Initialize()
	defer session.Uninitialize()
	
	webui.InitRoute()
	
	rpc.InitRoute()
		
    http.ListenAndServe(":8888", nil)
 
}
