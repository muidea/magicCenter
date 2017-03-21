package common

import "net/http"

// StaticHandler 静态页处理器
type StaticHandler interface {
	HandleView(basePath string, w http.ResponseWriter, r *http.Request)
	HandleResource(basePath string, w http.ResponseWriter, r *http.Request)
}
