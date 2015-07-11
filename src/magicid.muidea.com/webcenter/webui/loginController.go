package webui

import (
    "net/http"
    "html/template"
    "log"
    "magicid.muidea.com/webcenter/session"
)
 
type loginController struct {
}

func (this *loginController)LoginAction(w http.ResponseWriter, r *http.Request, session *session.Session) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	
	access_token := session.AccessToken()
    t, err := template.ParseFiles("template/html/login.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, access_token)
}
