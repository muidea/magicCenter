package common

import "martini"

// Authority 鉴权
type Authority interface {
	AdminAuthVerify() martini.Handler

	LoginAuthVerify() martini.Handler

	Authority() martini.Handler
}
