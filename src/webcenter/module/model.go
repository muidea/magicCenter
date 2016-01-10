package module

import (
	"fmt"
	"webcenter/modelhelper"
)

type Module interface {
	ID() int
	Name() string
	Description() string
	Uri() string
	EnableState() bool
	Enable()
	Disable()
	DefaultState() bool
	Default()
	Undefault()
	Internal() bool
}

type module struct {
	id int
	name string
	description string
	uri string
	enableFlag int
	defaultFlag int
	styleFlag int
}

func (this *module)ID() int {
	return this.id
}

func (this *module)Name() string {
	return this.name
}

func (this *module)Description() string {
	return this.description
}

func (this *module)Uri() string {
	return this.uri
}

func (this *module)EnableState() bool {
	return this.enableFlag == 1
}

func (this *module)Enable() {
	this.enableFlag = 1
}

func (this *module)Disable() {
	this.enableFlag = 0
}

func (this *module)DefaultState() bool {
	return this.defaultFlag == 1
}

func (this *module)Default() {
	this.defaultFlag = 1
}

func (this *module)Undefault() {
	this.defaultFlag = 0
}

func (this *module)Internal() bool {
	return this.styleFlag == 0
}

func create(model modelhelper.Model, name,description,uri string) (Module,bool) {
	m := &module{}
	m.name = name
	m.description = description
	m.uri = uri
	m.enableFlag = 0
	m.defaultFlag = 0
	m.styleFlag = 0
	
	sql := fmt.Sprintf("insert into module(name,description,uri,enableflag,defaultflag,styleflag) values('%s','%s','%s',%d,%d,%d)", m.name, m.description, m.uri, m.enableFlag, m.defaultFlag, m.styleFlag)
	if !model.Execute(sql) {
		panic("execute sql failed")
	}
	
	sql = fmt.Sprintf("select id from module where name='%s' and uri='%s'", m.name, m.uri)
	if !model.Query(sql) {
		panic("execute sql failed")
	}
	
	result := false
	if model.Next() {
		model.GetValue(&m.id)
		result = true
	}
	
	return m, result
}

func destroy(model modelhelper.Model, id int) {
	sql := fmt.Sprintf("delete from module where id=%d", id)
	if !model.Execute(sql) {
		panic("execute sql failed")
	}
}

func query(model modelhelper.Model, id int) (Module,bool) {
	m := &module{}
	sql := fmt.Sprintf("select id, name,description,uri,enableflag,defaultflag,styleflag from module where id=%d", id)	
	if !model.Query(sql) {
		panic("execute sql failed")
	}
	
	result := false
	if model.Next() {
		model.GetValue(&m.id, &m.name, &m.description, &m.uri, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		result = true
	}
	
	return m, result	
}

func queryAll(model modelhelper.Model) []Module {
	moduleList := []Module{}
	
	sql := fmt.Sprintf("select id, name,description,uri,enableflag,defaultflag,styleflag from module")	
	if !model.Query(sql) {
		panic("execute sql failed")
	}
	
	for model.Next() {
		m := &module{}
		model.GetValue(&m.id, &m.name, &m.description, &m.uri, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		
		moduleList = append(moduleList, m)
	}
	
	return moduleList
}

func save(model modelhelper.Model, m Module) bool {
	_, found := query(model, m.ID())
	if !found {
		return false
	}
	
	sql := fmt.Sprintf("update module set name ='%s', description ='%s', uri ='%s', enableflag =%d, defaultflag =%d, styleflag =%d where id=%d", m.Name(), m.Description(), m.Uri(), m.EnableState(), m.DefaultState(), m.Internal(), m.ID())
	if !model.Execute(sql) {
		panic("execute sql failed")
	}
	
	return true
}






