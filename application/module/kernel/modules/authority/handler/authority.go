package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/dal"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateAuthorityHandler 新建CASHandler
func CreateAuthorityHandler(moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	dbhelper, _ := dbhelper.NewHelper()

	i := impl{
		dbhelper:        dbhelper,
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
	dbhelper        dbhelper.DBHelper
	sessionRegistry common.SessionRegistry
	casHandler      common.CASHandler
	cacheData       cache.Cache
}

/*
1、先判断authToken是否一致，如果不一致则，认为无权限
*/
func (i *impl) VerifyAuthority(res http.ResponseWriter, req *http.Request) bool {
	//url := req.URL.Path

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

func (i *impl) QueryModuleACL(module string) []model.ACL {
	return dal.QueryModuleACL(i.dbhelper, module)
}

func (i *impl) InsertACL(url, method, module string, status int) (model.ACL, bool) {
	return dal.InsertACL(i.dbhelper, url, method, module, status)
}

func (i *impl) DeleteACL(id int) bool {
	return dal.DeleteACL(i.dbhelper, id)
}

func (i *impl) EnableACL(ids []int) bool {
	return dal.EnableACL(i.dbhelper, ids)
}

func (i *impl) DisableACL(ids []int) bool {
	return dal.DisableACL(i.dbhelper, ids)
}

func (i *impl) QueryACLAuthGroup(id int) []int {
	authGroups := []int{}
	acl, ok := dal.QueryACL(i.dbhelper, id)
	if !ok {
		return authGroups
	}

	return acl.AuthGroup
}

func (i *impl) UpdateACLAuthGroup(id int, authGroups []int) bool {
	acl, ok := dal.QueryACL(i.dbhelper, id)
	if !ok {
		return ok
	}

	acl.AuthGroup = authGroups
	return dal.UpateACL(i.dbhelper, acl)
}

func (i *impl) QueryUserAuthGroup(user int) []int {
	return dal.QueryUserAuthGroup(i.dbhelper, user)
}

func (i *impl) UpdateUserAuthGroup(user int, authGroups []int) bool {
	return dal.UpdateUserAuthGroup(i.dbhelper, user, authGroups)
}

func (i *impl) QueryUserModule(user int) []string {
	return dal.QueryUserModule(i.dbhelper, user)
}

func (i *impl) UpdateUserModule(user int, modules []string) bool {
	return dal.UpdateUserModule(i.dbhelper, user, modules)
}
