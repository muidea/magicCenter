package webui

import (
    "net/http"
    "html/template"
    "log"
)
 
type loginController struct {
}

func (this *loginController)LoginAction(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("template/html/login.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, nil)
}
