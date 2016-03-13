package module

import (
	"log"
	"webcenter/router"
)

const (
	GET = "get"
	POST = "post"
)

type Route interface {
	Type() string
	Pattern() string
	Handler() interface{}
}

type Module interface {	
	ID() string
	Name() string
	Description() string
	
	Uri() string
	Routes() []Route	

	Startup()
	Cleanup()
}

type route struct {
	rType string
	rPattern string
	rHandler interface{}
}

func (this *route) Type() string {
	return this.rType
}

func (this *route) Pattern() string {
	return this.rPattern
}

func (this *route) Handler() interface{} {
	return this.rHandler
}

func NewRoute(rType, rPattern string, rHandler interface{}) Route {
	r := route{}
	r.rType = rType
	r.rPattern = rPattern
	r.rHandler = rHandler
	
	return &r	
}

var moduleIDMap = map[string]Module{}

func QueryAllModule() []Module {
	modules := []Module{}
	
	for _, m := range moduleIDMap {
		modules = append(modules, m)
	}
	
	return modules
}

func FindModule(id string) (Module, bool) {
	m,found := moduleIDMap[id]
	
	return m,found
}

func RegisterModule(m Module) {
	moduleIDMap[m.ID()] = m	
}

func UnregisterModule(id string) {
	delete(moduleIDMap, id)
}

func StartupAllModules() {
	log.Println("StartupAllModules all modules")
	
	for _, m := range moduleIDMap {
		
		routes := m.Routes()
		for i, _ := range routes {
			rt := routes[i]
			
			if rt.Type() == GET {				
				pattern := m.Uri() + rt.Pattern()
				router.AddGetRoute(pattern, rt.Handler())

			} else if rt.Type() == POST {
				pattern := m.Uri() + rt.Pattern()
				router.AddPostRoute(pattern, rt.Handler())
				
			} else {
				panic("illegal route type, type:" + rt.Type() )
			}
		}
		
		m.Startup()
	}
}

func CleanupAllModules() {
	for _, m := range moduleIDMap {
		m.Cleanup()
	}	
}

