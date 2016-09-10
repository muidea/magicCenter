package bll

import (
	"magiccenter/kernel/dashboard/module/dal"
	"magiccenter/kernel/dashboard/module/model"
	"magiccenter/module"
	"magiccenter/util/dbhelper"
)

func queryAllModuleInternal(helper dbhelper.DBHelper) []model.Module {

	// 由于部分Module可能还未启用，所以这里需要取DB和系统加载信息的全集
	modules := dal.QueryAllModule(helper)

	sysModule := module.QueryAllModule()

	for _, sysMod := range sysModule {
		mod, found := dal.QueryModule(helper, sysMod.ID())
		if !found {
			mod = model.Module{}
			mod.ID = sysMod.ID()
			mod.Name = sysMod.Name()
			mod.Description = sysMod.Description()
			mod.URL = sysMod.URL()
			mod.EnableFlag = 0

			modules = append(modules, mod)
		}
	}

	return modules
}

// QueryAllModules 查询所有Module
func QueryAllModules() []model.Module {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return queryAllModuleInternal(helper)
}

/*
func QueryModuleDetail(id string) (model.ModuleLayout, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	detail := model.ModuleLayout{}
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
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	helper.BeginTransaction()

	ok := true
	modules := queryAllModuleInternal(helper)
	for i, _ := range modules {
		m := &modules[i]

		found := false
		for _, id := range enableList {
			if m.Id == id {
				found = true
				break
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

func AddModuleBlock(name, tag string, style int, owner string) (model.ModuleLayout, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	detail := model.ModuleLayout{}
	_, ok := dal.InsertBlock(helper, name, tag, style, owner)
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

func RemoveModuleBlock(id int, owner string) (model.ModuleLayout, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	detail := model.ModuleLayout{}
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

func SavePageBlock(owner, url string, blocks []int) (model.ModuleLayout, bool) {
	helper, err := dbhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	detail := model.ModuleLayout{}
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
*/
