package bll

import (
	"fmt"
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/module/dal"
    "magiccenter/kernel/module/model"
    "magiccenter/configuration"
    contentBll "magiccenter/kernel/content/bll"
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
			
			if view.Style != 0 {
				// 说明是显示内容,所以这里要继续把Block下对应item的内容取出来	
				for ii, _ := range view.Items {
					item := &view.Items[ii]
					article, found := contentBll.QueryArticleById(item.Id)
					if found {
						content := model.Content{}
						content.Article = article
						content.Url = fmt.Sprintf("%s/view/?id=%d",uri, article.Id)
						pageView.Contents = append(pageView.Contents, content)
					}					
				}			
			}
		}
	}
	pageView.Url = page.Url
	pageView.Owner = page.Owner
	
	return pageView, found
}

func QueryContentView(module, url string, id int) (model.PageContentView, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	contentView := model.PageContentView{}
	
	page, found := dal.QueryPage(helper, module, url)
	if !found {
		return contentView, found
	}
	
	m, found := dal.QueryModule(helper, page.Owner)
	if !found {
		return contentView, found
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
			contentView.Blocks = append(contentView.Blocks, view)
			
			if view.Style != 0 {
				// 说明是显示内容,所以这里要继续把Block下对应item的内容取出来	
				for ii, _ := range view.Items {
					item := &view.Items[ii]
					article, found := contentBll.QueryArticleById(item.Id)
					if found {
						content := model.Content{}
						content.Article = article
						content.Url = fmt.Sprintf("%s/view/?id=%d",uri, article.Id)
						contentView.Contents = append(contentView.Contents, content)
					}					
				}			
			}
		}
	}
	
	article, found := contentBll.QueryArticleById(id)
	if found {
		contentView.Content = article 
	}
	
	contentView.Url = page.Url
	contentView.Owner = page.Owner
	
	return contentView, found
}





