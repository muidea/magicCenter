package webui
 
import (
    "net/http"
    "encoding/json"
    "magicid.muidea.com/webcenter/datamanager"
)
 
type Result struct{
    Ret int
    Reason string
    Data interface{}
}
 
type ajaxController struct {
}
 
func (this *ajaxController)AjaxAction(w http.ResponseWriter, r *http.Request) {	
    w.Header().Set("content-type", "application/json")
    w.Header().Set("charset", "utf-8")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 0, "invalid Param", nil)
        return
    }
     
    loginAccount := r.FormValue("login_account")
    loginPassword := r.FormValue("login_password")
     
    if loginAccount == "" || loginPassword == "" {
        OutputJson(w, 0, "invalid Param", nil)
        return
    }
    
    userManager := datamanager.GetUserManager()
    user, found:= userManager.FindUserByEMail(loginAccount)
    if !found || !user.ValidPassword(loginPassword) {
    	OutputJson(w, 0, "invalid account or password", nil)
    	return
    }
    
    // 存入cookie,使用cookie存储
    cookie := http.Cookie{Name: "loginAccount", Value: loginAccount, Path: "/"}
    http.SetCookie(w, &cookie)
     
    OutputJson(w, 1, "Login OK", nil)
    return
}
 
func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

