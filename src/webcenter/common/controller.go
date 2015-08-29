package common

import (
	"webcenter/session"
)



type AdminView struct {
    Account string
    AccessToken string
}

type Result struct {
	ErrCode int
	Reason string
}

type Controller struct {
	session *session.Session
}

