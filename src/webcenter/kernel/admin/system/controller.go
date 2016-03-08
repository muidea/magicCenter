package system

import (
    "webcenter/util/modelhelper"
    "webcenter/module"
    "webcenter/kernel"
    "webcenter/kernel/admin/common"
)

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
	disableList []string
	defaultModule []string
}

type ApplyModuleResult struct {
	common.Result
}

type Module struct {
	Id string
	Name string
}

type Block struct {
	Id int
	Name string
}

type Page struct {
	Url string
	Blocks []int
}

type QueryModuleInfoParam struct {
	id string
}

type QueryModuleInfoResult struct {
	common.Result
	Module Module
	Blocks []Block
	Pages []Page
}

type DeleteModuleBlockParam struct {
	id int
	owner string
}

type DeleteModuleBlockResult struct {
	common.Result
	Blocks []Block
}

type SaveModuleBlockParam struct {
	owner string
	block string
}

type SaveModuleBlockResult struct {
	common.Result
	Owner string
	Blocks []Block
}

type SavePageBlockParam struct {
	url string
	blocks []int
}

type SavePageBlockResult struct {
	common.Result
	Page Page
}

type systemController struct {
	
}

func (this *systemController)UpdateSystemInfoAction(param UpdateSystemInfoParam) UpdateSystemInfoResult {
	result := UpdateSystemInfoResult{}
	
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	if param.name != "" && param.name != kernel.Name() {
		kernel.UpdateName(param.name)
		UpdateSystemName(model, param.name)
	}
	
	if param.logo != "" && param.logo != kernel.Logo() {
		kernel.UpdateLogo(param.logo)
		UpdateSystemLogo(model, param.logo)
	}
	
	if param.domain != "" && param.domain != kernel.Domain() {
		kernel.UpdateDomain(param.domain)
		UpdateSystemDomain(model, param.domain)
	}
		
	if param.emailServer != "" && param.emailServer != kernel.MailServer() {
		kernel.UpdateMailServer(param.emailServer)
		UpdateSystemEMailServer(model, param.emailServer)
	}

	if param.emailAccount != "" && param.emailAccount != kernel.MailAccount() {
		kernel.UpdateMailAccount(param.emailAccount)
		UpdateSystemEMailAccount(model, param.emailAccount)
	}

	if param.emailPassword != "" && param.emailPassword != kernel.MailPassword() {
		kernel.UpdateMailPassword(param.emailPassword)
		UpdateSystemEMailPassword(model, param.emailPassword)
	}
	
	result.ErrCode = 0
	result.Reason = "保存站点信息成功"
	
	return result
}

func (this *systemController)ApplyModuleAction(param ApplyModuleParam) ApplyModuleResult {
	result := ApplyModuleResult{}
	
	for _, v := range param.enableList {
		module.EnableModule(v)
	}
	
	for _, v := range param.disableList {
		module.DisableModule(v)
	}
	
	module.UndefaultAllModule()
	for _, v := range param.defaultModule {
		module.DefaultModule(v)
	}
	
	result.ErrCode = 0;
	result.Reason = "操作成功"
	
	return result
}

func (this *systemController)QueryModuleInfoAction(param QueryModuleInfoParam) QueryModuleInfoResult {
	result := QueryModuleInfoResult{}
	
	m,found := module.QueryModule(param.id)
	if found {
		result.Module.Name = m.Name()
		result.Module.Id = m.ID()
		blocks := module.QueryModuleBlocks(param.id)
		for _, b := range blocks {
			item := Block{}
			item.Id = b.ID()
			item.Name = b.Name()
			
			result.Blocks = append(result.Blocks, item)
		}
		
		urls := m.Urls()
		for _, u := range urls {
			p := module.QueryPage(u)
			
			page := Page{}
			page.Url = u
			
			blocks := p.Blocks()
			for _, b := range blocks {
				page.Blocks = append(page.Blocks, b.ID())
			}
			
			result.Pages = append(result.Pages, page)
		}
				
		result.ErrCode = 0
		result.Reason = "查询成功"
	} else {
		result.ErrCode = 1
		result.Reason = "指定Module不存在"
	}
	
	return result
}


func (this *systemController)DeleteModuleBlockAction(param DeleteModuleBlockParam) DeleteModuleBlockResult {
	result := DeleteModuleBlockResult{}
	
	module.DeleteModuleBlock(param.id)
	
	blocks := module.QueryModuleBlocks(param.owner)
	for _, b := range blocks {
		item := Block{}
		item.Id = b.ID()
		item.Name = b.Name()		
		result.Blocks = append(result.Blocks, item)
	}	
	
	result.ErrCode = 0
	result.Reason = "删除成功"
	
	return result
}

func (this *systemController)SaveModuleBlockAction(param SaveModuleBlockParam) SaveModuleBlockResult {
	result := SaveModuleBlockResult{}

	_, ok := module.InsertModuleBlock(param.block,param.owner)
	if ok {
		result.ErrCode = 0
		result.Reason = "保存数据成功"
		result.Owner = param.owner
		blocks := module.QueryModuleBlocks(param.owner)
		for _, b := range blocks {
			block := Block{}
			block.Id = b.ID()
			block.Name = b.Name()
			
			result.Blocks = append(result.Blocks, block)
		}
	} else {
		result.ErrCode = 1
		result.Reason = "保存数据失败"
	}
	
	return result
}

func (this *systemController)SavePageBlockAction(param SavePageBlockParam) SavePageBlockResult {
	result := SavePageBlockResult{}

	blocks := module.SavePageBlocks(param.url,param.blocks)
	result.ErrCode = 0
	result.Reason = "保存数据成功"
	result.Page.Url = param.url
	result.Page.Blocks = blocks
	
	return result
}
