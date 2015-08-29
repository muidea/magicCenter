package admin

import (
    "webcenter/session"
)

type AdminView struct {
	Accesscode string
	NickName string
}

type adminController struct {
}
 
func (this *adminController)Action(session *session.Session) AdminView {
	view := AdminView{}
	
	return view
}
