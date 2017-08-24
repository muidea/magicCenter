package common

import (
	"net/http"
)

// CorsHandler Cors处理器
type CorsHandler interface {
	CheckCors(res http.ResponseWriter, req *http.Request) bool
}
