package bll

import (
    "webcenter/util/modelhelper"
    "webcenter/kernel/dal"
)

type Module struct {
	Id string
	Name string
	Description string
	Enable bool
}

type Block struct {
	Id int
	Name string
}

type Page struct {
	Url string
	Blocks []Block
}

type ModuleDetail struct {
	Module
	Blocks []Block
	Pages []Page
}

func QueryAllModules() []Module {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	moduleList := []Module{}
	
	modules := dal.QueryAllModule(helper)
	for i, _ := range modules {
		m := &modules[i]
		
		module := Module{Id:m.Id, Name:m.Name, Description:m.Description, Enable:m.EnableFlag == 1}
				
		moduleList = append(moduleList, module)
	}

	return moduleList	
}

func QueryModuleDetail(id string) (ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	detail := ModuleDetail{}
	
	module, found := dal.QueryModule(helper, id)
	if found {
		detail.Id = module.Id
		detail.Name = module.Name
		detail.Description = module.Description
		detail.Enable = module.EnableFlag == 1
		
		blocks := dal.QueryBlocks(helper, id)
		for _, b := range blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			detail.Blocks = append(detail.Blocks, block)
		}
		
		pages := dal.QueryPages(helper, id)
		for _, p := range pages {
			page := Page{}
			page.Url = p.Url
			
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)				
			}
			
			detail.Pages = append(detail.Pages, page)
		}
	}
	
	return detail, found
}

func EnableModules(enableList []string) ([]Module, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	helper.BeginTransaction()
	
	ok := true
	modules := dal.QueryAllModule(helper)
	for i, _ := range modules {
		m := &modules[i]
		
		found := false
		for _, id := range enableList {
			if m.Id == id {
				found = true
				break;
			}
		}

		if found {
			if m.EnableFlag == 0 {
				m.EnableFlag = 1
				_, ok = dal.SaveModule(helper, *m)
			}
		} else {
			if m.EnableFlag == 1 {
				m.EnableFlag = 0
				_, ok = dal.SaveModule(helper, *m)
			}
		}
		
		if !ok {
			break
		}
	}

	moduleList := []Module{}
	if ok {
		helper.Commit()
		
		modules := dal.QueryAllModule(helper)
		for i, _ := range modules {
			m := &modules[i]
			
			module := Module{Id:m.Id, Name:m.Name, Description:m.Description, Enable:m.EnableFlag == 1}
					
			moduleList = append(moduleList, module)
		}
	} else {
		helper.Rollback()
	}
	
	return moduleList, ok
}

func AddModuleBlock(name, owner string) (ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := ModuleDetail{}
	_, ok := dal.InsertBlock(helper, name, owner)
	if !ok {
		return detail, ok
	}
	
	module, found := dal.QueryModule(helper, owner)
	if found {
		detail.Id = module.Id
		detail.Name = module.Name
		detail.Description = module.Description
		detail.Enable = module.EnableFlag == 1
		
		blocks := dal.QueryBlocks(helper, owner)
		for _, b := range blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			detail.Blocks = append(detail.Blocks, block)
		}
		
		pages := dal.QueryPages(helper, owner)
		for _, p := range pages {
			page := Page{}
			page.Url = p.Url
			
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)				
			}
			
			detail.Pages = append(detail.Pages, page)
		}
	}
	
	return detail, found	
}

func RemoveModuleBlock(id int, owner string) (ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := ModuleDetail{}
	ok := dal.DeleteBlock(helper, id)
	if !ok {
		return detail, ok
	}
	
	module, found := dal.QueryModule(helper, owner)
	if found {
		detail.Id = module.Id
		detail.Name = module.Name
		detail.Description = module.Description
		detail.Enable = module.EnableFlag == 1
		
		blocks := dal.QueryBlocks(helper, owner)
		for _, b := range blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			detail.Blocks = append(detail.Blocks, block)
		}
		
		pages := dal.QueryPages(helper, owner)
		for _, p := range pages {
			page := Page{}
			page.Url = p.Url
			
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)				
			}
			
			detail.Pages = append(detail.Pages, page)
		}
	}
	
	return detail, found	
}

func SavePageBlock(owner,url string, blocks []int) (ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := ModuleDetail{}
	helper.BeginTransaction()
	_, ok := dal.SavePage(helper, owner, url, blocks)
	if !ok {
		helper.Rollback()
		
		return detail, ok
	}
	
	helper.Commit()
	
	module, found := dal.QueryModule(helper, owner)
	if found {
		detail.Id = module.Id
		detail.Name = module.Name
		detail.Description = module.Description
		detail.Enable = module.EnableFlag == 1
		
		blocks := dal.QueryBlocks(helper, owner)
		for _, b := range blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			detail.Blocks = append(detail.Blocks, block)
		}
		
		pages := dal.QueryPages(helper, owner)
		for _, p := range pages {
			page := Page{}
			page.Url = p.Url
			
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)				
			}
			
			detail.Pages = append(detail.Pages, page)
		}
	}
	
	return detail, found	
}


