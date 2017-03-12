package dal

import (
	"fmt"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/common/resource"
)

// QueryAllLink 查询全部Link
func QueryAllLink(helper dbhelper.DBHelper) []model.Link {
	linkList := []model.Link{}
	sql := fmt.Sprintf(`select id, name, url, logo, creater from link`)
	helper.Query(sql)

	for helper.Next() {
		link := model.Link{}
		helper.GetValue(&link.ID, &link.Name, &link.URL, &link.Logo, &link.Creater)

		linkList = append(linkList, link)
	}

	for index, _ := range linkList {
		link := &linkList[index]
		ress := resource.QueryRelativeResource(helper, link.ID, model.LINK)
		for _, r := range ress {
			link.Catalog = append(link.Catalog, r.RId())
		}
	}

	return linkList
}

// QueryLinkByCatalog 查询指定分类下的Link
func QueryLinkByCatalog(helper dbhelper.DBHelper, id int) []model.Link {
	linkList := []model.Link{}

	resList := resource.QueryReferenceResource(helper, id, model.CATALOG, model.LINK)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name, url, logo, creater from link where id =%d`, r.RId())
		helper.Query(sql)

		if helper.Next() {
			link := model.Link{}
			helper.GetValue(&link.ID, &link.Name, &link.URL, &link.Logo, &link.Creater)
			linkList = append(linkList, link)
		}
	}

	for _, link := range linkList {
		ress := resource.QueryRelativeResource(helper, link.ID, model.LINK)
		for _, r := range ress {
			link.Catalog = append(link.Catalog, r.RId())
		}
	}

	return linkList
}

// QueryLinkByRang 查询指定范围的Link
func QueryLinkByRang(helper dbhelper.DBHelper, begin int, offset int) []model.Link {
	linkList := []model.Link{}
	sql := fmt.Sprintf(`select id, name, url, logo, creater from link order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		link := model.Link{}
		helper.GetValue(&link.ID, &link.Name, &link.URL, &link.Logo, &link.Creater)

		linkList = append(linkList, link)
	}

	for _, link := range linkList {
		ress := resource.QueryRelativeResource(helper, link.ID, model.LINK)
		for _, r := range ress {
			link.Catalog = append(link.Catalog, r.RId())
		}
	}

	return linkList
}

// QueryLinkByID 查询指定Link
func QueryLinkByID(helper dbhelper.DBHelper, id int) (model.Link, bool) {
	link := model.Link{}
	sql := fmt.Sprintf(`select id, name, url, logo, creater from link where id =%d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&link.ID, &link.Name, &link.URL, &link.Logo, &link.Creater)
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
	sql := fmt.Sprintf(`delete from link where id =%d`, id)
	num, result := helper.Execute(sql)
	if num > 0 && result {
		lnk := resource.CreateSimpleRes(id, model.LINK, "")
		result = resource.DeleteResource(helper, lnk)
	}

	return result
}

// CreateLink 新建Link
func CreateLink(helper dbhelper.DBHelper, name, url, logo string, uID int, catalogs []int) (model.Link, bool) {
	lnk := model.Link{}
	lnk.Name = name
	lnk.URL = url
	lnk.Logo = logo
	lnk.Catalog = catalogs
	lnk.Creater = uID
	// insert
	sql := fmt.Sprintf(`insert into link (name,url,logo,creater) values ('%s','%s','%s', %d)`, lnk.Name, lnk.URL, lnk.Logo, lnk.Creater)
	_, result := helper.Execute(sql)
	if result {
		sql = fmt.Sprintf(`select id from link where name='%s' and url ='%s' and creater=%d`, lnk.Name, lnk.URL, lnk.Creater)

		helper.Query(sql)
		if helper.Next() {
			helper.GetValue(&lnk.ID)
		}
	}

	if result {
		res := resource.CreateSimpleRes(lnk.ID, model.LINK, lnk.Name)
		for _, c := range lnk.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return lnk, result
}

// SaveLink 保存Link
func SaveLink(helper dbhelper.DBHelper, lnk model.Link) (model.Link, bool) {
	// modify
	sql := fmt.Sprintf(`update link set name ='%s', url ='%s', logo='%s', creater=%d where id=%d`, lnk.Name, lnk.URL, lnk.Logo, lnk.Creater, lnk.ID)
	num, result := helper.Execute(sql)

	if result && num == 1 {
		res := resource.CreateSimpleRes(lnk.ID, model.LINK, lnk.Name)
		for _, c := range lnk.Catalog {
			ca := resource.CreateSimpleRes(c, model.CATALOG, "")
			res.AppendRelative(ca)
		}
		result = resource.SaveResource(helper, res)
	}

	return lnk, result
}
