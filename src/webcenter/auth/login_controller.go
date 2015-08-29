package auth

import (
    "webcenter/session"    
)

type LoginView struct {
	Accesscode string
}

type loginController struct {
}
 
func (this *loginController)Action(session *session.Session) LoginView {
	view := LoginView{}
	
	return view
}
