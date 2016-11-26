package bll

import (
	"magiccenter/common"
	"magiccenter/common/model"
	"magiccenter/kernel/api/dal"
	"magiccenter/system"
)

func queryAllModuleInternal(helper common.DBHelper) []model.Module {

	// 由于部分Module可能还未启用，所以这里需要取DB和系统加载信息的全集
	modules := dal.QueryAllModule(helper)

	modulehub := system.GetModuleHub()
	sysModule := modulehub.QueryAllModule()

	for _, sysMod := range sysModule {
		if sysMod.Type() == common.KERNEL {
			// 不处理Kernal模块
			continue
		}

		mod, found := dal.QueryModule(helper, sysMod.ID())
		if !found {
			mod.ID = sysMod.ID()
			mod.Name = sysMod.Name()
			mod.Description = sysMod.Description()
			mod.URL = sysMod.URL()
			mod.Type = sysMod.Type()
			mod.Status = sysMod.Status()

			modules = append(modules, mod)
		}
	}

	return modules
}

// QueryAllModules 查询所有Module
func QueryAllModules() []model.Module {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	return queryAllModuleInternal(helper)
}

// QueryModule 查询指定Module
func QueryModule(id string) (model.Module, bool) {
	mod := model.Module{}
	modulehub := system.GetModuleHub()
	sysMod, found := modulehub.FindModule(id)
	if found {
		mod.ID = sysMod.ID()
		mod.Name = sysMod.Name()
		mod.Description = sysMod.Description()
		mod.URL = sysMod.URL()
		mod.Type = sysMod.Type()
		mod.Status = sysMod.Status()
	}

	return mod, found
}

// EnableModules 启动模块
func EnableModules(enableList []string) ([]model.Module, bool) {
	helper, err := system.GetDBHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()

	helper.BeginTransaction()

	ok := true
	modules := queryAllModuleInternal(helper)
	for _, m := range modules {
		found := false
		for _, id := range enableList {
			if m.ID == id {
				found = true
				break
			}
		}

		if found {
			if m.Status == 0 {
				m.Status = 1
				_, ok = dal.SaveModule(helper, m)
			}
		} else {
			if m.Status == 1 {
				m.Status = 0
				_, ok = dal.SaveModule(helper, m)
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
		m.Status = 0
		m, found = dal.SaveModule(helper, m)
	}

	if found {
		detail.Id = m.Id
		detail.Name = m.Name
		detail.Description = m.Description
		detail.Status = m.Status
		detail.Blocks = dal.QueryBlocks(helper, id)

		rts := instance.Routes()
		for _, r := range rts {
			p, _ := dal.QueryPage(helper, id, r.Pattern())
			detail.Pages = append(detail.Pages, p)
		}
	}

	return detail, found
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
		m.Status = 0
		m, found = dal.SaveModule(helper, m)
	}

	if found {
		detail.Id = m.Id
		detail.Name = m.Name
		detail.Description = m.Description
		detail.Status = m.Status
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
		m.Status = 0
		m, found = dal.SaveModule(helper, m)
	}

	if found {
		detail.Id = m.Id
		detail.Name = m.Name
		detail.Description = m.Description
		detail.Status = m.Status
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
		m.Status = 0
		m, found = dal.SaveModule(helper, m)
	}

	if found {
		detail.Id = m.Id
		detail.Name = m.Name
		detail.Description = m.Description
		detail.Status = m.Status
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
