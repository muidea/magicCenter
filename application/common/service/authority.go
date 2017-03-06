package service

import "net/http"

// Authority 鉴权
type Authority interface {
	Verify(res http.ResponseWriter, req *http.Request) bool
}
