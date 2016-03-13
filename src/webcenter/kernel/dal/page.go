package dal

import (
	"fmt"
	"webcenter/util/modelhelper"
)

type Page struct {
	Owner string
	Url string
	Blocks []Block
}

func QueryPage(helper modelhelper.Model, owner, url string) (Page, bool) {
	page := Page{}
	ret := false
	
	sql := fmt.Sprintf("select id,name,owner from block where id in (select block from page where owner='%s' and url='%s')", owner, url)
	helper.Query(sql)
	
	for helper.Next() {
		b := Block{}
		helper.GetValue(&b.Id, &b.Name, &b.Owner)
		
		page.Blocks = append(page.Blocks, b)
		
		ret = true
	}
	
	for i, _ := range page.Blocks {
		b := &page.Blocks[i]
		b.Items = QueryItems(helper, b.Id)
	}
	page.Owner = owner
	page.Url = url
	
	return page, ret
}

func QueryPages(helper modelhelper.Model, owner string) []Page {
	
	sql := fmt.Sprintf("select distinct url from page where owner='%s'", owner)
	helper.Query(sql)
	
	urlList := []string{}
	for helper.Next() {
		url := ""
		helper.GetValue(&url)
		
		urlList = append(urlList, url)
	}
	
	pageList := []Page{}
	for _, url := range urlList {
		page, _ := QueryPage(helper, owner, url)
		pageList = append(pageList, page)
	}
	
	return pageList
}

func SavePage(helper modelhelper.Model, page Page) (Page, bool) {
	ret := false
	sql := fmt.Sprintf("delete from page where owner='%s' and url='%s'", page.Owner, page.Url)
	_, ret = helper.Execute(sql)
	if ret {
		for _, block := range page.Blocks {
			sql = fmt.Sprintf("insert into page(owner,url,block) values('%s','%s',%d)", page.Owner, page.Url, block.Id)
			helper.Execute(sql)
		}
	}
	
	return page,ret
}
