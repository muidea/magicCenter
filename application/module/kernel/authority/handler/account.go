package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
)

type accountActionHandler struct {
	moduleHub common.ModuleHub
	authCache cache.Cache
}

// model.UserDetail 登陆用户
// string 本次登陆的token
// bool 是否登陆成功
func (i *accountActionHandler) LoginAccount(account, password string) (model.UserDetail, string, bool) {
	user := model.UserDetail{}
	found := false
	modHub, ok := i.moduleHub.FindModule(common.AccountModuleID)
	if !ok {
		return user, "", false
	}

	endPoint := modHub.EndPoint()
	switch endPoint.(type) {
	case common.AccountHandler:
		accountHandler := endPoint.(common.AccountHandler)
		user, found = accountHandler.FindUserByAccount(account, password)
	}

	authToken := ""
	if found {
		authToken = i.authCache.PutIn(user, cache.MaxAgeValue)
	}

	return user, authToken, found
}

// authID 登陆token
func (i *accountActionHandler) LogoutAccount(authID string) bool {

	_, found := i.authCache.FetchOut(authID)
	if found {
		i.authCache.Remove(authID)
	}

	return found
}

func (i *accountActionHandler) IsLogin(authID string) bool {
	_, found := i.authCache.FetchOut(authID)
	return found
}
