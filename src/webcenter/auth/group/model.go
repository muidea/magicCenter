package group

import (
	"fmt"
	"webcenter/modelhelper"
)

var ADMIN_GROUP = 0
var COMMON_GROUP = 1

type GroupInfo struct {
	Id int
	Name string
	UserCount int
	Catalog int	
}

type Group struct {
	Id int
	Name string
	Catalog int
}

func newGroup() Group {
	group := Group{}
	group.Id = -1
	
	return group;
}

func (group Group)AdminGroup() bool {
	return group.Catalog == ADMIN_GROUP
}

func QueryAllGroup(model modelhelper.Model) []GroupInfo {
	groupInfoList := []GroupInfo{}
	sql := fmt.Sprintf("select g.id, g.`name`, count(u.id) count, g.catalog from `group` g, `user` u where g.id = u.`group` union select g.id, g.`name`, 0 count, g.catalog from `group` g where g.id not in ( select  `group` from `user` )")
	if !model.Query(sql) {
		panic("query failed")
	}

	for model.Next() {
		group := GroupInfo{}
		model.GetValue(&group.Id, &group.Name, &group.UserCount, &group.Catalog)
		
		groupInfoList = append(groupInfoList, group)
	}

	return groupInfoList
}

func QueryGroupById(model modelhelper.Model, id int) (Group, bool) {
	group := newGroup()
	sql := fmt.Sprintf("select id,name,catalog from `group` where id=%d",id)
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false;
	for model.Next() {
		model.GetValue(&group.Id, &group.Name, &group.Catalog)
		result = true
	}

	return group, result
}

func DeleteGroup(model modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from `group` where id =%d", id)
	result := model.Execute(sql)
	if !result {
		panic("query failed")
	}
	
	return result	
}

func SaveGroup(model modelhelper.Model, group Group) bool {
	sql := fmt.Sprintf("select id from `group` where id=%d", group.Id)
	if !model.Query(sql) {
		panic("query failed")
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `group` (name,catalog) values ('%s',1)", group.Name)
	} else {
		// modify
		sql = fmt.Sprintf("update `group` set name ='%s' where id=%d", group.Name, group.Id)
	}
	
	return model.Execute(sql)
}



