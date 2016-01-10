package system

import (
    "webcenter/application"
    "webcenter/modelhelper"
)

func init() {
	model, err := modelhelper.NewModel()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	name, found := GetSystemName(model)
	if found {
		application.UpdateName(name)
	}
	
	logo, found := GetSystemLogo(model)
	if found {
		application.UpdateLogo(logo)
	}
	
	domain, found := GetSystemDomain(model)
	if found {
		application.UpdateDomain(domain)
	}
	
	emailServer, found := GetSystemEMailServer(model)
	if found {
		application.UpdateMailServer(emailServer)
	}
	
	emailAccount, found := GetSystemEMailAccount(model)
	if found {
		application.UpdateMailAccount(emailAccount)
	}
	
	emailPassword, found := GetSystemEMailPassword(model)
	if found {
		application.UpdateMailPassword(emailPassword)
	}
	
}