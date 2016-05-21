package bll

import (
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/module/dal"
    "magiccenter/kernel/module/model"
    "magiccenter/configuration"
)



//
// 获取module指定url的内容
// 
func QueryPageView(module, url string) (model.PageView, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	pageView := model.PageView{}
	
	page, found := dal.QueryPage(helper, module, url)
	if !found {
		return pageView, found
	}
	
	m, found := dal.QueryModule(helper, page.Owner)
	if !found {
		return pageView, found
	}

	uri := ""
	defaultModule, _ := configuration.GetOption(configuration.SYS_DEFULTMODULE)
	// 如果不是默认模块，则uri为module的Uri
	if defaultModule != page.Owner {
		uri = m.Uri
	}
	
	for index, _ := range page.Blocks {
		block := &page.Blocks[index]
		
		view,found := dal.QueryBlockView(helper, uri, block.Id)
		if found {
			pageView.Blocks = append(pageView.Blocks, view)
		}
	}
	pageView.Url = page.Url
	pageView.Owner = page.Owner
	
	return pageView, found
}



