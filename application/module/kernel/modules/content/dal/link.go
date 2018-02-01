package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllLink 查询全部Link
func QueryAllLink(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}
	sql := fmt.Sprintf(`select id, name,createdate,creater from content_link`)
	helper.Query(sql)

	for helper.Next() {
		link := model.Summary{}
		helper.GetValue(&link.ID, &link.Name, &link.CreateDate, &link.Creater)

		summaryList = append(summaryList, link)
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.LINK)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryLinkByCatalog 查询指定分类下的Link
func QueryLinkByCatalog(helper dbhelper.DBHelper, id int) []model.Summary {
	summaryList := []model.Summary{}

	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.LINK)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name,createdate,creater from content_link where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			link := model.Summary{}
			helper.GetValue(&link.ID, &link.Name, &link.CreateDate, &link.Creater)
			summaryList = append(summaryList, link)
		}
	}

	for index, value := range summaryList {
		summary := &summaryList[index]
		ress := resource.QueryRelativeResource(helper, value.ID, model.LINK)
		for _, r := range ress {
			summary.Catalog = append(summary.Catalog, r.RId())
		}
	}

	return summaryList
}

// QueryLinkByID 查询指定Link
func QueryLinkByID(helper dbhelper.DBHelper, id int) (model.LinkDetail, bool) {
	link := model.LinkDetail{}
	sql := fmt.Sprintf(`select id, name, url, logo, createdate, creater from content_link where id =%d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&link.ID, &link.Name, &link.URL, &link.Logo, &link.CreateDate, &link.Creater)
		result = true
	}

	if result {
		ress := resource.QueryRelativeResource(helper, link.ID, model.LINK)
		for _, r := range ress {
			link.Catalog = append(link.Catalog, r.RId())
		}
	}

	return link, result
}

// DeleteLinkByID 删除指定Link
func DeleteLinkByID(helper dbhelper.DBHelper, id int) bool {
	sql := fmt.Sprintf(`delete from content_link where id =%d`, id)
	num, result := helper.Execute(sql)
	if num > 0 && result {
		lnk := resource.CreateSimpleRes(id, model.LINK, "", "", -1)
		result = resource.DeleteResource(helper, lnk)
	}

	return result
}

// CreateLink 新建Link
func CreateLink(helper dbhelper.DBHelper, name, url, logo, createDate string, uID int, catalogs []int) (model.Summary, bool) {
	lnk := model.Summary{}
	lnk.Name = name
	lnk.Catalog = catalogs
	lnk.CreateDate = createDate
	lnk.Creater = uID

	// insert
	sql := fmt.Sprintf(`insert into content_link (name,url,logo, createDate, creater) values ('%s','%s','%s','%s', %d)`, name, url, logo, createDate, uID)
	_, result := helper.Execute(sql)
	if result {
		sql = fmt.Sprintf(`select id from content_link where name='%s' and url ='%s' and creater=%d`, name, url, uID)

		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&lnk.ID)
		}
	}

	if result {
		res := resource.CreateSimpleRes(lnk.ID, model.LINK, lnk.Name, lnk.CreateDate, lnk.Creater)
		for _, c := range lnk.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "", "", -1)
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return lnk, result
}

// SaveLink 保存Link
func SaveLink(helper dbhelper.DBHelper, lnk model.LinkDetail) (model.Summary, bool) {
	// modify
	sql := fmt.Sprintf(`update content_link set name ='%s', url ='%s', logo='%s', createdate='%s', creater=%d where id=%d`, lnk.Name, lnk.URL, lnk.Logo, lnk.CreateDate, lnk.Creater, lnk.ID)
	num, result := helper.Execute(sql)

	if result && num == 1 {
		res := resource.CreateSimpleRes(lnk.ID, model.LINK, lnk.Name, lnk.CreateDate, lnk.Creater)
		for _, c := range lnk.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "", "", -1)
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	} else {
		result = false
	}

	return model.Summary{ID: lnk.ID, Name: lnk.Name, Catalog: lnk.Catalog, CreateDate: lnk.CreateDate, Creater: lnk.Creater}, result
}
