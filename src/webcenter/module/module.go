package module

import (
	"log"
    "webcenter/util/modelhelper"
    "webcenter/module/model"
)

const GET = "get"
const POST = "post"

type Entity interface {
	ID() string
	Name() string
	Description() string
	EnableState() bool
	Enable()
	Disable()
	DefaultState() bool
	Default()
	Undefault()
	Internal() bool
}

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

type Router interface {
	AddGetRoute(pattern string, handler interface{})
	RemoveGetRoute(pattern string)
	
	AddPostRoute(pattern string, handler interface{})
	RemovePostRoute(pattern string)	
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
		
		modules := QueryAllModules()
		for _, m := range modules {
			entityIDMap[m.ID()] = m
		}
		
		initializeFlag = true
	}
	
	log.Println(entityIDMap)
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

func StartupAllModules(r Router) {
	log.Println("StartupAllModules all modules")
	
	for i, m := range moduleIDMap {
		e,ok := entityIDMap[i]
		if !ok {
			log.Println("illegal module id")
			continue
		}
		
		if !e.EnableState() {
			continue			
		}
		
		routes := m.Routes()
		for i, _ := range routes {
			rt := routes[i]
			
			pattern := rt.Pattern()
			if !e.DefaultState() {
				pattern = m.Uri() + rt.Pattern()
			}
			
			if rt.Type() == GET {
				r.AddGetRoute(pattern, rt.Handler())
			} else if rt.Type() == POST {
				r.AddPostRoute(pattern, rt.Handler())
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
		
		if e.EnableState() {
			m.Cleanup()
		}
	}	
}

func QueryAllModules() []Entity {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	entities := model.QueryAll(helper)
	allEntity := []Entity{}
	for _, e := range entities {
		allEntity = append(allEntity, e)
	}
	
	return allEntity
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

	model.Destroy(helper,id)
	
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
	
	return model.Save(helper,e)
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
	
	return model.Save(helper,e)
}

func UndefaultAllModule() {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer helper.Release()
	
	for _, e := range entityIDMap {
		e.Undefault()
		
		model.Save(helper,e)
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
	
	return model.Save(helper,e)
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
	
	return model.Save(helper,e)
}


