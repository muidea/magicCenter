package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/dal"
	"muidea.com/magicCenter/foundation/net"
)

// CreateAuthorityHandler 新建CASHandler
func CreateAuthorityHandler(moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	dbhelper, _ := dbhelper.NewHelper()

	i := impl{
		dbhelper:        dbhelper,
		sessionRegistry: sessionRegistry}

	casModule, _ := moduleHub.FindModule(common.CASModuleID)
	entryPoint := casModule.EntryPoint()
	switch entryPoint.(type) {
	case common.CASHandler:
		i.casHandler = entryPoint.(common.CASHandler)
	default:
		panic("can\\'t find CASModule")
	}

	return &i
}

type impl struct {
	dbhelper        dbhelper.DBHelper
	sessionRegistry common.SessionRegistry
	casHandler      common.CASHandler
}

func (i *impl) refreshUserStatus(session common.Session, remoteAddr string) {
	obj, ok := session.GetOption(common.AuthTokenID)
	if !ok {
		panic("")
	}

	// 找到sessionToken了，则说明该用户已经登录了，这里就必须保证两端的token一致否则也要认为鉴权非法
	// 用户登录过Token必然不为空
	// req.RemoteAddr
	sessionToken := obj.(string)

	if i.casHandler != nil {
		i.casHandler.RefreshToken(sessionToken, remoteAddr)
	}
}

/*
1、先获取当前route对应的授权组
*/
func (i *impl) VerifyAuthority(res http.ResponseWriter, req *http.Request) bool {
	url, id := net.SplitRESTAPI(req.URL.Path)
	urlPattern := net.FormatRoutePattern(url, id)
	urlMethod := req.Method

	acl, ok := dal.QueryACL(i.dbhelper, urlPattern, urlMethod)
	if !ok {
		// 找不到ACL，当成没有权限来处理
		return false
	}

	retVal := false
	for _, cur := range acl.AuthGroup {
		if cur == common.VisitorAuthGroup.ID {
			// 如果当前URL的授权组包含VisitorAuthGroup，则直接认为有授权
			retVal = true
			break
		}
	}

	session := i.sessionRegistry.GetSession(res, req)
	user, loginOK := session.GetAccount()
	if loginOK {
		i.refreshUserStatus(session, req.RemoteAddr)
	}

	if retVal || !loginOK {
		return retVal
	}

	authGroups := dal.QueryUserAuthGroup(i.dbhelper, user.ID)
	for _, cur := range acl.AuthGroup {
		for _, item := range authGroups {
			if cur == item {
				retVal = true
				break
			}
		}

		if retVal {
			break
		}
	}

	return retVal
}

func (i *impl) QueryModuleACL(module string) []model.ACL {
	return dal.QueryModuleACL(i.dbhelper, module)
}

func (i *impl) QueryACL(url, method string) (model.ACL, bool) {
	return dal.QueryACL(i.dbhelper, url, method)
}

func (i *impl) InsertACL(url, method, module string, status int, authGroups []int) (model.ACL, bool) {
	return dal.InsertACL(i.dbhelper, url, method, module, status, authGroups)
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
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return authGroups
	}

	return acl.AuthGroup
}

func (i *impl) UpdateACLAuthGroup(id int, authGroups []int) bool {
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
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
