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

func (this *Group)inert(dao *dao.Dao) bool {
	sql := fmt.Sprintf("insert into `group` value (%d, %s, %d)", this.Id, this.Name, this.Catalog)
	if !dao.Execute(sql) {
		log.Printf("insert group failed, sql:%s", sql)
		return false
	}
	
	return true	
}

func (this *Group)update(dao *dao.Dao) bool {
	sql := fmt.Sprintf("update `group` set name ='%s', catalog=%d where id =%d", this.Name, this.Catalog, this.Id)
	if !dao.Execute(sql) {
		log.Printf("update group failed, sql:%s", sql)
		return false
	}
	
	return true	
}


func (this *Group)remove(dao *dao.Dao) {
	sql := fmt.Sprintf("delete from `group` where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete group failed, sql:%s", sql)
		return
	}
}



