package bll

import (
    "magiccenter/util/modelhelper"
    "magiccenter/kernel/module/dal"
    "magiccenter/kernel/module/model"
    "magiccenter/module"    
)

func QueryAllModules() []model.Module {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	return dal.QueryAllModule(helper)
}

func QueryModuleDetail(id string) (model.ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := model.ModuleDetail{}
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
		detail.Blocks = dal.QueryBlocks(helper, id)
		
		rts := instance.Routes()
		for _, r := range rts {
			p, _ := dal.QueryPage(helper, id, r.Pattern())
			detail.Pages = append(detail.Pages, p)
		}
	}
	
	return detail, found
}

func EnableModules(enableList []string) ([]model.Module, bool) {
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

	moduleList := []model.Module{}
	if ok {
		helper.Commit()
		
		moduleList = dal.QueryAllModule(helper)
	} else {
		helper.Rollback()
	}
	
	return moduleList, ok
}

func AddModuleBlock(name, owner string) (model.ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := model.ModuleDetail{}
	_, ok := dal.InsertBlock(helper, name, owner)
	if !ok {
		return detail, ok
	}
	
	instance, found := module.FindModule(owner)
	if !found {
		return detail, found
	}

	m, found := dal.QueryModule(helper, owner)
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
		detail.Blocks = dal.QueryBlocks(helper, owner)
		
		rts := instance.Routes()
		for _, r := range rts {
			p, _ := dal.QueryPage(helper, owner, r.Pattern())
			detail.Pages = append(detail.Pages, p)
		}
	}	
	
	return detail, found	
}

func RemoveModuleBlock(id int, owner string) (model.ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := model.ModuleDetail{}
	ok := dal.DeleteBlock(helper, id)
	if !ok {
		return detail, ok
	}
	
	instance, found := module.FindModule(owner)
	if !found {
		return detail, found
	}

	m, found := dal.QueryModule(helper, owner)
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
		detail.Blocks = dal.QueryBlocks(helper, owner)
		
		rts := instance.Routes()
		for _, r := range rts {
			p, _ := dal.QueryPage(helper, owner, r.Pattern())
			detail.Pages = append(detail.Pages, p)
		}
	}	
	
	return detail, found	
}

func SavePageBlock(owner,url string, blocks []int) (model.ModuleDetail, bool) {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	detail := model.ModuleDetail{}
	helper.BeginTransaction()
	_, ok := dal.SavePage(helper, owner, url, blocks)
	if !ok {
		helper.Rollback()
		
		return detail, ok
	}
	
	helper.Commit()
	
	instance, found := module.FindModule(owner)
	if !found {
		return detail, found
	}

	m, found := dal.QueryModule(helper, owner)
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
		detail.Blocks = dal.QueryBlocks(helper, owner)
		
		rts := instance.Routes()
		for _, r := range rts {
			p, _ := dal.QueryPage(helper, owner, r.Pattern())
			detail.Pages = append(detail.Pages, p)
		}
	}
	
	return detail, found	
}


