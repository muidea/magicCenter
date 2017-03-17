package common

import (
	"muidea.com/magicCenter/application/common/model"
)

// AuthorityHandler 鉴权处理Handler
type AuthorityHandler interface {
	//@in account 账号
	//@in password 密码
	//@ret model.UserDetail 登陆用户
	//@ret string 本次登陆的鉴权token
	//@ret bool 是否登陆成功
	LoginAccount(account, password string) (model.UserDetail, string, bool)

	//@in authID 鉴权token
	//@ret bool 是否登出成功
	LogoutAccount(authID string) bool

	//@in authID 鉴权token
	IsLogin(authID string) bool
}
