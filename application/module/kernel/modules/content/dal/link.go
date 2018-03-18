package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
	"muidea.com/magicCenter/foundation/util"
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

// QueryLinks 查询指定链接
func QueryLinks(helper dbhelper.DBHelper, ids []int) []model.Link {
	linkList := []model.Link{}

	if len(ids) == 0 {
		return linkList
	}

	sql := fmt.Sprintf(`select id, name from content_link where id in(%s)`, util.IntArray2Str(ids))
	helper.Query(sql)

	for helper.Next() {
		link := model.Link{}
		helper.GetValue(&link.ID, &link.Name)

		linkList = append(linkList, link)
	}

	return linkList
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
	result := false
	helper.BeginTransaction()

	for {
		sql := fmt.Sprintf(`delete from content_link where id =%d`, id)
		_, result = helper.Execute(sql)
		if result {
			res, ok := resource.QueryResource(helper, id, model.LINK)
			if ok {
				result = resource.DeleteResource(helper, res, true)
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

// CreateLink 新建Link
func CreateLink(helper dbhelper.DBHelper, name, url, logo, createDate string, uID int, catalogs []int) (model.Summary, bool) {
	lnk := model.Summary{}
	lnk.Name = name
	lnk.Catalog = catalogs
	lnk.CreateDate = createDate
	lnk.Creater = uID
	result := false
	helper.BeginTransaction()

	for {
		// insert
		sql := fmt.Sprintf(`insert into content_link (name,url,logo, createDate, creater) values ('%s','%s','%s','%s', %d)`, name, url, logo, createDate, uID)
		_, result = helper.Execute(sql)
		if result {
			sql = fmt.Sprintf(`select id from content_link where name='%s' and url ='%s' and creater=%d`, name, url, uID)

			helper.Query(sql)
			if helper.Next() {
				helper.GetValue(&lnk.ID)
			} else {
				result = false
				break
			}
		}

		if result {
			res := resource.CreateSimpleRes(lnk.ID, model.LINK, lnk.Name, lnk.CreateDate, lnk.Creater)
			for _, c := range lnk.Catalog {
				ca, ok := resource.QueryResource(helper, c, model.CATALOG)
				if ok {
					res.AppendRelative(ca)
				} else {
					result = false
					break
				}
			}

			result = resource.CreateResource(helper, res, true)
		}

		break
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return lnk, result
}

// SaveLink 保存Link
func SaveLink(helper dbhelper.DBHelper, lnk model.LinkDetail) (model.Summary, bool) {
	summary := model.Summary{Unit: model.Unit{ID: lnk.ID, Name: lnk.Name}, Catalog: lnk.Catalog, CreateDate: lnk.CreateDate, Creater: lnk.Creater}
	result := false
	helper.BeginTransaction()

	for {
		// modify
		sql := fmt.Sprintf(`update content_link set name ='%s', url ='%s', logo='%s', createdate='%s', creater=%d where id=%d`, lnk.Name, lnk.URL, lnk.Logo, lnk.CreateDate, lnk.Creater, lnk.ID)
		_, result = helper.Execute(sql)

		if result {
			res, ok := resource.QueryResource(helper, lnk.ID, model.LINK)
			if !ok {
				result = false
				break
			}

			res.ResetRelative()
			for _, c := range lnk.Catalog {
				ca, ok := resource.QueryResource(helper, c, model.CATALOG)
				if ok {
					res.AppendRelative(ca)
				} else {
					result = false
					break
				}
			}
			if result {
				result = resource.SaveResource(helper, res, true)
			}
		}

		break
	}

	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}

	return summary, result
}
