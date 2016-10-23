package handler

import (
	"encoding/json"
	"html/template"
	"magiccenter/common"
	"magiccenter/configuration"
	"net/http"
	"path"

	utilPath "muidea.com/path"
)

// Result 处理结果
type Result struct {
	Result common.Result
}

// NoSupportActionHandler 不支持
func NoSupportActionHandler(w http.ResponseWriter, r *http.Request) {
	result := Result{}
	result.Result.ErrCode = 1
	result.Result.Reason = "操作不支持"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// HTMLViewHandler 静态文件处理器
func HTMLViewHandler(w http.ResponseWriter, r *http.Request, htmlfile string) {
	w.Header().Set("content-type", "text/html")
	w.Header().Set("charset", "utf-8")

	// 如果指定的页面不存在，这里返回404页面
	staticPath, _ := configuration.GetOption(configuration.StaticPath)
	fullPath := path.Join(staticPath, htmlfile)
	if !utilPath.PathExist(fullPath) {
		fullPath = path.Join(staticPath, "404.html")
	}

	t, err := template.ParseFiles(fullPath)
	if err != nil {
		panic("parse files failed")
	}

	t.Execute(w, nil)
}
