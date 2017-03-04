package common

import "github.com/go-martini/martini"

// Authority 鉴权
type Authority interface {
	AdminAuthVerify() martini.Handler

	LoginAuthVerify() martini.Handler

	Authority() martini.Handler
}
