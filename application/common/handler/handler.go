package handler

/*
提供基础的Handler实现
1、操作不支持
2、静态页面访问
*/

import (
	"encoding/json"
	"html/template"
	"net/http"

	"muidea.com/magicCenter/application/common"
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

	htmlFile := system.GetHTMLPath(htmlfile)
	t, err := template.ParseFiles(htmlFile)
	if err != nil {
		panic("parse files failed")
	}

	t.Execute(w, nil)
}
