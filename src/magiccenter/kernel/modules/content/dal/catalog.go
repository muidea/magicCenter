package dal


import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/content/model"
	"magiccenter/kernel/account/dal"
)


func QueryAllCatalog(helper modelhelper.Model) []model.Catalog {
	catalogList := []model.Catalog{}
	
	sql := fmt.Sprintf(`select id, name from catalog`)
	helper.Query(sql)

	for helper.Next() {
		c := model.Catalog{}
		helper.GetValue(&c.Id, &c.Name)
		
		catalogList = append(catalogList, c)
	}
		
	return catalogList
}


func QueryAllCatalogDetail(helper modelhelper.Model) []model.CatalogDetail {
	catalogList := []model.CatalogDetail{}
	
	sql := fmt.Sprintf(`select id, name, creater from catalog`)
	helper.Query(sql)

	for helper.Next() {
		c := model.CatalogDetail{}
		helper.GetValue(&c.Id, &c.Name, &c.Creater.Id)
		
		catalogList = append(catalogList, c)
	}
	
	for i, _ := range catalogList {
		c := &catalogList[i]
		
		user, found := dal.QueryUserById(helper, c.Creater.Id)
		if found {
			c.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, c.Id, model.CATALOG)		
		for _, r := range ress {
			ca := model.Catalog{}
			ca.Id = r.RId()
			ca.Name = r.RName()
			
			c.Parent = append(c.Parent, ca)
		}
	}
	
	return catalogList
}

func QueryCatalogById(helper modelhelper.Model, id int) (model.CatalogDetail, bool) {
	catalog := model.CatalogDetail{}
	sql := fmt.Sprintf(`select id, name, creater from catalog where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&catalog.Id, &catalog.Name, &catalog.Creater.Id)
		result = true
	}
	
	if result {
		user, found := dal.QueryUserById(helper, catalog.Creater.Id)
		if found {
			catalog.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, id, model.CATALOG)
		
		for _, r := range ress {
			ca := model.Catalog{}
			ca.Id = r.RId()
			ca.Name = r.RName()
			
			catalog.Parent = append(catalog.Parent, ca)
		}
	}
	
	return catalog, result	
}

func QueryAvalibleParentCatalog(helper modelhelper.Model, id int) []model.Catalog {
	catalogList := []model.Catalog{}
	sql := fmt.Sprintf(`select id, name from catalog where id < %d`, id)
	helper.Query(sql)

	for helper.Next() {
		catalog := model.Catalog{}
		helper.GetValue(&catalog.Id, &catalog.Name)
		
		catalogList = append(catalogList, catalog)
	}

	return catalogList
}

func QuerySubCatalog(helper modelhelper.Model, id int) []model.Catalog {
	catalogList := []model.Catalog{}
	
	resList := QueryReferenceResource(helper, id, model.CATALOG, model.CATALOG)
	for _, r := range resList {
		catalog := model.Catalog{}
		catalog.Id = r.RId()
		catalog.Name = r.RName()
		
		catalogList = append(catalogList, catalog)
	}

	return catalogList		
}

func DeleteCatalog(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf(`delete from catalog where id=%d`, id)
	
	num, result := helper.Execute(sql)
	if num > 1 && result {
		ca := model.Catalog{}
		ca.Id = id
		
		result = DeleteResource(helper, &ca)
	}
	
	return result	
}

func SaveCatalog(helper modelhelper.Model, catalog model.CatalogDetail) bool {
	sql := fmt.Sprintf(`select id from catalog where id=%d`, catalog.Id)
	helper.Query(sql)

	result := false;
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into catalog (name,creater) values ('%s',%d)`, catalog.Name, catalog.Creater.Id)
		_, result = helper.Execute(sql)
		
		if result {
			sql = fmt.Sprintf(`select id from catalog where name='%s' and creater=%d`, catalog.Name, catalog.Creater.Id)
			helper.Query(sql)
			if helper.Next() {
				helper.GetValue(&catalog.Id)
			}
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update catalog set name ='%s', creater =%d where id=%d`, catalog.Name, catalog.Creater.Id, catalog.Id)
		_, result = helper.Execute(sql)
	}
	
	if result {
		result = SaveResource(helper, &catalog)
	}
			
	return result
}






