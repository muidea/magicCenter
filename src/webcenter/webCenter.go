package main
 
import (
    "net/http"
    "log"
    _"webcenter/datamanager"
    _"webcenter/session"
    _"webcenter/common"
    _"webcenter/user"
    _"webcenter/patrol"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0");
		
    http.ListenAndServe(":8888", nil)
 
}
