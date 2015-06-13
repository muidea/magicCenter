package webui

import (
    "log"
    "net/http"
    "html/template"
)

type User struct {
    UserName string
}

type adminController struct {
}

func (this *adminController)AdminAction(w http.ResponseWriter, r *http.Request, user string) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")	
    t, err := template.ParseFiles("template/html/admin.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, &User{user})
}