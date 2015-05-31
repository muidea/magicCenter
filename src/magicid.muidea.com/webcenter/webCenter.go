package main
 
import (
    "net/http"
    "log"
    "magicid.muidea.com/webcenter/webui"
    "magicid.muidea.com/webcenter/rpc"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0");
	
	webui.InitRoute()
	
	rpc.InitRoute()
	
    http.ListenAndServe(":8888", nil)
 
}
