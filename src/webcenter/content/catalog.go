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
		
		catalog.Creater.Query(dao)
		
		catalogList = append(catalogList, catalog)
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


func (this *Catalog)insert(dao *dao.Dao) bool {
	sql := fmt.Sprintf("insert into catalog value (%d, %s, %d)", this.Id, this.Name, this.Creater.Id)
	if !dao.Execute(sql) {
		log.Printf("inser catalog failed, sql:%s", sql)
		return false
	}
		
	return true
}

func (this *Catalog)update(dao *dao.Dao) bool {
	sql := fmt.Sprintf("update catalog set name ='%s' where id =%d", this.Name, this.Id)
	if !dao.Execute(sql) {
		log.Printf("update catalog failed, sql:%s", sql)
		return false
	}
	
	return true
}

func (this *Catalog)remove(dao *dao.Dao) {
	sql := fmt.Sprintf("delete from catalog where id =%d", this.Id)
	if !dao.Execute(sql) {
		log.Printf("delete catalog failed, sql:%s", sql)
	}
}



