package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllCatalog 查询所有分类
func QueryAllCatalog(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}

	sql := fmt.Sprintf(`select id, name,createdate,creater from catalog`)
	helper.Query(sql)

	for helper.Next() {
		c := model.Summary{}
		helper.GetValue(&c.ID, &c.Name, &c.CreateDate, &c.Creater)

		summaryList = append(summaryList, c)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.CATALOG)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryCatalogByID 查询指定分类
func QueryCatalogByID(helper dbhelper.DBHelper, id int) (model.CatalogDetail, bool) {
	catalog := model.CatalogDetail{}
	sql := fmt.Sprintf(`select id, name, description, createdate, creater from catalog where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&catalog.ID, &catalog.Name, &catalog.Description, &catalog.CreateDate, &catalog.Creater)
		result = true
	}

	if result {
		ress := resource.QueryRelativeResource(helper, id, model.CATALOG)

		for _, r := range ress {
			catalog.Catalog = append(catalog.Catalog, r.RId())
		}
	}
	return catalog, result
}

// QueryCatalogByCatalog 查询指定分类的子类
func QueryCatalogByCatalog(helper dbhelper.DBHelper, id int) []model.Summary {
	summaryList := []model.Summary{}

	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.CATALOG)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name,createdate,creater from catalog where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			summary := model.Summary{}
			helper.GetValue(&summary.ID, &summary.Name, &summary.CreateDate, &summary.Creater)

			summaryList = append(summaryList, summary)
		}
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.CATALOG)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// DeleteCatalog 删除指定类
func DeleteCatalog(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from catalog where id=%d`, id)

	num, result := helper.Execute(sql)
	if num >= 1 && result {
		ca := resource.CreateSimpleRes(id, model.CATALOG, "")
		result = resource.DeleteResource(helper, ca)
	}

	return result
}

// CreateCatalog 新建分类
func CreateCatalog(helper dbhelper.DBHelper, name, description, createdate string, parent []int, creater int) (model.Summary, bool) {
	catalog := model.Summary{}
	catalog.Name = name
	catalog.Creater = creater
	catalog.Catalog = parent
	catalog.CreateDate = createdate

	// insert
	sql := fmt.Sprintf(`insert into catalog (name, description, createdate, creater) values ('%s','%s','%s',%d)`, name, description, createdate, creater)
	num, result := helper.Execute(sql)

	if num == 1 && result {
		sql = fmt.Sprintf(`select id from catalog where name='%s' and creater=%d`, name, creater)
		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&catalog.ID)
		}
	}

	if result {
		res := resource.CreateSimpleRes(catalog.ID, model.CATALOG, catalog.Name)
		for _, c := range parent {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return catalog, result
}

// SaveCatalog 保存分类
func SaveCatalog(helper dbhelper.DBHelper, catalog model.CatalogDetail) (model.Summary, bool) {
	// modify
	sql := fmt.Sprintf(`update catalog set name ='%s', description='%s', createdate='%s', creater =%d where id=%d`, catalog.Name, catalog.Description, catalog.CreateDate, catalog.Creater, catalog.ID)
	num, result := helper.Execute(sql)

	if num == 1 && result {
		res := resource.CreateSimpleRes(catalog.ID, model.CATALOG, catalog.Name)
		for _, c := range catalog.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	} else {
		result = false
	}

	return model.Summary{ID: catalog.ID, Name: catalog.Name, Catalog: catalog.Catalog, CreateDate: catalog.CreateDate, Creater: catalog.Creater}, result
}
