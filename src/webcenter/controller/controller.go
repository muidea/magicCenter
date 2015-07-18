package controller

import (
	"webcenter/controller/session"
	"webcenter/controller/webui"
	"webcenter/controller/patrol"
	"webcenter/controller/rpc"
)

func Initialize() {
	session.Initialize()

	webui.InitRoute()
	
	patrol.InitRoute()
	
	rpc.InitRoute()
			
}

func Uninitialized() {
	
	session.Uninitialize()	
}
