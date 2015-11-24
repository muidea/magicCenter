package group

import (
	"fmt"
	"log"
	"webcenter/modelhelper"
)

var ADMIN_GROUP = 0
var COMMON_GROUP = 1


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


func GetAllGroup(model modelhelper.Model) []Group {
	groupList := []Group{}
	sql := fmt.Sprintf("select id,name,catalog from `group`")
	if !model.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return groupList
	}

	for model.Next() {
		group := newGroup()
		model.GetValue(&group.Id, &group.Name, &group.Catalog)
		
		groupList = append(groupList, group)
	}

	return groupList
}

func GetAllSubGroup(model modelhelper.Model, id int) []Group {
	groupList := []Group{}
	sql := fmt.Sprintf("select id,name,catalog from `group` where catalog=%d",id)
	if !model.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return groupList
	}

	for model.Next() {
		group := newGroup()
		model.GetValue(&group.Id, &group.Name, &group.Catalog)
		
		groupList = append(groupList, group)
	}

	return groupList
}


func GetGroupById(model modelhelper.Model, id int) (Group, bool) {
	group := newGroup()
	sql := fmt.Sprintf("select id,name,catalog from `group` where id=%d",id)
	if !model.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return group, false
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
		log.Printf("delete group failed, sql:%s", sql)
	}
	
	return result	
}

func SaveGroup(model modelhelper.Model, group Group) bool {
	sql := fmt.Sprintf("select id from `group` where id=%d", group.Id)
	if !model.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return false
	}

	result := false;
	for model.Next() {
		var id = 0
		result = model.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `group` (name,catalog) values ('%s',%d)", group.Name, group.Catalog)
	} else {
		// modify
		sql = fmt.Sprintf("update `group` set name ='%s', catalog =%d where id=%d", group.Name, group.Catalog, group.Id)
	}
	
	result = model.Execute(sql)
	
	return result
}



