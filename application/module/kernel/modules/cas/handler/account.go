package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
)

type accountManager struct {
	moduleHub common.ModuleHub
}

func createAccountManager(modHub common.ModuleHub) accountManager {
	return accountManager{moduleHub: modHub}
}

// model.UserDetail 登陆用户
// bool 是否登陆成功
func (i *accountManager) findUser(account, password string) (model.UserDetail, bool) {
	user := model.UserDetail{}
	found := false
	modHub, ok := i.moduleHub.FindModule(common.AccountModuleID)
	if !ok {
		return user, false
	}

	endPoint := modHub.EntryPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		accountHandler := endPoint.(common.AccountHandler)
		user, found = accountHandler.FindUserByAccount(account, password)
	}
	return user, found
}
