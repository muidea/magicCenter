package service

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
)

// SessionRegistry 会话仓库
type SessionRegistry interface {
	GetSession(w http.ResponseWriter, r *http.Request) common.Session
	UpdateSession(session common.Session) bool
}
