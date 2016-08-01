package dal

import (
	"fmt"
	"magiccenter/kernel/dashboard/module/model"
	"magiccenter/util/modelhelper"
)

// QueryPage 查询page信息
func QueryPage(helper modelhelper.Model, owner, url string) (model.Page, bool) {
	page := model.Page{}
	ret := true

	sql := fmt.Sprintf("select id,name,owner from block where id in (select block from page where owner='%s' and url='%s')", owner, url)
	helper.Query(sql)

	for helper.Next() {
		b := model.Block{}
		helper.GetValue(&b.ID, &b.Name, &b.Owner)

		page.Blocks = append(page.Blocks, b)
	}

	page.Owner = owner
	page.URL = url

	return page, ret
}

// QueryPages 查询Module的Pages
func QueryPages(helper modelhelper.Model, owner string) []model.Page {

	sql := fmt.Sprintf("select distinct url from page where owner='%s' order by url", owner)
	helper.Query(sql)

	urlList := []string{}
	for helper.Next() {
		url := ""
		helper.GetValue(&url)

		urlList = append(urlList, url)
	}

	pageList := []model.Page{}
	for _, url := range urlList {
		page, _ := QueryPage(helper, owner, url)
		pageList = append(pageList, page)
	}

	return pageList
}

// SavePage 保存页面信息
func SavePage(helper modelhelper.Model, owner, url string, blocks []int) (model.Page, bool) {
	ret := false
	sql := fmt.Sprintf("delete from page where owner='%s' and url='%s'", owner, url)
	_, ret = helper.Execute(sql)
	if ret {
		ret = true
		for _, b := range blocks {
			sql = fmt.Sprintf("insert into page(owner,url,block) values('%s','%s',%d)", owner, url, b)
			num, ok := helper.Execute(sql)
			if num != 1 || !ok {
				ret = false
				break
			}
		}
	}

	if !ret {
		page := model.Page{}
		return page, ret
	}

	return QueryPage(helper, owner, url)
}
