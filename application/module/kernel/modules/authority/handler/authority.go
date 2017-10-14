package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateAuthorityHandler 新建CASHandler
func CreateAuthorityHandler(moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	i := impl{
		sessionRegistry: sessionRegistry,
		cacheData:       cache.NewCache()}

	casModule, _ := moduleHub.FindModule(common.CASModuleID)
	entryPoint := casModule.EntryPoint()
	switch entryPoint.(type) {
	case common.CASHandler:
		i.casHandler = entryPoint.(common.CASHandler)
	}

	return &i
}

type impl struct {
	sessionRegistry common.SessionRegistry
	casHandler      common.CASHandler
	cacheData       cache.Cache
}

/*
1、先判断authToken是否一致，如果不一致则，认为无权限
*/
func (i *impl) VerifyAuthority(res http.ResponseWriter, req *http.Request) bool {
	session := i.sessionRegistry.GetSession(res, req)

	_, ok := session.GetAccount()
	if !ok {
		return false
	}

	urlToken := req.URL.Query().Get(common.AuthTokenID)
	sessionToken := ""
	obj, ok := session.GetOption(common.AuthTokenID)
	if ok {
		// 找到sessionToken了，则说明该用户已经登录了，这里就必须保证两端的token一致否则也要认为鉴权非法
		// 用户登录过Token必然不为空
		sessionToken = obj.(string)

		if i.casHandler != nil {
			i.casHandler.RefreshToken(sessionToken, req.RemoteAddr)
		}

		return urlToken == sessionToken
	}

	return true
}

func (i *impl) QueryUserAuthGroup(user model.User) model.AuthGroupInfo {
	return model.AuthGroupInfo{}
}

func (i *impl) UpdateUserAuthGroup(user model.User, authGroupInfo model.AuthGroupInfo) {

}

func (i *impl) QueryGroupAuthGroup(group model.Group) model.AuthGroupInfo {
	return model.AuthGroupInfo{}
}

func (i *impl) UpdateGroupAuthGroup(group model.Group, authGroupInfo model.AuthGroupInfo) {

}
