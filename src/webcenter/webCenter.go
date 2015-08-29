package main
 
import (
    "net/http"
    "log"
    _"webcenter/session"
    _"webcenter/common"
    _"webcenter/auth"
    _"webcenter/admin"
    _"webcenter/content"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0");
		
    http.ListenAndServe(":8888", nil)
 
}

