package datamanager

import (
	"fmt"
	"log"
	"webcenter/model/dao"
)

type RoutelineManager struct {
	routelineInfo map[int]Routeline
	dao           *dao.Dao
}

func (this *RoutelineManager) Load() bool {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		log.Printf("fetch dao failed, err:%s", err.Error())
		return false
	}

	this.routelineInfo = make(map[int]Routeline)
	this.dao = dao
	return true
}

func (this *RoutelineManager) Unload() {
	this.dao.Release()
	this.dao = nil
	this.routelineInfo = nil
}

func (this *RoutelineManager) AddRouteline(routeline Routeline) bool {
	sql := fmt.Sprintf("insert into magicid_db.routeline value (%d, %s, %s, %d)", routeline.Id, routeline.Name, routeline.Description,routeline.Creater)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}

	return true
}

func (this *RoutelineManager) ModRouteline(routeline Routeline) bool {
	sql := fmt.Sprintf("update magicid_db.routeline set name ='%s', description='%s' where id =%d", routeline.Name, routeline.Description, routeline.Id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return false
	}

	this.routelineInfo[routeline.Id] = routeline
	return true
}

func (this *RoutelineManager) DelRouteline(id int) {
	delete(this.routelineInfo, id)

	sql := fmt.Sprintf("delete from magicid_db.routeline where id =%d", id)
	if !this.dao.Execute(sql) {
		log.Printf("execute failed, sql:%s", sql)
		return
	}
}

func (this *RoutelineManager) FindRoutelineById(id int) (Routeline, bool) {
	routeline, found := this.routelineInfo[id]
	if !found {
		sql := fmt.Sprintf("select * from magicid_db.routeline where id=%d", id)
		if !this.dao.Query(sql) {
			log.Printf("query failed, sql:%s", sql)
			return routeline, false
		}

		for this.dao.Next() {
			routeline := Routeline{}
			this.dao.GetField(&routeline.Id, &routeline.Name, &routeline.Description, &routeline.Creater)
			this.routelineInfo[routeline.Id] = routeline
		}
	}
	routeline, found = this.routelineInfo[id]

	return routeline, found
}

func (this *RoutelineManager) GetAll() []Routeline {
	ret := []Routeline{}
	if len(this.routelineInfo) == 0 {
		log.Println("select routeline info from database.")
		sql := "select * from magicid_db.routeline order by id"
		if !this.dao.Query(sql) {
			log.Printf("query failed, sql:%s", sql)
			return ret
		}

		for this.dao.Next() {
			routeline := Routeline{}
			this.dao.GetField(&routeline.Id, &routeline.Name, &routeline.Description, &routeline.Creater)
			this.routelineInfo[routeline.Id] = routeline
		}
	}

	for _, routeline := range this.routelineInfo {
		ret = append(ret, routeline)
	}

	return ret
}
