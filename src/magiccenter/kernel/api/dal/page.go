package dal

import (
	"fmt"
	"magiccenter/common"
	"magiccenter/common/model"
)

// QueryPage 查询page信息
func QueryPage(helper common.DBHelper, owner, url string) (model.Page, bool) {
	page := model.Page{}
	ret := true

	sql := fmt.Sprintf("select block from page where owner='%s' and url='%s'", owner, url)
	helper.Query(sql)

	for helper.Next() {
		b := -1
		helper.GetValue(&b)
		page.Blocks = append(page.Blocks, b)
	}

	page.Owner = owner
	page.URL = url

	return page, ret
}

// QueryPages 查询Module的Pages
func QueryPages(helper common.DBHelper, owner string) []model.Page {

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
func SavePage(helper common.DBHelper, page model.Page) bool {
	ret := false
	sql := fmt.Sprintf("delete from page where owner='%s' and url='%s'", page.Owner, page.URL)
	_, ret = helper.Execute(sql)
	if ret {
		ret = true
		for _, b := range page.Blocks {
			sql = fmt.Sprintf("insert into page(owner,url,block) values('%s','%s',%d)", page.Owner, page.URL, b)
			num, ok := helper.Execute(sql)
			if num != 1 || !ok {
				ret = false
				break
			}
		}
	}

	return ret
}

// DeletePage 删除指定页面
func DeletePage(helper common.DBHelper, owner, url string) bool {
	ret := false
	sql := fmt.Sprintf("delete from page where owner='%s' and url='%s'", owner, url)
	num, ret := helper.Execute(sql)
	return num == 1 && ret
}
