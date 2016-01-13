package blog

import (
	"webcenter/module"
)

const ID = "f17133ec-63e9-4b46-8757-e6ca1af6fe3e"
const URI = "/blog/"

type blog struct {
	
}

var instance *blog = nil

func init() {
	instance = &blog{}
	
	module.RegisterModule(instance)
}

func (this *blog) Startup(e module.Entity) {
	registerRouter()	
}

func (this *blog) Cleanup() {
	
}

func (this *blog) ID() string {
	return ID
}

func (this *blog) Uri() string {
	return URI
}

