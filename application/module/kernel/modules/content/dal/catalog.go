package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/application/common"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/resource"
	"muidea.com/magicCenter/foundation/util"
	"muidea.com/magicCommon/model"
)

func loadCatalogID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from content_catalog`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// QueryAllCatalog 查询所有分类
func QueryAllCatalog(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}

	sql := fmt.Sprintf(`select id, name,createdate,creater from content_catalog`)
	helper.Query(sql)

	for helper.Next() {
		c := model.Summary{Type: model.CATALOG}
		helper.GetValue(&c.ID, &c.Name, &c.CreateDate, &c.Creater)

		summaryList = append(summaryList, c)
	}
	helper.Finish()

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.CATALOG)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryCatalogs 查询指定分类
func QueryCatalogs(helper dbhelper.DBHelper, ids []int) []model.Catalog {
	catalogList := []model.Catalog{}

	if len(ids) == 0 {
		return catalogList
	}

	sql := fmt.Sprintf(`select id, name from content_catalog where id in(%s)`, util.IntArray2Str(ids))
	helper.Query(sql)
	defer helper.Finish()

	for helper.Next() {
		summary := model.Catalog{}
		helper.GetValue(&summary.ID, &summary.Name)

		catalogList = append(catalogList, summary)
	}

	return catalogList
}

// QueryCatalogByID 查询指定ID的Catalog
func QueryCatalogByID(helper dbhelper.DBHelper, id int) (model.CatalogDetail, bool) {
	catalog := model.CatalogDetail{}
	sql := fmt.Sprintf(`select id, name, description, createdate, creater from content_catalog where id = %d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&catalog.ID, &catalog.Name, &catalog.Description, &catalog.CreateDate, &catalog.Creater)
		result = true
	}
	helper.Finish()

	if result {
		ress := resource.QueryRelativeResource(helper, id, model.CATALOG)

		for _, r := range ress {
			catalog.Catalog = append(catalog.Catalog, r.RId())
		}
	}
	return catalog, result
}

// QueryCatalogByName 查询指定Name的Catalog
func QueryCatalogByName(helper dbhelper.DBHelper, name string) (model.CatalogDetail, bool) {
	catalog := model.CatalogDetail{}
	sql := fmt.Sprintf(`select id, name, description, createdate, creater from content_catalog where name = '%s'`, name)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&catalog.ID, &catalog.Name, &catalog.Description, &catalog.CreateDate, &catalog.Creater)
		result = true
	}
	helper.Finish()

	if result {
		ress := resource.QueryRelativeResource(helper, catalog.ID, model.CATALOG)

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
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
		summaryList = append(summaryList, summary)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, value.Type)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// DeleteCatalog 删除指定类
func DeleteCatalog(helper dbhelper.DBHelper, id int) bool {
	result := false
	helper.BeginTransaction()

	for {
		sql := fmt.Sprintf(`delete from content_catalog where id=%d`, id)

		_, result = helper.Execute(sql)
		if result {
			res, ok := resource.QueryResource(helper, id, model.CATALOG)
			if ok {
				result = resource.DeleteResource(helper, res, true)
			} else {
				result = ok
			}
		}

		break
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return result
}

// UpdateCatalog 更新Catalog
func UpdateCatalog(helper dbhelper.DBHelper, catalogs []model.Catalog, updateDate string, updater int) ([]model.Catalog, bool) {
	ids := []int{}
	result := false
	if len(catalogs) > 0 {
		helper.BeginTransaction()
		for _, val := range catalogs {
			result = true
			detail, existFlag := QueryCatalogByName(helper, val.Name)

			if existFlag {
				modifyFlag := false
				if detail.Creater != updater {
					detail.Creater = updater
					modifyFlag = true
				}

				if modifyFlag {
					detail.CreateDate = updateDate
					_, result = SaveCatalog(helper, detail, true)
					if !result {
						break
					}
				}

				ids = append(ids, detail.ID)
			} else {
				detail, ok := CreateCatalog(helper, val.Name, "", updateDate, []int{common.DefaultContentCatalog.ID}, updater, true)
				if ok {
					ids = append(ids, detail.ID)
				} else {
					result = false
				}
			}
		}

		if result {
			helper.Commit()
			return QueryCatalogs(helper, ids), true
		}

		helper.Rollback()
		return []model.Catalog{}, false
	}

	return []model.Catalog{}, true
}

// CreateCatalog 新建分类
func CreateCatalog(helper dbhelper.DBHelper, name, description, createDate string, parent []int, creater int, enableTransaction bool) (model.Summary, bool) {
	catalog := model.Summary{Unit: model.Unit{Name: name}, Type: model.CATALOG, Catalog: parent, CreateDate: createDate, Creater: creater}

	if !enableTransaction {
		helper.BeginTransaction()
	}

	id := allocCatalogID()
	result := false
	for {
		sql := fmt.Sprintf(`select id from content_catalog where name='%s'`, name)
		helper.Query(sql)
		if helper.Next() {
			// 说明对应的Catalog已经存在，返回Create失败
			helper.Finish()
			break
		}
		helper.Finish()

		// insert
		sql = fmt.Sprintf(`insert into content_catalog (id, name, description, createdate, creater) values (%d, '%s','%s','%s',%d)`, id, name, description, createDate, creater)
		_, result = helper.Execute(sql)
		if !result {
			break
		}

		catalog.ID = id
		res := resource.CreateSimpleRes(catalog.ID, model.CATALOG, catalog.Name, catalog.CreateDate, catalog.Creater)
		for _, c := range parent {
			ca, ok := resource.QueryResource(helper, c, model.CATALOG)
			if ok {
				res.AppendRelative(ca)
			} else {
				result = false
				break
			}
		}

		if result {
			result = resource.CreateResource(helper, res, true)
		}

		break
	}

	if !enableTransaction {
		if result {
			helper.Commit()
		} else {
			helper.Rollback()
		}
	}

	return catalog, result
}

// SaveCatalog 保存分类
func SaveCatalog(helper dbhelper.DBHelper, catalog model.CatalogDetail, enableTransaction bool) (model.Summary, bool) {
	if !enableTransaction {
		helper.BeginTransaction()
	}
	summary := model.Summary{Unit: model.Unit{ID: catalog.ID, Name: catalog.Name}, Type: model.CATALOG, Catalog: catalog.Catalog, CreateDate: catalog.CreateDate, Creater: catalog.Creater}

	result := false
	for {
		// modify
		sql := fmt.Sprintf(`update content_catalog set description='%s', createdate='%s', creater =%d where id=%d`, catalog.Description, catalog.CreateDate, catalog.Creater, catalog.ID)
		_, result = helper.Execute(sql)

		if result {
			res, ok := resource.QueryResource(helper, catalog.ID, model.CATALOG)
			if !ok {
				result = false
				break
			}

			res.ResetRelative()
			for _, c := range catalog.Catalog {
				ca, ok := resource.QueryResource(helper, c, model.CATALOG)
				if ok {
					res.AppendRelative(ca)
				}
			}
			result = resource.SaveResource(helper, res, true)
		}

		break
	}

	if !enableTransaction {
		if result {
			helper.Commit()
		} else {
			helper.Rollback()
		}
	}

	return summary, result
}
