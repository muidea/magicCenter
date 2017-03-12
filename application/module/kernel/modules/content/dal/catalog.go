package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllCatalog 查询所有分类
func QueryAllCatalog(helper dbhelper.DBHelper) []model.Catalog {
	catalogList := []model.Catalog{}

	sql := fmt.Sprintf(`select id, name from catalog`)
	helper.Query(sql)

	for helper.Next() {
		c := model.Catalog{}
		helper.GetValue(&c.ID, &c.Name)

		catalogList = append(catalogList, c)
	}

	return catalogList
}

// QueryAllCatalogDetail 查询所有分类详情
func QueryAllCatalogDetail(helper dbhelper.DBHelper) []model.CatalogDetail {
	catalogList := []model.CatalogDetail{}

	sql := fmt.Sprintf(`select id, name, creater from catalog`)
	helper.Query(sql)

	for helper.Next() {
		c := model.CatalogDetail{}
		helper.GetValue(&c.ID, &c.Name, &c.Creater)

		catalogList = append(catalogList, c)
	}

	for index, value := range catalogList {
		cur := &catalogList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.CATALOG)
		for _, r := range ress {
			cur.Parent = append(cur.Parent, r.RId())
		}
	}

	return catalogList
}

// QueryCatalogByID 查询指定分类
func QueryCatalogByID(helper dbhelper.DBHelper, id int) (model.CatalogDetail, bool) {
	catalog := model.CatalogDetail{}
	sql := fmt.Sprintf(`select id, name, creater from catalog where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&catalog.ID, &catalog.Name, &catalog.Creater)
		result = true
	}

	if result {
		ress := resource.QueryRelativeResource(helper, id, model.CATALOG)

		for _, r := range ress {
			catalog.Parent = append(catalog.Parent, r.RId())
		}
	}

	return catalog, result
}

// QueryAvalibleParentCatalog 查询可用的父类
// 可用父类的判断规则是比指定分类ID小的就视为可用分类
func QueryAvalibleParentCatalog(helper dbhelper.DBHelper, id int) []model.Catalog {
	catalogList := []model.Catalog{}
	sql := fmt.Sprintf(`select id, name from catalog where id < %d`, id)
	helper.Query(sql)

	for helper.Next() {
		catalog := model.Catalog{}
		helper.GetValue(&catalog.ID, &catalog.Name)

		catalogList = append(catalogList, catalog)
	}

	return catalogList
}

// QuerySubCatalog 查询指定分类的子类
func QuerySubCatalog(helper dbhelper.DBHelper, id int) []model.Catalog {
	catalogList := []model.Catalog{}

	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.CATALOG)
	for _, r := range resList {
		catalog := model.Catalog{}
		catalog.ID = r.RId()
		catalog.Name = r.RName()

		catalogList = append(catalogList, catalog)
	}

	return catalogList
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
func CreateCatalog(helper dbhelper.DBHelper, name string, parent []int, creater int) (model.Catalog, bool) {
	catalog := model.Catalog{}
	catalog.Name = name

	// insert
	sql := fmt.Sprintf(`insert into catalog (name,creater) values ('%s',%d)`, name, creater)
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
func SaveCatalog(helper dbhelper.DBHelper, catalog model.CatalogDetail) (model.Catalog, bool) {
	sql := fmt.Sprintf(`select id from catalog where id=%d`, catalog.ID)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if result {
		// modify
		sql = fmt.Sprintf(`update catalog set name ='%s', creater =%d where id=%d`, catalog.Name, catalog.Creater, catalog.ID)
		_, result = helper.Execute(sql)
	}

	if result {
		res := resource.CreateSimpleRes(catalog.ID, model.CATALOG, catalog.Name)
		for _, c := range catalog.Parent {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return model.Catalog{ID: catalog.ID, Name: catalog.Name}, result
}
