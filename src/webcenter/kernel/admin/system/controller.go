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

type QueryModuleBlockParam struct {
	id string
}

type QueryModuleBlockResult struct {
	common.Result
	Module Module
	Blocks []Block	
}

type SaveModuleBlockParam struct {
	module string
	block string
}

type SaveModuleBlockResult struct {
	common.Result
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

func (this *systemController)QueryModuleBlockAction(param QueryModuleBlockParam) QueryModuleBlockResult {
	result := QueryModuleBlockResult{}
	
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
		
		result.ErrCode = 0
		result.Reason = "查询成功"
	} else {
		result.ErrCode = 1
		result.Reason = "指定Module不存在"
	}
	
	return result
}

func (this *systemController)SaveModuleBlockAction(param SaveModuleBlockParam) SaveModuleBlockResult {
	result := SaveModuleBlockResult{}

	_, ok := module.InsertModuleBlock(param.block,param.module)
	if ok {
		result.ErrCode = 0
		result.Reason = "保存数据成功"
	} else {
		result.ErrCode = 1
		result.Reason = "保存数据失败"
	}
	
	return result
}

