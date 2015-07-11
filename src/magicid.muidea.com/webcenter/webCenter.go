package main
 
import (
    "net/http"
    "log"
    "magicid.muidea.com/webcenter/webui"
    "magicid.muidea.com/webcenter/rpc"
    "magicid.muidea.com/webcenter/datamanager"
    "magicid.muidea.com/webcenter/session"
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
