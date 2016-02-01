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

type Module interface {
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
var entityIDMap = map[string]Entity{}
var moduleIDMap = map[string]Module{}

func init() {
	log.Println("module init")
	
	if !initializeFlag {
		entityIDMap = make(map[string]Entity)
		moduleIDMap = make(map[string]Module)

		helper, err := modelhelper.NewHelper()
		if err != nil {
			panic("construct model failed")
		}
		defer helper.Release()
		
		modules := queryAllEntity(helper)
		for _, m := range modules {
			entityIDMap[m.ID()] = m
		}
		
		initializeFlag = true
	}
}

func RegisterModule(m Module) bool {
	_, ok := entityIDMap[m.ID()]
	if !ok {
		return false
	}
	
	_, ok = moduleIDMap[m.ID()]
	if ok {
		panic("duplicate register module, id:" + m.ID())
	}
	
	moduleIDMap[m.ID()] = m
	
	return true
}

func UnregisterModule(id string) {
	_, ok := moduleIDMap[id]
	if ok {
		delete(moduleIDMap, id)
	} else {
		log.Println("illegal module id, id:" + id)
	}
}

func StartupAllModules() {
	log.Println("StartupAllModules all modules")
	
	for i, m := range moduleIDMap {
		e,ok := entityIDMap[i]
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
	for i, m := range moduleIDMap {
		e,ok := entityIDMap[i]
		if !ok {
			log.Println("illegal module id")
			continue
		}
		
		if e.EnableStatus() == 1 {
			m.Cleanup()
		}
	}	
}

func QueryAllEntities() []Entity {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	return queryAllEntity(helper)
}

func QueryModuleEntity(id string) (Entity,bool) {
	e, found := entityIDMap[id]
	
	return e,found
}

func QueryModule(id string) (Module,bool) {
	m, found := moduleIDMap[id]
	
	return m,found
}

func InstallModules(modulePath string) bool {
	return true
}

func UninstallModules(id string) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("illegal module id,id:" + id)
	}
	defer helper.Release()

	deleteEntity(helper,id)
	
	delete(entityIDMap, id)
}

func EnableModule(id string) bool {
	e, ok := entityIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
	
	m, ok := moduleIDMap[id]
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
	
	return saveEntity(helper,e)
}

func DisableModule(id string) bool {
	e, ok := entityIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
	
	m, ok := moduleIDMap[id]
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
	
	return saveEntity(helper,e)
}

func UndefaultAllModule() {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	for _, e := range entityIDMap {
		e.Undefault()
		
		saveEntity(helper,e)
	}
}

func DefaultModule(id string) bool {
	e, ok := entityIDMap[id]
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
		
	return saveEntity(helper,e)
}

func UndefaultModule(id string) bool {
	e, ok := entityIDMap[id]
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
	
	return saveEntity(helper, e)
}


