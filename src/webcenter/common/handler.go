package common

import (
	"net/http"
)

func init() {
	registerRouter()
}

func registerRouter() {
    http.Handle("/resources/css/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/scripts/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/images/", http.FileServer(http.Dir("template")))
    http.Handle("/resources/upload/", http.FileServer(http.Dir("template")))
}
