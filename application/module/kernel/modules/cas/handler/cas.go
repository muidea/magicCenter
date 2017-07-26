package handler

import (
	"net/http"

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

/*
1、先判断authToken是否一致，如果不一致则，认为无权限
*/
func (i *impl) VerifyAccount(res http.ResponseWriter, req *http.Request) bool {
	session := i.sessionRegistry.GetSession(res, req)

	_, ok := session.GetAccount()
	if !ok {
		return false
	}

	urlToken := ""
	authToken, found := req.URL.Query()[common.AuthTokenID]
	if found {
		// 如果url里没有token，则认为token为空，判断时使用空字符串进行比较
		urlToken = authToken[0]
	}

	sessionToken := ""
	obj, ok := session.GetOption(common.AuthTokenID)
	if ok {
		// 找到sessionToken了，则说明该用户已经登录了，这里就必须保证两端的token一致否则也要认为鉴权非法
		// 用户登录过Token必然不为空
		sessionToken = obj.(string)
		return urlToken == sessionToken
	}

	return true
}
