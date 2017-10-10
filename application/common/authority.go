package common

import (
	"net/http"
)

// AuthorityHandler 鉴权处理器
type AuthorityHandler interface {
	VerifyAuthority(res http.ResponseWriter, req *http.Request) bool
}
