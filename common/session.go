package common

import (
	"net/http"

	"github.com/muidea/magicCommon/model"
)

// Session 会话
type Session interface {
	ID() string

	GetOption(key string) (interface{}, bool)
	SetOption(key string, value interface{})
	RemoveOption(key string)

	// 获取当前登陆账号
	GetAccount() (model.User, bool)
	SetAccount(user model.User)
	ClearAccount()

	OptionKey() []string

	Flush()
}

// SessionRegistry 会话仓库
type SessionRegistry interface {
	GetSession(w http.ResponseWriter, r *http.Request) Session
	UpdateSession(session Session) bool
	FlushSession(session Session)
}
