package dal

import (
	"fmt"
	"webcenter/util/modelhelper"
)
type Module struct {
	Id string
	Name string
	Description string
	EnableFlag int
}


func DeleteModule(helper modelhelper.Model, id string) {
	sql := fmt.Sprintf("delete from module where id='%s'", id)
	num, ret := helper.Execute(sql)
	if num == 1 && ret {
		sql = fmt.Sprintf("delete from block where owner='%s'", id)
		helper.Execute(sql)
	}
}

func QueryModule(helper modelhelper.Model, id string) (Module, bool) {
	m := Module{}
	sql := fmt.Sprintf("select id, name, description, enableflag from module where id='%s'", id)	
	helper.Query(sql)
	
	result := false
	if helper.Next() {
		helper.GetValue(&m.Id, &m.Name, &m.Description, &m.EnableFlag)
		result = true
	}
	
	return m, result	
}

func QueryAllModule(helper modelhelper.Model) []Module {
	moduleList := []Module{}
	
	sql := fmt.Sprintf("select id, name, description, enableflag from module order by styleflag")	
	helper.Query(sql)
	
	for helper.Next() {
		m := Module{}
		helper.GetValue(&m.Id, &m.Name, &m.Description, &m.EnableFlag)
		
		moduleList = append(moduleList, m)
	}
	
	return moduleList
}

func SaveModule(helper modelhelper.Model, m Module) (Module, bool) {
	result := false
	_, found := QueryModule(helper, m.Id)
	if found {
		sql := fmt.Sprintf("update module set name ='%s', description ='%s', enableflag =%d where Id='%s'", m.Name, m.Description, m.EnableFlag, m.Id)
		_, result = helper.Execute(sql)
	} else {
		sql := fmt.Sprintf("insert into module(id, name, description, enableflag) values ('%s','%s','%s',%d)", m.Id, m.Name, m.Description, m.EnableFlag)
		_, result = helper.Execute(sql)
	}
	
	return m, result
}


