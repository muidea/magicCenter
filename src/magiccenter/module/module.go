package module

import (
	"log"
	"magiccenter/router"
	"magiccenter/configuration"
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
	log.Printf("register module, name:%s, id:%s", m.Name(), m.ID())
	
	moduleIDMap[m.ID()] = m	
}

func UnregisterModule(id string) {
	log.Printf("register module, id:%s", id)
	
	delete(moduleIDMap, id)
}

func StartupAllModules() {
	log.Println("StartupAllModules all modules")
	
	defaultModule, _ := configuration.GetOption(configuration.SYS_DEFULTMODULE)
	
	for _, m := range moduleIDMap {
		
		routes := m.Routes()
		for i, _ := range routes {
			rt := routes[i]
			
			pattern := m.Uri() + rt.Pattern()
			if m.ID() == defaultModule {
				pattern = rt.Pattern()
			}
				
			
			if rt.Type() == GET {
				router.AddGetRoute(pattern, rt.Handler(), nil)

			} else if rt.Type() == POST {
				router.AddPostRoute(pattern, rt.Handler(), nil)
				
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

