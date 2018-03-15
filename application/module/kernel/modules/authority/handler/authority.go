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
	path, id := net.SplitRESTAPI(req.URL.Path)
	pattern := net.FormatRoutePattern(path, id)
	method := req.Method

	acl, ok := dal.FilterACL(i.dbhelper, pattern, method)
	if !ok {
		// 找不到对应的ACL，则认为没有权限
		return false
	}

	// 如果ACL的授权组为访客组，则直接认为有授权
	if acl.AuthGroup == common.VisitorAuthGroup.ID {
		return true
	}

	authToken := req.URL.Query().Get("authToken")
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
	userModuleAuthGroup := dal.QueryUserModuleAuthGroup(i.dbhelper, onlineAccount.User.ID)
	for _, val := range userModuleAuthGroup.ModuleAuthGroup {
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

func (i *impl) UpdateACLStatus(enableList []int, disableList []int) bool {
	return dal.UpdateACLStatus(i.dbhelper, enableList, disableList)
}

func (i *impl) QueryACLAuthGroup(id int) (model.AuthGroup, bool) {
	authGroup := model.AuthGroup{}
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return authGroup, ok
	}

	switch acl.AuthGroup {
	case common.VisitorAuthGroup.ID:
		authGroup = common.VisitorAuthGroup
	case common.UserAuthGroup.ID:
		authGroup = common.UserAuthGroup
	case common.MaintainerAuthGroup.ID:
		authGroup = common.MaintainerAuthGroup
	}

	return authGroup, ok
}

func (i *impl) UpdateACLAuthGroup(id, authGroup int) bool {
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return ok
	}

	acl.AuthGroup = authGroup
	return dal.UpateACL(i.dbhelper, acl)
}

func (i *impl) QueryUserModuleAuthGroup(user int) model.UserModuleAuthGroup {
	return dal.QueryUserModuleAuthGroup(i.dbhelper, user)
}

func (i *impl) UpdateUserModuleAuthGroup(user int, moduleAuthGroups []model.ModuleAuthGroup) bool {
	return dal.UpdateUserModuleAuthGroup(i.dbhelper, user, moduleAuthGroups)
}

func (i *impl) QueryUserACL(user int) []model.ACLDetail {
	acls := []model.ACLDetail{}

	//authGroup := dal.QueryUserModuleAuthGroup(i.dbhelper, user)
	//acls = dal.QueryAvalibleACLByAuthGroup(i.dbhelper, authGroup)

	return acls
}

func (i *impl) QueryModuleUserAuthGroup(module string) model.ModuleUserAuthGroup {
	return dal.QueryModuleUserAuthGroup(i.dbhelper, module)
}

func (i *impl) UpdateModuleUserAuthGroup(module string, userAuthGroups []model.UserAuthGroup) bool {
	return dal.UpdateModuleUserAuthGroup(i.dbhelper, module, userAuthGroups)
}
