package module

import (
	"fmt"
	"webcenter/util/modelhelper"
)

type Block interface {
	ID() int
	Name() string
	Owner() string
}

type block struct {
	id int
	name string
	owner string
}

func (this *block)ID() int {
	return this.id
}

func (this *block)Name() string {
	return this.name
}

func (this *block)Owner() string {
	return this.owner
}

type Entity interface {
	ID() string
	Name() string
	Description() string
	Enable()
	Disable()
	EnableStatus() int
	Default()
	Undefault()
	DefaultStatus() int
	Internal() bool
	StyleFlag() int
}

type entity struct {
	id string
	name string
	description string
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

func (this *entity)Enable() {
	this.enableFlag = 1
}

func (this *entity)Disable() {
	this.enableFlag = 0
}

func (this *entity)EnableStatus() int {
	return this.enableFlag
}

func (this *entity)Default() {
	this.defaultFlag = 1
}

func (this *entity)Undefault() {
	this.defaultFlag = 0
}

func (this *entity)DefaultStatus() int {
	return this.defaultFlag
}

func (this *entity)Internal() bool {
	return this.styleFlag == 0
}

func (this *entity)StyleFlag() int {
	return this.styleFlag
}

func newEntity(id,name,description string) Entity {
	e := &entity{}
	e.id = id
	e.name = name
	e.description = description
	
	return e
}

func insertBlock(helper modelhelper.Model, name,owner string) {
	sql := fmt.Sprintf("insert into module_block (name,owner) values('%s','%s')", name, owner)
	helper.Execute(sql)
}

func deleteBlock(helper modelhelper.Model, name string) {
	sql := fmt.Sprintf("delete from module_block where name='%s'", name)
	helper.Execute(sql)
}

func queryEntityBlocks(helper modelhelper.Model, owner string) []Block {	
	blockList := []Block{}
	sql := fmt.Sprintf("select id,name,owner from module_block where owner='%s'", owner)
	if helper.Query(sql) {
		for helper.Next() {
			b := &block{}
			helper.GetValue(&b.id, &b.name, &b.owner)
			
			blockList = append(blockList, b)
		}
	} else {
		panic("execute sql failed")
	}
	
	return blockList
}

func deleteEntity(helper modelhelper.Model, id string) {
	helper.BeginTransaction()
	
	sql := fmt.Sprintf("delete from module_block where owner='%s'", id)
	if helper.Execute(sql) {
		sql = fmt.Sprintf("delete from module where id='%s'", id)
		if helper.Execute(sql) {
			helper.Commit()
		} else {
			helper.Rollback()			
		}
	} else {
		helper.Rollback()
	}
	
}

func queryEntity(helper modelhelper.Model, id string) (Entity, bool) {
	m := &entity{}
	sql := fmt.Sprintf("select id, name, description, enableflag, defaultflag, styleflag from module where id='%s'", id)	
	if !helper.Query(sql) {
		panic("execute sql failed")
	}
	
	result := false
	if helper.Next() {
		helper.GetValue(&m.id, &m.name, &m.description, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		result = true
	}
		
	return m, result	
}

func queryAllEntity(helper modelhelper.Model) []Entity {
	moduleList := []Entity{}
	
	sql := fmt.Sprintf("select id, name, description, enableflag, defaultflag, styleflag from module order by styleflag")	
	if !helper.Query(sql) {
		panic("execute sql failed")
	}
	
	for helper.Next() {
		m := &entity{}
		helper.GetValue(&m.id, &m.name, &m.description, &m.enableFlag, &m.defaultFlag, &m.styleFlag)
		
		moduleList = append(moduleList, m)
	}
	
	return moduleList
}

func saveEntity(helper modelhelper.Model, m Entity) bool {
	_, found := queryEntity(helper, m.ID())
	if found {
		sql := fmt.Sprintf("update module set Name ='%s', Description ='%s', enableflag =%d, defaultflag =%d where Id='%s'", m.Name(), m.Description(), m.EnableStatus(), m.DefaultStatus(), m.ID())
		if !helper.Execute(sql) {
			panic("execute sql failed")
		}
	} else {
		sql := fmt.Sprintf("insert into module(id, name, description, enableflag, defaultflag, styleflag) values ('%s','%s','%s',%d,%d,%d)", m.ID(), m.Name(), m.Description(), m.EnableStatus(), m.DefaultStatus(), m.StyleFlag())
		if !helper.Execute(sql) {
			panic("execute sql failed")
		}
	}
		
	return true
}






