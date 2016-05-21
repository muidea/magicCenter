package bll

import (
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/module/dal"
    "magiccenter/kernel/module/model"
    contentdal "magiccenter/kernel/content/dal"
    "magiccenter/module"
)

func QueryModuleContent(id string) (model.ModuleContent, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := model.ModuleContent{}
	instance, found := module.FindModule(id)
	if !found {
		return detail, found
	}

	m, found := dal.QueryModule(helper, id)
	if !found {
		m.Id = instance.ID()
		m.Name = instance.Name()
		m.Description = instance.Description()
		m.EnableFlag = 0
		m, found = dal.SaveModule(helper, m)
	}
	
	if found {
		detail.Id = m.Id
		detail.Name = m.Name
		detail.Description = m.Description
		detail.EnableFlag = m.EnableFlag
		detail.Blocks = dal.QueryBlockDetails(helper, id)
	}
	
	return detail, found
}

func SaveBlockItem(block int, articleList, catalogList, linkList []int) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	helper.BeginTransaction()
	
	dal.ClearItems(helper, block)
	
	for _, ar := range articleList {
		article, found := contentdal.QueryArticleById(helper, ar)
		if found {
			dal.AddItem(helper, article.RId(), article.RType(), block) 
		}
	}
	
	for _, ca := range catalogList {
		catalog, found := contentdal.QueryCatalogById(helper, ca)
		if found {
			dal.AddItem(helper, catalog.RId(), catalog.RType(), block) 
		}
	}
	
	for _, lnk := range linkList {
		link, found := contentdal.QueryLinkById(helper, lnk)
		if found {
			dal.AddItem(helper, link.RId(), link.RType(), block) 
		}
	}
	
	helper.Commit()
}


