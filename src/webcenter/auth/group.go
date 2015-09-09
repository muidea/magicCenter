package auth

import (
	"fmt"
	"log"
	"muidea.com/dao"
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


func GetAllGroup(dao * dao.Dao) []Group {
	groupList := []Group{}
	sql := fmt.Sprintf("select id,name,catalog from `group`")
	if !dao.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return groupList
	}

	for dao.Next() {
		group := newGroup()
		dao.GetField(&group.Id, &group.Name, &group.Catalog)
		
		groupList = append(groupList, group)
	}

	return groupList
}

func GetAllSubGroup(id int,dao * dao.Dao) []Group {
	groupList := []Group{}
	sql := fmt.Sprintf("select id,name,catalog from `group` where catalog=%d",id)
	if !dao.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return groupList
	}

	for dao.Next() {
		group := newGroup()
		dao.GetField(&group.Id, &group.Name, &group.Catalog)
		
		groupList = append(groupList, group)
	}

	return groupList
}


func (this *Group)IsAdminGroup() bool {
	return this.Catalog == ADMIN_GROUP
}

func (this *Group)query(dao *dao.Dao) bool {
		sql := fmt.Sprintf("select id,name,catalog from `group` where id=%d", this.Id)
		if !dao.Query(sql) {
			log.Printf("query group failed, sql:%s", sql)
			return false
		}
		
		result := false
		for dao.Next() {
			dao.GetField(&this.Id, &this.Name, &this.Catalog)
			result = true
		}
		
		return result
}


func (this *Group)delete(dao *dao.Dao) {
	sql := fmt.Sprintf("delete from `group` where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete group failed, sql:%s", sql)
		return
	}
}


func (this *Group)save(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id from `group` where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query group failed, sql:%s", sql)
		return false
	}

	result := false;
	for dao.Next() {
		var id = 0
		result = dao.GetField(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf("insert into `group` (name,catalog) values ('%s',%d)", this.Name, this.Catalog)
	} else {
		// modify
		sql = fmt.Sprintf("update `group` set name ='%s', catalog =%d where id=%d", this.Name, this.Catalog, this.Id)
	}
	
	result = dao.Execute(sql)
	
	return result
}


