package handler

import (
	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
)

// CreateAuthorityHandler 新建CASHandler
func CreateAuthorityHandler(modHub common.ModuleHub, sessionRegistry common.SessionRegistry) common.AuthorityHandler {
	dbhelper, _ := dbhelper.NewHelper()
	i := impl{
		sessionRegistry:  sessionRegistry,
		authGroupManager: createAuthGroupManager(dbhelper),
		aclManager:       createACLManager(dbhelper),
		cacheData:        cache.NewCache()}

	return &i
}

type impl struct {
	sessionRegistry  common.SessionRegistry
	authGroupManager authGroupManager
	aclManager       aclManager
	cacheData        cache.Cache
}

func (i *impl) QueryAuthGroup(module string) ([]model.AuthGroup, bool) {
	return i.authGroupManager.queryAuthGroup(module)
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

func (i *impl) QueryACL(module string, status int) ([]model.ACL, bool) {
	return i.aclManager.queryACL(module, status)
}

func (i *impl) AddACL(url, method, module string) (model.ACL, bool) {
	return i.aclManager.addACL(url, method, module)
}

func (i *impl) DelACL(url, method, module string) bool {
	return i.aclManager.delACL(url, method, module)
}

func (i *impl) EnableACL(acls []int) bool {
	return i.aclManager.enableACL(acls)
}

func (i *impl) DisableACL(acls []int) bool {
	return i.aclManager.disableACL(acls)
}

func (i *impl) AdjustACLAuthGroup(acl int, authGroup []int) (model.ACL, bool) {
	return i.aclManager.adjustACLAuthGroup(acl, authGroup)
}