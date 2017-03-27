package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
	"muidea.com/magicCenter/foundation/net"
)

// CreateAuthorityHandler 新建AuthorityHandler
func CreateAuthorityHandler(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	i := impl{
		sessionRegistry:  sessionRegistry,
		accountManager:   createAccountManager(modHub),
		authGroupManager: createAuthGroupManager(),
		aclManager:       createACLManager(),
		cacheData:        cache.NewCache()}

	return &i
}

type impl struct {
	sessionRegistry  common.SessionRegistry
	accountManager   accountManager
	authGroupManager authGroupManager
	aclManager       aclManager
	cacheData        cache.Cache
}

func (i *impl) LoginAccount(account, password string) (model.UserDetail, string, bool) {
	user, ok := i.accountManager.findUser(account, password)
	if !ok {
		return user, "", ok
	}

	token := i.cacheData.PutIn(user, cache.MaxAgeValue)
	return user, token, ok
}

func (i *impl) LogoutAccount(authToken string) bool {
	_, ok := i.cacheData.FetchOut(authToken)
	if !ok {
		return false
	}

	i.cacheData.Remove(authToken)
	return true
}

/*
1、先判断authToken是否一致，如果不一致则，认为无权限
*/
func (i *impl) VerifyAuth(res http.ResponseWriter, req *http.Request) bool {
	session := i.sessionRegistry.GetSession(res, req)

	authGroup := []int{}
	url := req.URL.Path + ":" + req.Method
	user, ok := session.GetAccount()
	if ok {
		// Session里找不到用户则说明用户没有登录
		authGroup, ok = i.authGroupManager.getUserAuthGroup(user.ID)
		if !ok {
			return false
		}
	}
	if !i.aclManager.verifyAuthGroup(url, authGroup) {
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

func (i *impl) GetAuthGroup(module string) ([]model.AuthGroup, bool) {
	return i.authGroupManager.findAuthGroup(module)
}

func (i *impl) InsertAuthGroup(authGroups []model.AuthGroup) bool {
	return i.authGroupManager.insertAuthGroup(authGroups)
}

func (i *impl) DeleteAuthGroup(authGroups []model.AuthGroup) bool {
	return i.authGroupManager.deleteAuthGroup(authGroups)
}

func (i *impl) AdjustUserAuthGroup(userID int, authGroup []int) bool {
	return i.authGroupManager.adjustUserAuthGroup(userID, authGroup)
}

func (i *impl) GetUserAuthGroup(userID int) ([]int, bool) {
	return i.authGroupManager.getUserAuthGroup(userID)
}

func (i *impl) QueryACL(module string) ([]model.ACL, bool) {
	return i.aclManager.queryACL(module)
}

func (i *impl) AddACL(url, module string, route common.Route) bool {
	url = net.JoinURL(url, route.Pattern()) + ":" + route.Method()

	return i.aclManager.addACL(url, module)
}

func (i *impl) DelACL(url, module string, route common.Route) bool {
	url = net.JoinURL(url, route.Pattern()) + ":" + route.Method()

	return i.aclManager.delACL(url, module)
}

func (i *impl) AdjustACLAuthGroup(moduleURL string, route common.Route, authGroup []int) bool {
	url := net.JoinURL(moduleURL, route.Pattern()) + ":" + route.Method()

	return i.aclManager.adjustACLAuthGroup(url, authGroup)
}
