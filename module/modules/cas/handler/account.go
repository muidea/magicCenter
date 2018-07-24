package handler

import (
	"time"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCommon/model"
)

type accountManager struct {
	accountHandler common.AccountHandler
}

func createAccountManager(moduleHub common.ModuleHub) (accountManager, bool) {
	s := accountManager{accountHandler: nil}
	mod, ok := moduleHub.FindModule(common.AccountModuleID)
	if !ok {
		return s, false
	}

	entryPoint := mod.EntryPoint()
	switch entryPoint.(type) {
	case common.AccountHandler:
		s.accountHandler = entryPoint.(common.AccountHandler)
	}

	return s, s.accountHandler != nil
}

// model.UserDetail 登陆用户
// bool 是否登陆成功
func (s *accountManager) userLogin(account, password, remoteAddr string) (model.OnlineEntryView, bool) {
	info := model.OnlineEntryView{}

	user, ok := s.accountHandler.FindUserByAccount(account, password)
	if !ok {
		return info, ok
	}

	info.ID = user.ID
	info.Name = user.Name
	info.Address = remoteAddr
	info.LoginTime = time.Now().Unix()
	info.UpdateTime = info.LoginTime
	info.IdentifyID = account

	return info, true
}
