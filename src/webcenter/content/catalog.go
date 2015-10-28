package content

import (
	"fmt"
	"log"
	"muidea.com/dao"
	"webcenter/auth"
)

type Catalog struct {
	Id int
	Name string
	Creater auth.User
	Pid int
}

func newCatalog() Catalog {
	catalog := Catalog{}
	catalog.Id = -1
	catalog.Creater = auth.NewUser()
	catalog.Pid = 0
	
	return catalog
}

func GetAllCatalog(dao * dao.Dao) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog")
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for dao.Next() {
		catalog := newCatalog()
		dao.GetField(&catalog.Id, &catalog.Name, &catalog.Creater.Id, &catalog.Pid)
		
		catalogList = append(catalogList, catalog)
	}
	
	for i :=0; i<len(catalogList); i++ {
		catalog := &catalogList[i]
		if !catalog.Creater.Query(dao) {
			catalog.Creater, _ = auth.QueryDefaultUser(dao)
		}
	}
	
	return catalogList
}

func GetAvalibleParentCatalog(id int, dao * dao.Dao) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where id < %d", id)
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for dao.Next() {
		catalog := newCatalog()
		dao.GetField(&catalog.Id, &catalog.Name, &catalog.Creater.Id, &catalog.Pid)
		
		catalogList = append(catalogList, catalog)
	}
	
	for i :=0; i<len(catalogList); i++ {
		catalog := &catalogList[i]
		if !catalog.Creater.Query(dao) {
			catalog.Creater, _ = auth.QueryDefaultUser(dao)
		}
	}
	
	return catalogList
}

func GetSubCatalog(id int, dao * dao.Dao) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where pid = %d", id)
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for dao.Next() {
		catalog := newCatalog()
		dao.GetField(&catalog.Id, &catalog.Name, &catalog.Creater.Id, &catalog.Pid)
		
		catalogList = append(catalogList, catalog)
	}
	
	for i :=0; i<len(catalogList); i++ {
		catalog := &catalogList[i]
		if !catalog.Creater.Query(dao) {
			catalog.Creater, _ = auth.QueryDefaultUser(dao)
		}
	}
	
	return catalogList
}

func (this *Catalog)Query(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return false
	}

	result := false
	for dao.Next() {
		dao.GetField(&this.Id, &this.Name, &this.Creater.Id, &this.Pid)
		result = true
	}
	
	if result {
		result = this.Creater.Query(dao)
		if !result {
			this.Creater, result = auth.QueryDefaultUser(dao)
		}
	}
	
	return result	
}


func (this *Catalog)delete(dao *dao.Dao) {
	sql := fmt.Sprintf("delete from catalog where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete catalog failed, sql:%s", sql)
	}
}

func (this *Catalog)save(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id from catalog where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
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
		sql = fmt.Sprintf("insert into catalog (name,creater,pid) values ('%s',%d,%d)", this.Name, this.Creater.Id, this.Pid)
	} else {
		// modify
		sql = fmt.Sprintf("update catalog set name ='%s', creater =%d, pid=%d where id=%d", this.Name, this.Creater.Id, this.Pid, this.Id)
	}
	
	result = dao.Execute(sql)
	
	return result
}




