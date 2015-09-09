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
}

func newCatalog() Catalog {
	catalog := Catalog{}
	catalog.Id = -1
	catalog.Creater = auth.NewUser()
	
	return catalog
}

func GetAllCatalog(dao * dao.Dao) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater from catalog")
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for dao.Next() {
		catalog := newCatalog()
		dao.GetField(&catalog.Id, &catalog.Name, &catalog.Creater.Id)
		
		catalogList = append(catalogList, catalog)
	}
	
	for i :=0; i<len(catalogList); i++ {
		catalog := &catalogList[i]
		catalog.Creater.Query(dao)
	}
	
	return catalogList
}

func (this *Catalog)Query(dao *dao.Dao) bool {
	sql := fmt.Sprintf("select id,name,creater from catalog where id=%d", this.Id)
	if !dao.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return false
	}

	result := false
	for dao.Next() {
		dao.GetField(&this.Id, &this.Name, &this.Creater.Id)
		result = true
	}
	
	if result {
		result = this.Creater.Query(dao)
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
		sql = fmt.Sprintf("insert into catalog (name,creater) values ('%s',%d)", this.Name, this.Creater.Id)
	} else {
		// modify
		sql = fmt.Sprintf("update catalog set name ='%s', creater =%d where id=%d", this.Name, this.Creater.Id, this.Id)
	}
	
	result = dao.Execute(sql)
	
	return result
}




