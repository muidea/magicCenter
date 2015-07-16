package controller

import (
	"webcenter/controller/session"
	"webcenter/controller/webui"
	"webcenter/controller/rpc"
)

func Initialize() {
	session.Initialize()

	webui.InitRoute()
	
	rpc.InitRoute()
			
}

func Uninitialized() {
	
	session.Uninitialize()	
}