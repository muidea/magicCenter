package system

import (
    "webcenter/application"
    "webcenter/modelhelper"
    "webcenter/admin/common"
)


type UpdateParam struct {
	name string
	logo string
	domain string
	emailServer string
	emailAccount string
	emailPassword string
	accesscode string
}

type UpdateResult struct {
	common.Result
}


type systemController struct {
	
}

func (this *systemController)UpdateAction(param *UpdateParam) UpdateResult {
	result := UpdateResult{}
	
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	if param.name != "" && param.name != application.Name() {
		application.UpdateName(param.name)
		UpdateSystemName(model, param.name)
	}
	
	if param.logo != "" && param.logo != application.Logo() {
		application.UpdateLogo(param.logo)
		UpdateSystemLogo(model, param.logo)
	}
	
	if param.domain != "" && param.domain != application.Domain() {
		application.UpdateDomain(param.domain)
		UpdateSystemDomain(model, param.domain)
	}
		
	if param.emailServer != "" && param.emailServer != application.MailServer() {
		application.UpdateMailServer(param.emailServer)
		UpdateSystemEMailServer(model, param.emailServer)
	}

	if param.emailAccount != "" && param.emailAccount != application.MailAccount() {
		application.UpdateMailAccount(param.emailAccount)
		UpdateSystemEMailAccount(model, param.emailAccount)
	}

	if param.emailPassword != "" && param.emailPassword != application.MailPassword() {
		application.UpdateMailPassword(param.emailPassword)
		UpdateSystemEMailPassword(model, param.emailPassword)
	}
	
	result.ErrCode = 0
	result.Reason = "保存站点信息成功"
	
	return result
}




