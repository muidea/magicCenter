package user
 
import (
    "net/http"
    "encoding/json"
    "webcenter/common"
    "webcenter/session"
)
 
type VerifyResult struct{
	common.Result
    RedirectUrl string
}
 
type userController struct {
}
 
func (this *userController)VerifyAction(w http.ResponseWriter, r *http.Request, session *session.Session) {	
    w.Header().Set("content-type", "application/json")
    w.Header().Set("charset", "utf-8")
    
	var errCode int
	var reason string
	var redirectUrl string
		    
    err := r.ParseForm()
    if err != nil {
    	errCode = 1
    	reason = "非法输入"
        this.OutputLoginResult(w, errCode, reason, redirectUrl)
        return
    }
     
    loginAccount := r.FormValue("login_account")
    loginPassword := r.FormValue("login_password")
    if loginAccount == "" || loginPassword == "" {
    	errCode = 2
    	reason = "输入信息错误"
        this.OutputLoginResult(w, errCode, reason, redirectUrl)
        return
    }
    
    user, found:= FindUserByEmail(loginAccount)
    if !found || !user.ValidPassword(loginPassword) {
    	errCode = 2
    	reason = "非法账号"
        this.OutputLoginResult(w, errCode, reason, redirectUrl)
        return
    }
    
    // 存入cookie,使用cookie存储
    cookie := http.Cookie{Name: "session_id", Value: session.Id(), Path: "/"}
    http.SetCookie(w, &cookie)

	session.SetOption("account", loginAccount)

	errCode = 0;
	redirectUrl = "/admin/"
    this.OutputLoginResult(w, errCode, reason, redirectUrl)
    return
}
 
func (this *userController)OutputLoginResult(w http.ResponseWriter, errCode int, reason string, redirect string ) {
    out := &VerifyResult{}
    out.ErrCode = errCode
    out.Reason = reason
    out.RedirectUrl = redirect
    
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}

