package catalog


import (
	"fmt"
	"log"
	"webcenter/modelhelper"
)

type Catalog struct {
	Id int
	Name string
	Creater int
	Pid int
}

func newCatalog() Catalog {
	catalog := Catalog{}
	catalog.Id = -1
	catalog.Creater = -1
	catalog.Pid = 0
	
	return catalog
}


func GetAllCatalog(model modelhelper.Model) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog")
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for model.Next() {
		catalog := newCatalog()
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater, &catalog.Pid)
		
		catalogList = append(catalogList, catalog)
	}
		
	return catalogList
}

func QueryCatalogById(model modelhelper.Model, id int) (Catalog, bool) {
	catalog := Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where id=%d", id)
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalog, false
	}

	result := false
	for model.Next() {
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater, &catalog.Pid)
		result = true
	}

	return catalog, result	
}

func DeleteCatalogById(model modelhelper.Model, id int) bool {
	sql := fmt.Sprintf("delete from catalog where id=%d", id)
	
	result := model.Execute(sql)
	
	return result	
}

func GetAvalibleParentCatalog(model modelhelper.Model, id int) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where id < %d", id)
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for model.Next() {
		catalog := newCatalog()
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater, &catalog.Pid)
		
		catalogList = append(catalogList, catalog)
	}
	
	return catalogList
}

func GetSubCatalog(model modelhelper.Model, id int) []Catalog {
	catalogList := []Catalog{}
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where pid = %d", id)
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return catalogList
	}

	for model.Next() {
		catalog := newCatalog()
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater, &catalog.Pid)
		
		catalogList = append(catalogList, catalog)
	}
		
	return catalogList
}

func SaveCatalog(model modelhelper.Model, catalog Catalog) bool {
	sql := fmt.Sprintf("select id from catalog where id=%d", catalog.Id)
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
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
		sql = fmt.Sprintf("insert into catalog (name,creater,pid) values ('%s',%d,%d)", catalog.Name, catalog.Creater, catalog.Pid)
	} else {
		// modify
		sql = fmt.Sprintf("update catalog set name ='%s', creater =%d, pid=%d where id=%d", catalog.Name, catalog.Creater, catalog.Pid, catalog.Id)
	}
	
	result = model.Execute(sql)
	
	return result
}



func (this *Catalog)Query(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id,name,creater, pid from catalog where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
		return false
	}

	result := false
	for model.Next() {
		model.GetValue(&this.Id, &this.Name, &this.Creater, &this.Pid)
		result = true
	}
		
	return result	
}


func (this *Catalog)delete(model modelhelper.Model) {
	sql := fmt.Sprintf("delete from catalog where id =%d", this.Id)
	if !model.Execute(sql) {
		log.Printf("delete catalog failed, sql:%s", sql)
	}
}

func (this *Catalog)save(model modelhelper.Model) bool {
	sql := fmt.Sprintf("select id from catalog where id=%d", this.Id)
	if !model.Query(sql) {
		log.Printf("query catalog failed, sql:%s", sql)
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
		sql = fmt.Sprintf("insert into catalog (name,creater,pid) values ('%s',%d,%d)", this.Name, this.Creater, this.Pid)
	} else {
		// modify
		sql = fmt.Sprintf("update catalog set name ='%s', creater =%d, pid=%d where id=%d", this.Name, this.Creater, this.Pid, this.Id)
	}
	
	result = model.Execute(sql)
	
	return result
}


