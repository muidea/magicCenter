package module

import (
	"fmt"
	"webcenter/modelhelper"
)

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

type Module interface {
	Startup(e Entity)
	Cleanup()

	ID() string	
	Uri() string
}

type entity struct {
	id string
	name string
	description string
	uri string // 这个属性不记库，运行期赋值
	enableFlag int
	defaultFlag int
	styleFlag int
}

func (this *entity)ID() string {
	return this.id
}

func (this *entity)Name() string {
	return this.name
}

func (this *entity)Description() string {
	return this.description
}

func (this *entity)EnableState() bool {
	return this.enableFlag == 1
}

func (this *entity)Enable() {
	this.enableFlag = 1
}

func (this *entity)Disable() {
	this.enableFlag = 0
}

func (this *entity)DefaultState() bool {
	return this.defaultFlag == 1
}

func (this *entity)Default() {
	this.defaultFlag = 1
}

func (this *entity)Undefault() {
	this.defaultFlag = 0
}

func (this *entity)Internal() bool {
	return this.styleFlag == 0
}

func (this *entity)Uri() string {
	return this.uri
}

func create(model modelhelper.Model, id, name,description string) Entity {
	m := &entity{}
	m.id = id
	m.name = name
	m.description = description
	m.enableFlag = 0
	m.defaultFlag = 0
	m.styleFlag = 0
	
	sql := fmt.Sprintf("insert into module(id, name,description,uri,enableflag,defaultflag,styleflag) values('%s','%s','%s',%d,%d,%d)", m.id, m.name, m.description, m.enableFlag, m.defaultFlag, m.styleFlag)
	if !model.Execute(sql) {
		panic("execute sql failed")
	}
		
	return m
}

func destroy(model modelhelper.Model, id string) {
	sql := fmt.Sprintf("delete from module where id='%s'", id)
	if !model.Execute(sql) {
		panic("execute sql failed")
	}
}

func query(model modelhelper.Model, id string) (Entity,bool) {
	m := &entity{}
	sql := fmt.Sprintf("select id, name,description,enableflag,defaultflag,styleflag from module where id='%s'", id)	
	if !model.Query(sql) {
		panic("execute sql failed")
	}
	
	result := false
	if model.Next() {
		model.GetValue(&m.id, &m.name, &m.description, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		result = true
	}
	
	return m, result	
}

func queryAll(model modelhelper.Model) []Entity {
	moduleList := []Entity{}
	
	sql := fmt.Sprintf("select id,name,description,enableflag,defaultflag,styleflag from module order by styleflag")	
	if !model.Query(sql) {
		panic("execute sql failed")
	}
	
	for model.Next() {
		m := &entity{}
		model.GetValue(&m.id, &m.name, &m.description, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		
		moduleList = append(moduleList, m)
	}
	
	return moduleList
}

func save(model modelhelper.Model, m Entity) bool {
	_, found := query(model, m.ID())
	if !found {
		return false
	}
	
	name := m.Name()
	desc := m.Description()
	enableStatus := 0
	if m.EnableState() {
		enableStatus = 1
	}
	
	defaultStatus := 0
	if m.DefaultState() {
		defaultStatus = 1
	}
	sql := fmt.Sprintf("update module set name ='%s', description ='%s', enableflag =%d, defaultflag =%d where id='%s'", name, desc, enableStatus, defaultStatus, m.ID())
	if !model.Execute(sql) {
		panic("execute sql failed")
	}
	
	return true
}






