package bll

import (
	"webcenter/util/modelhelper"
	"webcenter/kernel"
	"webcenter/kernel/dal"
)

func UpdateSystemInfo(info kernel.SystemInfo) bool {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	result := true
	helper.BeginTransaction()
	for true {
		if !updateSystemName(helper, info.Name) {
			result = false
			break
		}
		
		if !updateSystemLogo(helper, info.Logo) {
			result = false
			break
		}
		
		if !updateSystemDomain(helper, info.Domain) {
			result = false
			break
		}
		
		if !updateSystemEMailServer(helper, info.MailServer) {
			result = false
			break
		}
		
		if !updateSystemEMailAccount(helper, info.MailAccount) {
			result = false
			break
		}
		
		if !updateSystemEMailPassword(helper, info.MailPassword) {
			result = false
			break
		}
		
		break;
	}
	
	if result {
		helper.Commit()
	} else {
		helper.Rollback()
	}
	
	return result
}

func GetSystemInfo() kernel.SystemInfo {
	helper, err := modelhelper.NewHelper()
	if err != nil {
		panic("construct helper failed")
	}
	defer helper.Release()
	
	info := kernel.SystemInfo{}
	info.Name, _ = dal.GetOption(helper, "@systemName")
	info.Logo, _ = dal.GetOption(helper, "@systemLogo")
	info.Domain, _ = dal.GetOption(helper, "@systemDomain")
	info.MailServer, _ = dal.GetOption(helper, "@systemEMailServer")
	info.MailAccount, _ = dal.GetOption(helper, "@systemEMailAccount")
	info.MailPassword, _ = dal.GetOption(helper, "@systemEMailPassword")
	return info
}

func updateSystemName(helper modelhelper.Model, systemName string) bool {
	return dal.SetOption(helper, "@systemName", systemName)
}

func updateSystemLogo(helper modelhelper.Model, systemLogo string) bool {
	return dal.SetOption(helper, "@systemLogo", systemLogo)
}

func updateSystemDomain(helper modelhelper.Model, systemDomain string) bool {
	return dal.SetOption(helper, "@systemDomain", systemDomain)
}

func updateSystemEMailServer(helper modelhelper.Model, systemEMailServer string) bool {
	return dal.SetOption(helper, "@systemEMailServer", systemEMailServer)
}

func updateSystemEMailAccount(helper modelhelper.Model, systemEMail string) bool {
	return dal.SetOption(helper, "@systemEMailAccount", systemEMail)
}

func updateSystemEMailPassword(helper modelhelper.Model, systemEMail string) bool {
	return dal.SetOption(helper, "@systemEMailPassword", systemEMail)
}
