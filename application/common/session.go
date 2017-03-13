package common

import (
	"net/http"

	"muidea.com/magicCenter/application/common/model"
)

// Session 会话
type Session interface {
	ID() string

	GetOption(key string) (interface{}, bool)
	SetOption(key string, value interface{})
	RemoveOption(key string)

	// 获取当前登陆账号
	GetAccount() (model.UserDetail, bool)
	SetAccount(user model.UserDetail)
	ClearAccount()

	OptionKey() []string
}

// SessionRegistry 会话仓库
type SessionRegistry interface {
	GetSession(w http.ResponseWriter, r *http.Request) Session
	UpdateSession(session Session) bool
}
