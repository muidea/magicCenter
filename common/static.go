package common

import "net/http"

// StaticHandler 静态页处理器
type StaticHandler interface {
	HandleResource(w http.ResponseWriter, r *http.Request)
}
