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
    t, err := template.ParseFiles("template/html/admin.html")
    if (err != nil) {
        log.Println(err)
    }
    t.Execute(w, &User{user})
}