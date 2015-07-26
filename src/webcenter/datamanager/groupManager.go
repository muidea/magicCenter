package datamanager

import (
	"fmt"
	"log"
	"muidea.com/dao"
)

type GroupManager struct {
	groupInfo map[int]Group
	dao *dao.Dao
}

func (this *GroupManager) Load() bool {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		log.Printf("fetch dao failed, err:%s", err.Error())
		return false
	}
	
	this.groupInfo = make(map[int]Group)
	this.dao = dao	
	return true
}

func (this *GroupManager) Unload() {
	this.dao.Release()
	this.dao = nil
	this.groupInfo = nil
}

func (this * GroupManager) AddGroup(group Group) bool {
	sql := fmt.Sprintf("insert into magicid_db.group value (%d, %s, %d)", group.id, group.name, group.catalog)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}
		
	return true
}

func (this * GroupManager) ModGroup(group Group) bool {
	sql := fmt.Sprintf("update magicid_db.group set name ='%s', catalog=%d where id =%d", group.name, group.catalog, group.id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}
	
	this.groupInfo[group.id] = group	
	return true
}

func (this * GroupManager) DelGroup(id int) {	
	delete(this.groupInfo, id)
	
	sql := fmt.Sprintf("delete from magicid_db.group where id =%d", id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return 
	}
}

func (this * GroupManager) FindGroupById(id int) (Group, bool) {
	group, found := this.groupInfo[id]
	if !found {
		sql := fmt.Sprintf("%s","select id,name,catalog from magicid_db.group where id=%d", id)
		if !this.dao.Query(sql) {
			log.Printf("query failed, sql:%s", sql)
			return group, false
		}
	
		for this.dao.Next() {
			group := Group{}
			this.dao.GetField(&group.id, &group.name, &group.catalog)
			this.groupInfo[group.id] = group
		}
	
	}
	group, found = this.groupInfo[id]
	
	return group, found
}

