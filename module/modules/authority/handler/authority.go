package handler

import (
	"log"
	"net/http"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/common/dbhelper"
	"muidea.com/magicCenter/module/modules/authority/dal"
	common_const "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

// CreateAuthorityHandler 新建CASHandler
func CreateAuthorityHandler(moduleHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	i := impl{
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

	i.aclHandler = createACLHandler()
	i.aclHandler.loadACL()

	return &i
}

type impl struct {
	moduleHub       common.ModuleHub
	sessionRegistry common.SessionRegistry
	casHandler      common.CASHandler
	accountHandler  common.AccountHandler
	aclHandler      *aclHandler
}

func (i *impl) refreshUserStatus(session common.Session, remoteAddr string) {
	obj, ok := session.GetOption(common_const.AuthToken)
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

	acl, ok := i.aclHandler.filterACL(req)
	if !ok {
		log.Printf("can't find acl, pattern:%s, method:%s", req.URL.Path, req.Method)
		// 找不到对应的ACL，则认为没有权限
		return false
	}

	// 如果ACL的授权组为访客组，则直接认为有授权
	if acl.AuthGroup == common_const.VisitorAuthGroup.ID {
		return true
	}

	authToken := req.URL.Query().Get(common_const.AuthToken)
	if len(authToken) == 0 {
		// 没有提供AuthToken则认为没有授权
		log.Printf("illegal authToken, empty authToken value.")
		return false
	}

	session := i.sessionRegistry.GetSession(res, req)
	sessionToken, ok := session.GetOption(common_const.AuthToken)
	if ok {
		ok = sessionToken.(string) == authToken
		if !ok {
			log.Printf("illegal authToken, query authToken:%s, session authToken:%s", authToken, sessionToken.(string))
		}

		return ok
	}

	onlineEntry, _, ok := i.casHandler.VerifyToken(authToken)
	if !ok {
		// 如果提供了authToken，但是校验不通过，则认为没有权限
		log.Printf("invalid authToken, authToken:%s", authToken)
		return false
	}
	if onlineEntry.Unit.ID == common_const.SystemAccountUser.ID {
		return true
	}

	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	avalibleFlag := false
	moduleAuthGroup := dal.QueryUserModuleAuthGroup(dbhelper, onlineEntry.Unit.ID)
	for _, val := range moduleAuthGroup {
		if val.Module == acl.Module && val.AuthGroup >= acl.AuthGroup {
			avalibleFlag = true
			break
		}
	}

	return avalibleFlag
}

func (i *impl) QueryAllACL() []model.ACL {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryAllACL(dbhelper)
}

func (i *impl) QueryACLByModule(module string) []model.ACL {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryACLByModule(dbhelper, module)
}

func (i *impl) QueryACLByID(id int) (model.ACL, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryACLByID(dbhelper, id)
}

func (i *impl) InsertACL(url, method, module string, status int, authGroup int) (model.ACL, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.InsertACL(dbhelper, url, method, module, status, authGroup)
}

func (i *impl) DeleteACL(id int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.DeleteACL(dbhelper, id)
}

func (i *impl) UpdateACL(acl model.ACL) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.UpdateACL(dbhelper, acl)
}

func (i *impl) UpdateACLStatus(enableList []int, disableList []int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.UpdateACLStatus(dbhelper, enableList, disableList)
}

func (i *impl) QueryACLAuthGroup(id int) (model.AuthGroup, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	authGroup := model.AuthGroup{}
	acl, ok := dal.QueryACLByID(dbhelper, id)
	if !ok {
		return authGroup, ok
	}

	authGroup = common_const.GetAuthGroup(acl.AuthGroup)

	return authGroup, ok
}

func (i *impl) UpdateACLAuthGroup(id, authGroup int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	acl, ok := dal.QueryACLByID(dbhelper, id)
	if !ok {
		return ok
	}

	acl.AuthGroup = authGroup
	return dal.UpdateACL(dbhelper, acl)
}

func (i *impl) QueryAllModuleUser() []model.ModuleUserInfo {
	moduleUserInfos := []model.ModuleUserInfo{}

	ids := i.moduleHub.GetAllModuleIDs()
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	for _, v := range ids {
		info := model.ModuleUserInfo{Module: v}
		info.User = dal.QueryModuleUser(dbhelper, v)
		moduleUserInfos = append(moduleUserInfos, info)
	}

	return moduleUserInfos
}

func (i *impl) QueryModuleUserAuthGroup(module string) []model.UserAuthGroup {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryModuleUserAuthGroup(dbhelper, module)
}

func (i *impl) UpdateModuleUserAuthGroup(module string, userAuthGroups []model.UserAuthGroup) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.UpdateModuleUserAuthGroup(dbhelper, module, userAuthGroups)
}

func (i *impl) QueryAllUserModule() []model.UserModuleInfo {
	userModuleInfos := []model.UserModuleInfo{}

	allUsers := i.accountHandler.GetAllUserIDs()
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	for _, v := range allUsers {
		info := model.UserModuleInfo{User: v}
		info.Module = dal.QueryUserModule(dbhelper, v)

		userModuleInfos = append(userModuleInfos, info)
	}

	return userModuleInfos
}

func (i *impl) QueryUserModuleAuthGroup(user int) []model.ModuleAuthGroup {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.QueryUserModuleAuthGroup(dbhelper, user)
}

func (i *impl) UpdateUserModuleAuthGroup(user int, moduleAuthGroups []model.ModuleAuthGroup) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		panic(err)
	}
	defer dbhelper.Release()

	return dal.UpdateUserModuleAuthGroup(dbhelper, user, moduleAuthGroups)
}

func (i *impl) QueryUserACL(user int) []model.ACL {
	acls := []model.ACL{}

	//authGroup := dal.QueryUserModuleAuthGroup(dbhelper, user)
	//acls = dal.QueryAvalibleACLByAuthGroup(dbhelper, authGroup)

	return acls
}

func (i *impl) QueryAllAuthGroupDef() []model.AuthGroup {
	authGroups := []model.AuthGroup{}

	authGroups = append(authGroups, common_const.VisitorAuthGroup)
	authGroups = append(authGroups, common_const.UserAuthGroup)
	authGroups = append(authGroups, common_const.MaintainerAuthGroup)

	return authGroups
}
