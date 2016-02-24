package module

import (
	"log"
	"webcenter/util/modelhelper"
	"webcenter/router"
)

const GET = "get"
const POST = "post"

type Route interface {
	Type() string
	Pattern() string
	Handler() interface{}
}

type Instance interface {
	Startup()
	Cleanup()

	ID() string	
	Uri() string
	Routes() []Route
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

var initializeFlag = false
var moduleIDMap = map[string]Module{}
var instanceIDMap = map[string]Instance{}

func init() {
	log.Println("init module manager")
	
	if !initializeFlag {
		moduleIDMap = make(map[string]Module)
		instanceIDMap = make(map[string]Instance)

		helper, err := modelhelper.NewHelper()
		if err != nil {
			panic("construct model failed")
		}
		defer helper.Release()
		
		modules := queryAllModule(helper)
		for _, m := range modules {
			moduleIDMap[m.ID()] = m
		}
		
		initializeFlag = true
	}
}

func RegisterModule(m Instance) bool {
	_, ok := moduleIDMap[m.ID()]
	if !ok {
		return false
	}
	
	_, ok = instanceIDMap[m.ID()]
	if ok {
		panic("duplicate register module, id:" + m.ID())
	}
	
	instanceIDMap[m.ID()] = m
	
	return true
}

func UnregisterModule(id string) {
	_, ok := instanceIDMap[id]
	if ok {
		delete(instanceIDMap, id)
	} else {
		log.Println("illegal module id, id:" + id)
	}
}

func StartupAllModules() {
	log.Println("StartupAllModules all modules")
	
	for i, m := range instanceIDMap {
		e,ok := moduleIDMap[i]
		if !ok {
			log.Println("illegal module id")
			continue
		}
		
		if e.EnableStatus() != 1 {
			continue
		}
		
		routes := m.Routes()
		for i, _ := range routes {
			rt := routes[i]
			
			if rt.Type() == GET {
				if e.DefaultStatus() == 1 {
					router.AddGetRoute(rt.Pattern(), rt.Handler())
				}
				
				pattern := m.Uri() + rt.Pattern()
				router.AddGetRoute(pattern, rt.Handler())

			} else if rt.Type() == POST {
				if e.DefaultStatus() == 1 {
					router.AddPostRoute(rt.Pattern(), rt.Handler())
				}
				
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
	for i, m := range instanceIDMap {
		e,ok := moduleIDMap[i]
		if !ok {
			log.Println("illegal module id")
			continue
		}
		
		if e.EnableStatus() == 1 {
			m.Cleanup()
		}
	}	
}

func QueryModule(id string) (Module, bool) {
	e, found := moduleIDMap[id]
	
	return e,found
}

func InstallModule(modulePath string) bool {
	return true
}

func UninstallModule(id string) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("illegal module id,id:" + id)
	}
	defer helper.Release()

	deleteModule(helper,id)
	
	delete(moduleIDMap, id)
}

func EnableModule(id string) bool {
	e, ok := moduleIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
	
	m, ok := instanceIDMap[id]
	if !ok {
		panic("illegal module id,id:" + id)
	}
	
	m.Startup()
	
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
		
	e.Enable()
	
	return saveModule(helper,e)
}

func DisableModule(id string) bool {
	e, ok := moduleIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
	
	m, ok := instanceIDMap[id]
	if !ok {
		panic("illegal module id,id:" + id)
	}
	
	m.Cleanup()
	
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
		
	e.Disable()
	
	return saveModule(helper,e)
}

func UndefaultAllModule() {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	for _, e := range moduleIDMap {
		e.Undefault()
		
		saveModule(helper,e)
	}
}

func DefaultModule(id string) bool {
	e, ok := moduleIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
	
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	e.Default()
		
	return saveModule(helper,e)
}

func UndefaultModule(id string) bool {
	e, ok := moduleIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
	
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	e.Undefault()
	
	return saveModule(helper, e)
}


