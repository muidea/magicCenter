package webui
 
import (
    "net/http"
    "encoding/json"
    "magicid.muidea.com/webcenter/datamanager"
    "magicid.muidea.com/webcenter/session"
)
 
type Result struct{
    Ret int
    Reason string
    Data string
}
 
type ajaxController struct {
}
 
func (this *ajaxController)AjaxAction(w http.ResponseWriter, r *http.Request, session *session.Session) {	
    w.Header().Set("content-type", "application/json")
    w.Header().Set("charset", "utf-8")
    err := r.ParseForm()
    if err != nil {
        OutputJson(w, 1, "invalid Param", "")
        return
    }
     
    loginAccount := r.FormValue("login_account")
    loginPassword := r.FormValue("login_password")
    accessToken := r.FormValue("access_token")
    if session.ValidToken(accessToken) {
    	return
    }

    if loginAccount == "" || loginPassword == "" {
        OutputJson(w, 1, "invalid Param", "")
        return
    }
    
    userManager := datamanager.GetUserManager()
    user, found:= userManager.FindUserByEMail(loginAccount)
    if !found || !user.ValidPassword(loginPassword) {
    	OutputJson(w, 1, "invalid account or password", "")
    	return
    }
    
    // 存入cookie,使用cookie存储
    cookie := http.Cookie{Name: "session_id", Value: session.Id(), Path: "/"}
    http.SetCookie(w, &cookie)

	session.SetOption("account", loginAccount)

    OutputJson(w, 0, "Login OK", "/admin/")
    return
}
 
func OutputJson(w http.ResponseWriter, ret int, reason string, redirect string ) {
    out := &Result{ret, reason, redirect}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

