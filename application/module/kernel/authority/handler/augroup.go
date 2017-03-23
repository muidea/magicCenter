package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/module/kernel/authority/dal"
)

type authGroupManager struct {
}

func createAuthGroupManager() authGroupManager {
	return authGroupManager{}
}

func (i *authGroupManager) adjustUserAuthGroup(userID int, authGroup []int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		return false
	}
	return dal.UpateUserAuthorityGroup(dbhelper, userID, authGroup)
}

func (i *authGroupManager) getUserAuthGroup(userID int) ([]int, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		return []int{}, false
	}

	return dal.GetUserAuthorityGroup(dbhelper, userID), true
}
