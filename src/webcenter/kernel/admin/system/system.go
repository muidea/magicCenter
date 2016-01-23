package system

import (
    "webcenter/util/modelhelper"
    "webcenter/kernel"
)

func init() {
	model, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct model failed")
	}
	defer model.Release()
	
	name, found := GetSystemName(model)
	if found {
		kernel.UpdateName(name)
	}
	
	logo, found := GetSystemLogo(model)
	if found {
		kernel.UpdateLogo(logo)
	}
	
	domain, found := GetSystemDomain(model)
	if found {
		kernel.UpdateDomain(domain)
	}
	
	emailServer, found := GetSystemEMailServer(model)
	if found {
		kernel.UpdateMailServer(emailServer)
	}
	
	emailAccount, found := GetSystemEMailAccount(model)
	if found {
		kernel.UpdateMailAccount(emailAccount)
	}
	
	emailPassword, found := GetSystemEMailPassword(model)
	if found {
		kernel.UpdateMailPassword(emailPassword)
	}
	
}