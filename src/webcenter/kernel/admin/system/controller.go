package system

import (
    "webcenter/kernel/bll"
    "webcenter/kernel/admin/common"
    "webcenter/configuration"
)

type Module struct {
	Id string
	Name string
	Description string
	Enable bool
	Default bool
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

type SystemInfoView struct {
	configuration.SystemInfo
}

type ModuleInfoView struct {
	Modules []Module
}

type UpdateSystemInfoParam struct {
	name string
	logo string
	domain string
	emailServer string
	emailAccount string
	emailPassword string
	accesscode string
}

type UpdateSystemInfoResult struct {
	common.Result
}

type ApplyModuleParam struct {
	enableList []string
	defaultModule string
}

type ApplyModuleResult struct {
	common.Result
	Modules []Module
}

type QueryModuleDetailParam struct {
	id string
}

type QueryModuleDetailResult struct {
	common.Result
	Module ModuleDetail
}

type DeleteModuleBlockParam struct {
	id int
	owner string
}

type DeleteModuleBlockResult struct {
	common.Result
	Module ModuleDetail
}

type SaveModuleBlockParam struct {
	owner string
	block string
}

type SaveModuleBlockResult struct {
	common.Result
	Module ModuleDetail
}

type SavePageBlockParam struct {
	owner string
	url string
	blocks []int
}

type SavePageBlockResult struct {
	common.Result
	Module ModuleDetail
}

const passwordMark = "********"

func GetSystemInfoViewAction() SystemInfoView {
	view := SystemInfoView{}
	info := configuration.GetSystemInfo()
	view.Name = info.Name
	view.Logo = info.Logo
	view.Domain = info.Domain
	view.MailServer = info.MailServer
	view.MailAccount = info.MailAccount
	view.MailPassword = passwordMark
	
	return view
}

func UpdateSystemInfoAction(param UpdateSystemInfoParam) UpdateSystemInfoResult {
	result := UpdateSystemInfoResult{}
	
	info := configuration.GetSystemInfo()
	info.Name = param.name
	info.Logo = param.logo
	info.Domain = param.domain
	info.MailServer = param.emailServer
	info.MailAccount = param.emailAccount
	info.MailPassword = param.emailPassword

	if configuration.UpdateSystemInfo(info) {
		
		result.ErrCode = 0
		result.Reason = "保存站点信息成功"		
	} else {
		result.ErrCode = 1
		result.Reason = "保存站点信息失败"
	}
		
	return result
}

func GetModuleInfoViewAction() ModuleInfoView {
	view := ModuleInfoView{}
	
	defaultModule, found := configuration.GetOption(configuration.SYS_DEFULTMODULE)
	modules := bll.QueryAllModules()
	for _, m := range modules {
		module := Module{}
		module.Id = m.Id
		module.Name = m.Name
		module.Description = m.Description
		module.Enable = m.Enable
		if found && defaultModule == m.Id {
			module.Default = true
		} else {
			module.Default = false
		}
		
		view.Modules = append(view.Modules, module)
	}
	
	return view
}

func ApplyModuleAction(param ApplyModuleParam) ApplyModuleResult {
	result := ApplyModuleResult{}
	
	configuration.SetOption(configuration.SYS_DEFULTMODULE, param.defaultModule)
	
	modules, ok := bll.EnableModules(param.enableList)
	if ok {
		result.ErrCode = 0;
		result.Reason = "操作成功"
		
		for i, _ := range modules {
			m := &modules[i]
			
			module := Module{}
			module.Id = m.Id
			module.Name = m.Name
			module.Description = m.Description
			module.Enable = m.Enable
			
			result.Modules = append(result.Modules, module)
		}
	} else {
		result.ErrCode = 1;
		result.Reason = "操作失败"		
	}
	
	return result
}

func QueryModuleDetailAction(param QueryModuleDetailParam) QueryModuleDetailResult {
	result := QueryModuleDetailResult{}
		
	m,found := bll.QueryModuleDetail(param.id)
	if found {
		result.Module.Id = m.Id
		result.Module.Name = m.Name
		result.Module.Description = m.Description
		result.Module.Enable = m.Enable
		for _, b := range m.Blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			result.Module.Blocks = append(result.Module.Blocks, block)
		}
		
		for _, p := range m.Pages {
			page := Page{}
			page.Url = p.Url
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)
			}
			
			result.Module.Pages = append(result.Module.Pages, page)
		}
				
		result.ErrCode = 0
		result.Reason = "查询成功"
	} else {
		result.ErrCode = 1
		result.Reason = "指定Module不存在"
	}
	
	return result
}


func DeleteModuleBlockAction(param DeleteModuleBlockParam) DeleteModuleBlockResult {
	result := DeleteModuleBlockResult{}
	
	m,ret := bll.RemoveModuleBlock(param.id, param.owner)
	if ret {
		result.Module.Id = m.Id
		result.Module.Name = m.Name
		result.Module.Description = m.Description
		result.Module.Enable = m.Enable
		for _, b := range m.Blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			result.Module.Blocks = append(result.Module.Blocks, block)
		}
		
		for _, p := range m.Pages {
			page := Page{}
			page.Url = p.Url
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)
			}
			
			result.Module.Pages = append(result.Module.Pages, page)
		}
				
		result.ErrCode = 0
		result.Reason = "查询成功"
	} else {
		result.ErrCode = 1
		result.Reason = "指定Module不存在"
	}
	
	return result
}

func SaveModuleBlockAction(param SaveModuleBlockParam) SaveModuleBlockResult {
	result := SaveModuleBlockResult{}

	m,ret := bll.AddModuleBlock(param.block, param.owner)
	if ret {
		result.Module.Id = m.Id
		result.Module.Name = m.Name
		result.Module.Description = m.Description
		result.Module.Enable = m.Enable
		for _, b := range m.Blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			result.Module.Blocks = append(result.Module.Blocks, block)
		}
		
		for _, p := range m.Pages {
			page := Page{}
			page.Url = p.Url
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)
			}
			
			result.Module.Pages = append(result.Module.Pages, page)
		}
				
		result.ErrCode = 0
		result.Reason = "查询成功"
	} else {
		result.ErrCode = 1
		result.Reason = "指定Module不存在"
	}
	
	return result
}

func SavePageBlockAction(param SavePageBlockParam) SavePageBlockResult {
	result := SavePageBlockResult{}

	m,ret := bll.SavePageBlock(param.owner, param.url, param.blocks)
	if ret {
		result.Module.Id = m.Id
		result.Module.Name = m.Name
		result.Module.Description = m.Description
		result.Module.Enable = m.Enable
		for _, b := range m.Blocks {
			block := Block{}
			block.Id = b.Id
			block.Name = b.Name
			result.Module.Blocks = append(result.Module.Blocks, block)
		}
		
		for _, p := range m.Pages {
			page := Page{}
			page.Url = p.Url
			for _, b := range p.Blocks {
				block := Block{}
				block.Id = b.Id
				block.Name = b.Name
				page.Blocks = append(page.Blocks, block)
			}
			
			result.Module.Pages = append(result.Module.Pages, page)
		}
				
		result.ErrCode = 0
		result.Reason = "查询成功"
	} else {
		result.ErrCode = 1
		result.Reason = "指定Module不存在"
	}
	
	return result	
}
