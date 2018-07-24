package common

import (
	"muidea.com/magicCommon/model"
)

// CASHandler 鉴权处理Handler
type CASHandler interface {
	//@in account 账号
	//@in password 密码
	//@in remoteAddr 登陆地址
	//@ret model.OnlineEntryView 登陆用户信息
	//@ret bool 是否登陆成功
	LoginAccount(account, password, remoteAddr string) (model.OnlineEntryView, string, bool)

	// 校验Endpoint
	LoginEndpoint(identifyID, authToken, remoteAddr string) (model.OnlineEntryView, string, bool)

	//@in authToken 鉴权token
	//@ret bool 是否登出成功
	Logout(authToken, remoteAddr string) bool

	// 校验权限是否OK
	VerifyToken(authToken string) (model.OnlineEntryView, string, bool)

	// 刷新Token
	RefreshToken(authToken string, remoteAddr string) bool

	// 查询Cas信息摘要
	GetSummary() model.CasSummary
}
