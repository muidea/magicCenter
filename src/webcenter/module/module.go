package module

import (
	"log"
    "webcenter/modelhelper"
)

var entityIDMap = map[string]Entity{}
var moduleIDMap = map[string]Module{}

func init() {
	log.Println("module init")
	entityIDMap = make(map[string]Entity)
	moduleIDMap = make(map[string]Module)
	
	modules := QueryAllModules()
	for _, m := range modules {
		entityIDMap[m.ID()] = m
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

func StarupAllModules() {
	for i, m := range moduleIDMap {
		e,ok := entityIDMap[i]
		if !ok {
			log.Println("illegal module id")
			continue
		}
		
		if e.EnableState() {
			m.Startup(e)
		}
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
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	return queryAll(model)	
}

func InstallModules(modulePath string) bool {
	return true
}

func UninstallModules(id string) {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("illegal module id,id:" + id)
	}
	defer model.Release()

	destroy(model,id)
	
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
	
	m.Startup(e)
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
		
	e.Enable()
	
	return save(model,e)
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
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
		
	e.Disable()
	
	return save(model,e)

}

func UndefaultAllModule() {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	for _, e := range entityIDMap {
		e.Undefault()
		
		save(model,e)
	}
}

func DefaultModule(id string) bool {
	e, ok := entityIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
		
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	e.Default()
	
	return save(model,e)
}

func UndefaultModule(id string) bool {
	e, ok := entityIDMap[id]
	if !ok {
		log.Println("illegal module id")
		return false
	}
		
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	e.Undefault()
	
	return save(model,e)
}


