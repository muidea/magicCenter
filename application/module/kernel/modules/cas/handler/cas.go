package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateCASHandler 新建CASHandler
func CreateCASHandler(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.CASHandler {
	i := impl{
		sessionRegistry: sessionRegistry,
		accountManager:  createAccountManager(modHub),
		cacheData:       cache.NewCache()}

	return &i
}

type impl struct {
	sessionRegistry common.SessionRegistry
	accountManager  accountManager
	cacheData       cache.Cache
}

func (i *impl) LoginAccount(account, password string) (model.UserDetail, string, bool) {
	user, ok := i.accountManager.findUser(account, password)
	if !ok {
		return user, "", ok
	}

	token := i.cacheData.PutIn(user, cache.MaxAgeValue)
	return user, token, ok
}

func (i *impl) LoginToken(token string) (string, bool) {
	return token, true
}

func (i *impl) Logout(authToken string) bool {
	_, ok := i.cacheData.FetchOut(authToken)
	if !ok {
		return false
	}

	i.cacheData.Remove(authToken)
	return true
}

func (i *impl) VerifyToken(authToken string) bool {
	return true
}
