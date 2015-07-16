package webui

import (
    "log"
    "net/http"
    "html/template"
    "webcenter/model"
    "webcenter/controller/session"
)

type adminPage struct {
    Account string
    AccessToken string
}

type adminPatrolPage struct {
	adminPage
	Routeline []model.Routeline
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


func (this *adminController)AdminPatrolAction(w http.ResponseWriter, r *http.Request, session *session.Session) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")	
	access_token := session.AccessToken()
	account, _ := session.GetOption("account")
	
    t, err := template.ParseFiles("template/html/admin_patrol.html")
    if (err != nil) {
        log.Println(err)
    }
 
 	routeline := model.GetRouteLine()
 	pageInfo := adminPatrolPage{}
 	pageInfo.Account = account.(string)
 	pageInfo.AccessToken = access_token
 	pageInfo.Routeline = routeline
    
    t.Execute(w, pageInfo)
}