package dal

import (
	"fmt"
	"magiccenter/kernel/module/model"
	"magiccenter/util/modelhelper"
)

func DeleteModule(helper modelhelper.Model, id string) {
	sql := fmt.Sprintf("delete from module where id='%s'", id)
	num, ret := helper.Execute(sql)
	if num == 1 && ret {
		sql = fmt.Sprintf("delete from block where owner='%s'", id)
		helper.Execute(sql)
	}
}

func QueryModule(helper modelhelper.Model, id string) (model.Module, bool) {
	m := model.Module{}
	sql := fmt.Sprintf("select id, name, description, uri, enableflag from module where id='%s'", id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&m.Id, &m.Name, &m.Description, &m.Uri, &m.EnableFlag)
		result = true
	}

	return m, result
}

func QueryAllModule(helper modelhelper.Model) []model.Module {
	moduleList := []model.Module{}

	sql := fmt.Sprintf("select id, name, description, uri, enableflag from module")
	helper.Query(sql)

	for helper.Next() {
		m := model.Module{}
		helper.GetValue(&m.Id, &m.Name, &m.Description, &m.Uri, &m.EnableFlag)

		moduleList = append(moduleList, m)
	}

	return moduleList
}

func SaveModule(helper modelhelper.Model, m model.Module) (model.Module, bool) {
	result := false
	_, found := QueryModule(helper, m.Id)
	if found {
		sql := fmt.Sprintf("update module set name ='%s', description ='%s', uri='%s', enableflag =%d where Id='%s'", m.Name, m.Description, m.Uri, m.EnableFlag, m.Id)
		_, result = helper.Execute(sql)
	} else {
		sql := fmt.Sprintf("insert into module(id, name, description, uri, enableflag) values ('%s','%s','%s','%s',%d)", m.Id, m.Name, m.Description, m.Uri, m.EnableFlag)
		_, result = helper.Execute(sql)
	}

	return m, result
}
