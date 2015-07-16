package main
 
import (
    "net/http"
    "log"
    "webcenter/controller"
    "webcenter/model"
)
 
func main() {
	log.Println("MagicID WebCenter V1.0");
	
	controller.Initialize()
	
	defer controller.Uninitialized();
	
	model.Initialize()
	
	defer model.Uninitialized();
	
    http.ListenAndServe(":8888", nil)
 
}
