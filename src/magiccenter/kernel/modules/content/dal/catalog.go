package dal

import (
	"fmt"
	"magiccenter/common/model"
	resdal "magiccenter/resource/dal"
	"magiccenter/util/dbhelper"
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

	for _, c := range catalogList {
		ress := resdal.QueryRelativeResource(helper, c.ID, model.CATALOG)
		for _, r := range ress {
			c.Parent = append(c.Parent, r.RId())
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
		ress := resdal.QueryRelativeResource(helper, id, model.CATALOG)

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

	resList := resdal.QueryReferenceResource(helper, id, model.CATALOG, model.CATALOG)
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
		ca := resdal.CreateSimpleRes(id, model.CATALOG, "")
		result = resdal.DeleteResource(helper, ca)
	}

	return result
}

// SaveCatalog 保存分类
func SaveCatalog(helper dbhelper.DBHelper, catalog model.CatalogDetail) bool {
	sql := fmt.Sprintf(`select id from catalog where id=%d`, catalog.ID)
	helper.Query(sql)

	result := false
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into catalog (name,creater) values ('%s',%d)`, catalog.Name, catalog.Creater)
		_, result = helper.Execute(sql)

		if result {
			sql = fmt.Sprintf(`select id from catalog where name='%s' and creater=%d`, catalog.Name, catalog.Creater)
			helper.Query(sql)
			if helper.Next() {
				helper.GetValue(&catalog.ID)
			}
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update catalog set name ='%s', creater =%d where id=%d`, catalog.Name, catalog.Creater, catalog.ID)
		_, result = helper.Execute(sql)
	}

	if result {
		res := resdal.CreateSimpleRes(catalog.ID, model.CATALOG, catalog.Name)
		for _, c := range catalog.Parent {
			ca := resdal.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resdal.SaveResource(helper, res)
	}

	return result
}
