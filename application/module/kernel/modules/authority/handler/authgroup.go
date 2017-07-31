package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

type authGroupManager struct {
	dbhelper         dbhelper.DBHelper
	module2AuthGroup map[string][]model.AuthGroup
}

func createAuthGroupManager(dbhelper dbhelper.DBHelper) authGroupManager {
	authGroupManager := authGroupManager{dbhelper: dbhelper, module2AuthGroup: make(map[string][]model.AuthGroup)}
	authGroupManager.loadAllAuthGroup()

	return authGroupManager
}

func (i *authGroupManager) loadAllAuthGroup() bool {
	/*
		authGroups := dal.GetAllAuthGroup(i.dbhelper)
		for _, authGroup := range authGroups {
			authGroups, found := i.module2AuthGroup[authGroup.Module]
			if !found {
				authGroups = []model.AuthGroup{}
			}

			authGroups = append(authGroups, authGroup)
			i.module2AuthGroup[authGroup.Module] = authGroups
		}
	*/
	return true
}

func (i *authGroupManager) queryAuthGroup(module string) ([]model.AuthGroup, bool) {
	/*
		if strings.ToLower(module) == "all" {
			authGroups := []model.AuthGroup{}
			for _, groups := range i.module2AuthGroup {
				authGroups = append(authGroups, groups...)
			}

			return authGroups, true
		}

		authGroups, found := i.module2AuthGroup[module]
		return authGroups, found
	*/
	return []model.AuthGroup{}, true
}

func (i *authGroupManager) insertAuthGroup(authGroups []model.AuthGroup) bool {
	/*
		for _, authGroup := range authGroups {
			authGroup, ok := dal.InsertAuthGroup(i.dbhelper, authGroup)
			if ok {
				authGroups, found := i.module2AuthGroup[authGroup.Module]
				if !found {
					authGroups = []model.AuthGroup{}
				}

				authGroups = append(authGroups, authGroup)
				i.module2AuthGroup[authGroup.Module] = authGroups
			}
		}
	*/
	return true
}

func (i *authGroupManager) deleteAuthGroup(authGroups []model.AuthGroup) bool {
	/*
		for _, v := range authGroups {
			dal.DeleteAuthGroup(i.dbhelper, v.ID)

			curAuthGroups, found := i.module2AuthGroup[v.Module]
			newAuthGroups := []model.AuthGroup{}
			if found {
				for _, c := range curAuthGroups {
					if c.ID != v.ID {
						newAuthGroups = append(newAuthGroups, c)
					}
				}
				if len(newAuthGroups) > 0 {
					i.module2AuthGroup[v.Module] = newAuthGroups
				}
			}
		}
	*/
	return true
}

func (i *authGroupManager) adjustUserAuthGroup(userID int, authGroup []int) bool {
	//return dal.UpateUserAuthorityGroup(i.dbhelper, userID, authGroup
	return true

}

func (i *authGroupManager) getUserAuthGroup(userID int) ([]int, bool) {
	//return dal.GetUserAuthorityGroup(i.dbhelper, userID), true
	return nil, true

}
