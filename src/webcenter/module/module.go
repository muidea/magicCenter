package module

import (
    "webcenter/modelhelper"
)

func QueryAllModules() []Module {
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

func UninstallModules(id int) {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()

	destroy(model,id)		
}

func EnableModule(id int) bool {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	m,found := query(model,id)
	if !found {
		return false
	}
	
	m.Enable()
	
	return save(model,m)
}

func DisableModule(id int) bool {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	m,found := query(model,id)
	if !found {
		return false
	}
	
	m.Disable()
	
	return save(model,m)

}

func DefaultModule(id int) bool {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	m,found := query(model,id)
	if !found {
		return false
	}
	
	m.Default()
	
	return save(model,m)
}

func UndefaultModule(id int) bool {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	m,found := query(model,id)
	if !found {
		return false
	}
	
	m.Undefault()
	
	return save(model,m)
}


