package handler

import (
	"html/template"
	"net/http"

	"path"

	"os"

	"muidea.com/magicCenter/application/common"
)

// CreateStaticHandler 新建StaticHandler
func CreateStaticHandler(rootPath string) common.StaticHandler {
	i := impl{rootPath: rootPath}

	return &i
}

type impl struct {
	rootPath string
}

func (i *impl) HandleView(basePath string, w http.ResponseWriter, r *http.Request) {
	fullPath := path.Join(i.rootPath, basePath, r.URL.Path)
	_, err := os.Stat(fullPath)
	if err == nil || os.IsExist(err) {
	} else {
		fullPath = path.Join(i.rootPath, "404.html")
	}

	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")
	t, err := template.ParseFiles(fullPath)
	if err != nil {
		panic("parse files failed")
	}

	t.Execute(w, nil)
}
