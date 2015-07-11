package webui

import (
    "log"
    "net/http"
    "html/template"
    "magicid.muidea.com/webcenter/session"
)

type adminPage struct {
    Account string
    AccessToken string
}

type adminController struct {
}

func (this *adminController)AdminAction(w http.ResponseWriter, r *http.Request, session *session.Session) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")	
	access_token := session.AccessToken()
	account, _ := session.GetOption("account")
	
    t, err := template.ParseFiles("template/html/admin.html")
    if (err != nil) {
        log.Println(err)
    }
 
 	pageInfo := adminPage{Account:account.(string), AccessToken:access_token}
    
    t.Execute(w, pageInfo)
}