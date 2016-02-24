package catalog


import (
	"fmt"
	"webcenter/util/modelhelper"
	"webcenter/kernel/admin/common"
	"webcenter/kernel/admin/content/base"
)

type CatalogInfo struct {
	Id int
	Name string
	Creater string
	Parent []string
}


type Catalog interface {
	common.Resource
	Creater() int
	SetId(id int)
	SetName(name string)
	SetCreater(user int)
	SetParent(pid []int)	
}

type catalog struct {
	id int
	name string
	creater int
	parent []int
}

func (this *catalog) Id() int {
	return this.id
}

func (this *catalog) Name() string {
	return this.name
}

func (this *catalog) Type() int {
	return base.CATALOG
}

func (this *catalog)Relative() []common.Resource {
	ress := []common.Resource{}
	
	for _, pid := range this.parent {
		res := common.NewSimpleRes(pid,"", base.CATALOG)
		ress = append(ress, res)
	}
	
	return ress
}

func (this *catalog) Creater() int {
	return this.creater
}

func (this *catalog) SetId(id int) {
	this.id = id
}

func (this *catalog) SetName(name string) {
	this.name = name
}

func (this *catalog) SetCreater(user int) {
	this.creater = user
}

func (this *catalog) SetParent(pid []int) {
	this.parent = pid
}

func NewCatalog() Catalog {
	c := &catalog{}	
	c.id = -1
	
	return c
}

func QueryAllCatalogInfo(model modelhelper.Model) []CatalogInfo {
	catalogInfoList := []CatalogInfo{}
		
	sql := fmt.Sprintf(`select c.id, c.name, u.nickname from catalog c, user u where c.creater = u.id`)
	model.Query(sql)

	for model.Next() {
		info := CatalogInfo{}
		info.Parent = []string{}
		model.GetValue(&info.Id, &info.Name, &info.Creater)
		
		catalogInfoList = append(catalogInfoList, info)
	}
	
	for index, info := range catalogInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.CATALOG)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			catalogInfoList[index].Parent = append(catalogInfoList[index].Parent, name)
		}
	}
	
	return catalogInfoList
}

func QueryCatalogInfoById(model modelhelper.Model, id int) (CatalogInfo, bool) {
	catalog := CatalogInfo{}
	sql := fmt.Sprintf(`select c.id, c.name, u.nickname from catalog c, user u where c.creater = u.id and c.id = %d`, id)
	model.Query(sql)

	result := false
	for model.Next() {
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater)
		result = true
	}
	if result {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, id, base.CATALOG)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			catalog.Parent = append(catalog.Parent, name)		
		}
	}
	
	return catalog, result	
}

func QueryAvalibleParentCatalogInfo(model modelhelper.Model, id int) []CatalogInfo {
	catalogInfoList := []CatalogInfo{}
	sql := fmt.Sprintf(`select c.id, c.name, u.nickname from catalog c, user u where c.creater = u.id and c.id < %d`, id)
	model.Query(sql)

	for model.Next() {
		catalog := CatalogInfo{}
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater)
		
		catalogInfoList = append(catalogInfoList, catalog)
	}
	
	for index, info := range catalogInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.CATALOG)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			catalogInfoList[index].Parent = append(catalogInfoList[index].Parent, name)
		}
	}
		
	return catalogInfoList
}

func QuerySubCatalogInfo(model modelhelper.Model, id int) []CatalogInfo {
	catalogInfoList := []CatalogInfo{}
	sql := fmt.Sprintf(`select distinct c.id id, c.name name, u.nickname from catalog c, user u ,resource_relative rr where c.creater = u.id and c.id in (
	select src from resource_relative where dst = %d and dstType = %d and srcType = %d
	)`, id, base.CATALOG, base.CATALOG)
	model.Query(sql)

	for model.Next() {
		catalog := CatalogInfo{}
		model.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater)
		
		catalogInfoList = append(catalogInfoList, catalog)
	}
	
	for index, info := range catalogInfoList {
		sql = fmt.Sprintf(`select r.name from resource r, resource_relative rr where r.id = rr.dst and r.type = rr.dstType and rr.src = %d and rr.srcType=%d`, info.Id, base.CATALOG)
		name := "-"
		model.Query(sql)
		for model.Next() {
			model.GetValue(&name)
			catalogInfoList[index].Parent = append(catalogInfoList[index].Parent, name)
		}
	}
		
	return catalogInfoList
}

func QueryCatalogById(model modelhelper.Model, id int) (Catalog, bool) {
	catalog := NewCatalog()
	sql := fmt.Sprintf(`select name, creater from catalog where id = %d`, id)
	model.Query(sql)

	name := ""
	creater := -1
	result := false
	for model.Next() {
		model.GetValue(&name, &creater)
		result = true
	}
	
	if !result {
		return catalog, result
	}
	
	catalog.SetId(id)
	catalog.SetName(name)
	catalog.SetCreater(creater)

	pid := []int{}
	if result {
		sql = fmt.Sprintf(`select dst resource_relative where src = %d and srcType=%d and dstType=%d`, id, base.CATALOG, base.CATALOG)
		dst := -1
		model.Query(sql)
		for model.Next() {
			model.GetValue(&dst)
			pid = append(pid, dst)			
		}
	}
	catalog.SetParent(pid)
		
	return catalog, result	
}

func DeleteCatalogById(model modelhelper.Model, id int) bool {
	model.BeginTransaction()
			
	sql := fmt.Sprintf(`delete from catalog where id=%d`, id)
	result := model.Execute(sql)
	if result {
		ca := catalog{}
		ca.id = id
		result = common.DeleteResource(model, &ca)
	}
	
	if !result {
		model.Rollback()
	} else {
		model.Commit()
	}	
	
	return result	
}

func SaveCatalog(model modelhelper.Model, catalog Catalog) bool {
	sql := fmt.Sprintf(`select id from catalog where id=%d`, catalog.Id())
	model.Query(sql)

	result := false;
	if model.Next() {
		var id = 0
		model.GetValue(&id)
		result = true
	}

	model.BeginTransaction()
	
	if !result {
		// insert
		sql = fmt.Sprintf(`insert into catalog (name,creater) values ('%s',%d)`, catalog.Name(), catalog.Creater())
		result = model.Execute(sql)
		
		if result {
			id := -1
			sql = fmt.Sprintf(`select id from catalog where name='%s' and creater=%d`, catalog.Name(), catalog.Creater())
			model.Query(sql)
			if model.Next() {
				model.GetValue(&id)
			}
			
			catalog.SetId(id)
		}			
	} else {
		// modify
		sql = fmt.Sprintf(`update catalog set name ='%s', creater =%d where id=%d`, catalog.Name(), catalog.Creater(), catalog.Id())
		result = model.Execute(sql)
	}
	
	if result {
		result = common.SaveResource(model, catalog)
	}
	
	if !result {
		model.Rollback()
	} else {
		model.Commit()
	}
		
	return result
}






