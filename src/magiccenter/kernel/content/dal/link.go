package dal

import (
	"fmt"
	"magiccenter/util/modelhelper"
	"magiccenter/kernel/content/model"
	"magiccenter/kernel/account/dal"
)

func QueryAllLink(helper modelhelper.Model) []model.LinkDetail {
	linkList := []model.LinkDetail{}
	sql := fmt.Sprintf(`select id, name, url, logo, creater from link`)
	helper.Query(sql)

	for helper.Next() {
		link := model.LinkDetail{}
		helper.GetValue(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Creater.Id)
		
		linkList = append(linkList, link)
	}
	
	for i, _ := range linkList {
		link := &linkList[i]
		user, found := dal.QueryUserById(helper, link.Creater.Id)
		if found {
			link.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, link.Id, model.LINK)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			link.Catalog = append(link.Catalog, catalog)
		}		
	}
		
	return linkList
}


func QueryLinkByCatalog(helper modelhelper.Model, id int) []model.LinkDetail {
	linkList := []model.LinkDetail{}
	
	resList := QueryReferenceResource(helper, id, model.CATALOG, model.LINK)
	for _, r := range resList {
		sql := fmt.Sprintf(`select id, name, url, logo, creater from link where id =%d`, r.RId())
		helper.Query(sql)
		
		if helper.Next() {
			link := model.LinkDetail{}
			helper.GetValue(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Creater.Id)			
			linkList = append(linkList, link)
		}
	}
	
	for i, _ := range linkList {
		link := &linkList[i]
		user, found := dal.QueryUserById(helper, link.Creater.Id)
		if found {
			link.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, link.Id, model.LINK)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			link.Catalog = append(link.Catalog, catalog)
		}		
	}
	
	return linkList	
}

func QueryLinkByRang(helper modelhelper.Model, begin int,offset int) []model.LinkDetail {
	linkList := []model.LinkDetail{}
	sql := fmt.Sprintf(`select id, name, url, logo, creater from link order by id where id >= %d limit %d`, begin, offset)
	helper.Query(sql)

	for helper.Next() {
		link := model.LinkDetail{}
		helper.GetValue(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Creater.Id)
		
		linkList = append(linkList, link)
	}
	
	for i, _ := range linkList {
		link := &linkList[i]
		user, found := dal.QueryUserById(helper, link.Creater.Id)
		if found {
			link.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, link.Id, model.LINK)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			link.Catalog = append(link.Catalog, catalog)
		}		
	}
			
	return linkList
}

func QueryLinkById(helper modelhelper.Model, id int) (model.LinkDetail, bool) {
	link := model.LinkDetail{}
	sql := fmt.Sprintf(`select id, name, url, logo, creater from link where id =%d`, id)
	helper.Query(sql)

	result := false
	if helper.Next() {
		helper.GetValue(&link.Id, &link.Name, &link.Url, &link.Logo, &link.Creater.Id)
		result = true
	}
	
	if result {
		user, found := dal.QueryUserById(helper, link.Creater.Id)
		if found {
			link.Creater.Name = user.Name
		}
		
		ress := QueryRelativeResource(helper, link.Id, model.LINK)
		for _, r := range ress {
			catalog := model.Catalog{}
			catalog.Id = r.RId()
			catalog.Name = r.RName()
			link.Catalog = append(link.Catalog, catalog)
		}
	}
	
	return link, result
}

func DeleteLinkById(helper modelhelper.Model, id int) bool {
	sql := fmt.Sprintf(`delete from link where id =%d`, id)	
	num, result := helper.Execute(sql)
	if num > 0 && result {
		link := model.LinkDetail{}
		link.Id = id
		result  = DeleteResource(helper, &link)
	}
	
	return result	
}

func SaveLink(helper modelhelper.Model, link model.LinkDetail) bool {
	sql := fmt.Sprintf(`select id from link where id=%d`, link.Id)
	helper.Query(sql)

	result := false;
	if helper.Next() {
		var id = 0
		helper.GetValue(&id)
		result = true
	}

	if !result {
		// insert
		sql = fmt.Sprintf(`insert into link (name,url,logo,creater) values ('%s','%s','%s', %d)`, link.Name, link.Url, link.Logo, link.Creater.Id)
		_, result = helper.Execute(sql)
		if result {
			sql = fmt.Sprintf(`select id from link where name='%s' and url ='%s' and creater=%d`, link.Name, link.Url, link.Creater.Id)
			
			helper.Query(sql)
			if helper.Next() {
				helper.GetValue(&link.Id)
			}
		}
	} else {
		// modify
		sql = fmt.Sprintf(`update link set name ='%s', url ='%s', logo='%s', creater=%d where id=%d`, link.Name, link.Url, link.Logo, link.Creater.Id, link.Id)
		_, result = helper.Execute(sql)
	}
	
	if result {
		result = SaveResource(helper, &link)
	}
	
	return result	
}
