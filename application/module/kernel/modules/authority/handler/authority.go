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

	acl, ok := dal.FilterACL(i.dbhelper, urlPattern)
	if !ok {
		// 找不到ACL，当成没有权限来处理
		return false
	}

	if acl.AuthGroup == common.VisitorAuthGroup.ID {
		return true
	}

	retVal := false
	// 到这里就说明必须要求访问用户要求属于UserAuthGroup或者MaintainerAuthGroup
	// 这里这里需要判断token是否合法
	session := i.sessionRegistry.GetSession(res, req)
	urlToken := req.URL.Query().Get(common.AuthTokenID)
	sessionToken, ok := session.GetOption(common.AuthTokenID)
	if !ok || sessionToken.(string) != urlToken {
		// 如果用户没有登录，或者urlToken与sessionToken不一致，则说明权限非法
		return false
	}

	_, loginOK := session.GetAccount()
	if loginOK {
		i.refreshUserStatus(session, req.RemoteAddr)
	}

	if retVal || !loginOK {
		return retVal
	}

	/*
		authGroup := dal.QueryUserModuleAuthGroup(i.dbhelper, user.ID)
		if authGroup >= acl.AuthGroup {
			retVal = true
		} else {
			retVal = false
		}
	*/

	return retVal
}

func (i *impl) QueryACLByModule(module string) []model.ACL {
	return dal.QueryACLByModule(i.dbhelper, module)
}

func (i *impl) QueryACLByID(id int) (model.ACL, bool) {
	return dal.QueryACLByID(i.dbhelper, id)
}

func (i *impl) InsertACL(url, method, module string, status int, authGroup int) (model.ACL, bool) {
	return dal.InsertACL(i.dbhelper, url, method, module, status, authGroup)
}

func (i *impl) DeleteACL(id int) bool {
	return dal.DeleteACL(i.dbhelper, id)
}

func (i *impl) UpdateACLStatus(enableList []int, disableList []int) bool {
	return dal.UpdateACLStatus(i.dbhelper, enableList, disableList)
}

func (i *impl) QueryACLAuthGroup(id int) (int, bool) {
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return 0, ok
	}

	return acl.AuthGroup, ok
}

func (i *impl) UpdateACLAuthGroup(id, authGroup int) bool {
	acl, ok := dal.QueryACLByID(i.dbhelper, id)
	if !ok {
		return ok
	}

	acl.AuthGroup = authGroup
	return dal.UpateACL(i.dbhelper, acl)
}

func (i *impl) QueryUserModuleAuthGroup(user int) model.UserModuleAuthGroupInfo {
	return dal.QueryUserModuleAuthGroup(i.dbhelper, user)
}

func (i *impl) UpdateUserModuleAuthGroup(user int, moduleAuthGroups []model.ModuleAuthGroup) bool {
	return dal.UpdateUserModuleAuthGroup(i.dbhelper, user, moduleAuthGroups)
}

func (i *impl) QueryUserACL(user int) []model.ACL {
	acls := []model.ACL{}

	//authGroup := dal.QueryUserModuleAuthGroup(i.dbhelper, user)
	//acls = dal.QueryAvalibleACLByAuthGroup(i.dbhelper, authGroup)

	return acls
}

func (i *impl) QueryModuleUserAuthGroup(module string) model.ModuleUserAuthGroupInfo {
	return dal.QueryModuleUserAuthGroup(i.dbhelper, module)
}

func (i *impl) UpdateModuleUserAuthGroup(module string, userAuthGroups []model.UserAuthGroup) bool {
	return dal.UpdateModuleUserAuthGroup(i.dbhelper, module, userAuthGroups)
}
