package common

import (
	"log"
	"webcenter/application"
)

type Result struct {
	ErrCode int
	Reason string
}

func init() {
	log.Print("common.init")
	
	application.BindStatic(application.ResourcePath());
	application.BindStatic(application.StaticPath());	
}