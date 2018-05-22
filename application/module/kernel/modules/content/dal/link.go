package dal

import (
	"database/sql"
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/resource"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

func loadLinkID(helper dbhelper.DBHelper) int {
	var maxID sql.NullInt64
	sql := fmt.Sprintf(`select max(id) from content_link`)
	helper.Query(sql)
	defer helper.Finish()

	if helper.Next() {
		helper.GetValue(&maxID)
	}

	return int(maxID.Int64)
}

// QueryAllLink 查询全部Link
func QueryAllLink(helper dbhelper.DBHelper) []model.Summary {
	summaryList := []model.Summary{}

	ress := resource.QueryResourceByType(helper, model.LINK)
	for _, v := range ress {
		summary := model.Summary{Unit: model.Unit{ID: v.RId(), Name: v.RName()}, Description: v.RDescription(), Type: v.RType(), CreateDate: v.RCreateDate(), Creater: v.ROwner()}

		for _, r := range v.Relative() {
			summary.Catalog = append(summary.Catalog, r.RId())
		}

		summaryList = append(summaryList, summary)
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
	defer helper.Finish()

	for helper.Next() {
		link := model.Link{}
		helper.GetValue(&link.ID, &link.Name)

		linkList = append(linkList, link)
	}

	return linkList
}

// QueryLinkByCatalog 查询指定分类下的Link
func QueryLinkByCatalog(helper dbhelper.DBHelper, catalog int) []model.Summary {
	summaryList := []model.Summary{}

	resList := resource.QueryReferenceResource(helper, catalog, model.CATALOG, model.LINK)
	for _, r := range resList {
		summary := model.Summary{Unit: model.Unit{ID: r.RId(), Name: r.RName()}, Description: r.RDescription(), Type: r.RType(), CreateDate: r.RCreateDate(), Creater: r.ROwner()}
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

// QueryLinkByID 查询指定Link
func QueryLinkByID(helper dbhelper.DBHelper, id int) (model.LinkDetail, bool) {
	link := model.LinkDetail{}
	sql := fmt.Sprintf(`select id, name, description, url, logo, createdate, creater from content_link where id =%d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&link.ID, &link.Name, &link.Description, &link.URL, &link.Logo, &link.CreateDate, &link.Creater)
		result = true
	}
	helper.Finish()

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

// CreateLink 新建Link
func CreateLink(helper dbhelper.DBHelper, name, description, url, logo, createDate string, creater int, catalogs []int) (model.Summary, bool) {
	lnk := model.Summary{Unit: model.Unit{Name: name}, Description: description, Type: model.LINK, Catalog: catalogs, CreateDate: createDate, Creater: creater}

	id := allocLinkID()
	result := false
	helper.BeginTransaction()

	for {
		// insert
		sql := fmt.Sprintf(`insert into content_link (id, name, description,url,logo, createDate, creater) values (%d,'%s','%s','%s','%s','%s', %d)`, id, name, description, url, logo, createDate, creater)
		_, result = helper.Execute(sql)
		if !result {
			break
		}

		lnk.ID = id
		res := resource.CreateSimpleRes(lnk.ID, model.LINK, lnk.Name, lnk.Description, lnk.CreateDate, lnk.Creater)
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
	summary := model.Summary{Unit: model.Unit{ID: lnk.ID, Name: lnk.Name}, Description: lnk.Description, Type: model.LINK, Catalog: lnk.Catalog, CreateDate: lnk.CreateDate, Creater: lnk.Creater}
	result := false
	helper.BeginTransaction()

	for {
		// modify
		sql := fmt.Sprintf(`update content_link set name ='%s', description ='%s', url ='%s', logo='%s', createdate='%s', creater=%d where id=%d`, lnk.Name, lnk.Description, lnk.URL, lnk.Logo, lnk.CreateDate, lnk.Creater, lnk.ID)
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
