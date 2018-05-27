package handler

import (
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/dal"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/model"
)

// CreateAuthorityHandler 新建CASHandler
func CreateAuthorityHandler(moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	dbhelper, _ := dbhelper.NewHelper()

	i := impl{
		dbhelper:        dbhelper,
		moduleHub:       moduleHub,
		sessionRegistry: sessionRegistry}

	casModule, _ := moduleHub.FindModule(common.CASModuleID)
	entryPoint := casModule.EntryPoint()
	switch entryPoint.(type) {
	case common.CASHandler:
		i.casHandler = entryPoint.(common.CASHandler)
	default:
		panic("can\\'t find CASModule")
	}

	accountModule, _ := moduleHub.FindModule(common.AccountModuleID)
	entryPoint = accountModule.EntryPoint()
	i.accountHandler = entryPoint.(common.AccountHandler)

	return &i
}

type impl struct {
	dbhelper        dbhelper.DBHelper
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	casHandler      common.CASHandler
	accountHandler  common.AccountHandler
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
	path, id := net.SplitRESTAPI(req.URL.Path)
	pattern := net.FormatRoutePattern(path, id)
	method := req.Method

	acl, ok := dal.FilterACL(i.dbhelper, pattern, method)
	if !ok {
		// 找不到对应的ACL，则认为没有权限
		return false
	}

	// 如果ACL的授权组为访客组，则直接认为有授权
	if acl.AuthGroup == common_const.VisitorAuthGroup.ID {
		return true
	}

	authToken := req.URL.Query().Get(common.AuthTokenID)
	if len(authToken) == 0 {
		// 没有提供AuthToken则认为没有授权
		return false
	}

	session := i.sessionRegistry.GetSession(res, req)
	sessionToken, ok := session.GetOption(common.AuthTokenID)
	if ok {
		return sessionToken.(string) == authToken
	}

	onlineAccount, ok := i.casHandler.VerifyToken(authToken)
	if !ok {
		// 如果提供了authToken，但是校验不通过，则认为没有权限
		return false
	}

	avalibleFlag := false
	moduleAuthGroup := dal.QueryUserModuleAuthGroup(i.dbhelper, onlineAccount.User.ID)
	for _, val := range moduleAuthGroup {
		if val.Module == acl.Module && val.AuthGroup >= acl.AuthGroup {
			avalibleFlag = true
			break
		}
	}

	if avalibleFlag {
		// 如果校验通过，则更新session里的相关信息
		session.SetAccount(onlineAccount.User)
		session.SetOption(common.AuthTokenID, authToken)
	}

	return avalibleFlag
}

func (i *impl) QueryAllACL() []model.ACL {
	return dal.QueryAllACL(i.dbhelper)
}

func (i *impl) QueryACLByModule(module string) []model.ACL {
	return dal.QueryACLByModule(i.dbhelper, module)
}

func (i *impl) QueryACLByID(id int) (model.ACLDetail, bool) {
	return dal.QueryACLByID(i.dbhelper, id)
}

func (i *impl) InsertACL(url, method, module string, status int, authGroup int) (model.ACLDetail, bool) {
	return dal.InsertACL(i.dbhelper, url, method, module, status, authGroup)
}

func (i *impl) DeleteACL(id int) bool {
	return dal.DeleteACL(i.dbhelper, id)
}

func (i *impl) UpdateACL(acl model.ACLDetail) bool {
	return dal.UpdateACL(i.dbhelper, acl)
}

func (i *impl) UpdateACLStatus(enableList []int, disableList []int) bool {
	return dal.UpdateACLStatus(i.dbhelper, enableList, disableList)
}

func (i *impl) QueryACLAuthGroup(id int) (model.AuthGroup, bool) {
	authGroup := model.AuthGroup{}
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return authGroup, ok
	}

	authGroup = common_const.GetAuthGroup(acl.AuthGroup)

	return authGroup, ok
}

func (i *impl) UpdateACLAuthGroup(id, authGroup int) bool {
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return ok
	}

	acl.AuthGroup = authGroup
	return dal.UpdateACL(i.dbhelper, acl)
}

func (i *impl) QueryAllModuleUser() []model.ModuleUserInfo {
	moduleUserInfos := []model.ModuleUserInfo{}

	ids := i.moduleHub.GetAllModuleIDs()
	for _, v := range ids {
		info := model.ModuleUserInfo{Module: v}
		info.User = dal.QueryModuleUser(i.dbhelper, v)
		moduleUserInfos = append(moduleUserInfos, info)
	}

	return moduleUserInfos
}

func (i *impl) QueryModuleUserAuthGroup(module string) []model.UserAuthGroup {
	return dal.QueryModuleUserAuthGroup(i.dbhelper, module)
}

func (i *impl) UpdateModuleUserAuthGroup(module string, userAuthGroups []model.UserAuthGroup) bool {
	return dal.UpdateModuleUserAuthGroup(i.dbhelper, module, userAuthGroups)
}

func (i *impl) QueryAllUserModule() []model.UserModuleInfo {
	userModuleInfos := []model.UserModuleInfo{}

	allUsers := i.accountHandler.GetAllUserIDs()
	for _, v := range allUsers {
		info := model.UserModuleInfo{User: v}
		info.Module = dal.QueryUserModule(i.dbhelper, v)

		userModuleInfos = append(userModuleInfos, info)
	}

	return userModuleInfos
}

func (i *impl) QueryUserModuleAuthGroup(user int) []model.ModuleAuthGroup {
	return dal.QueryUserModuleAuthGroup(i.dbhelper, user)
}

func (i *impl) UpdateUserModuleAuthGroup(user int, moduleAuthGroups []model.ModuleAuthGroup) bool {
	return dal.UpdateUserModuleAuthGroup(i.dbhelper, user, moduleAuthGroups)
}

func (i *impl) QueryAllEndpoint() []model.Endpoint {
	return dal.QueryAllEndpoint(i.dbhelper)
}

func (i *impl) QueryEndpointByID(id string) (model.Endpoint, bool) {
	return dal.QueryEndpointByID(i.dbhelper, id)
}

func (i *impl) InsertEndpoint(id, name, description string, user []int, status int, authToken string) (model.Endpoint, bool) {
	return dal.InsertEndpoint(i.dbhelper, id, name, description, user, status, authToken)
}

func (i *impl) UpdateEndpoint(endpoint model.Endpoint) (model.Endpoint, bool) {
	return dal.UpdateEndpoint(i.dbhelper, endpoint)
}

func (i *impl) DeleteEndpoint(id string) bool {
	return dal.DeleteEndpoint(i.dbhelper, id)
}

func (i *impl) QueryUserACL(user int) []model.ACLDetail {
	acls := []model.ACLDetail{}

	//authGroup := dal.QueryUserModuleAuthGroup(i.dbhelper, user)
	//acls = dal.QueryAvalibleACLByAuthGroup(i.dbhelper, authGroup)

	return acls
}

func (i *impl) QueryAllAuthGroupDef() []model.AuthGroup {
	authGroups := []model.AuthGroup{}

	authGroups = append(authGroups, common_const.VisitorAuthGroup)
	authGroups = append(authGroups, common_const.UserAuthGroup)
	authGroups = append(authGroups, common_const.MaintainerAuthGroup)

	return authGroups
}
