package handler

import (
	"log"

	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/authority/dal"
)

type aclManager struct {
	aclAuthGroup map[string]model.ACL
}

func createACLManager() aclManager {
	aclManager := aclManager{aclAuthGroup: make(map[string]model.ACL)}
	aclManager.loadAllACL()

	return aclManager
}

func (i *aclManager) loadAllACL() bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	acls := dal.LoadACL(dbhelper)
	for _, acl := range acls {
		i.aclAuthGroup[acl.URL] = acl
	}

	return true
}

func (i *aclManager) queryACL(module string) ([]model.ACL, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return []model.ACL{}, false
	}

	return dal.QueryACL(dbhelper, module), true
}

func (i *aclManager) addACL(url, module string) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	acl, ok := dal.InsertACL(dbhelper, url, module)
	if ok {
		i.aclAuthGroup[url] = acl
	}

	return true
}

func (i *aclManager) delACL(url, module string) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	acl, ok := i.aclAuthGroup[url]
	if ok {
		ok = dal.DeleteACL(dbhelper, acl.ID)
		if ok {
			delete(i.aclAuthGroup, url)
		}
	}

	return ok
}

func (i *aclManager) adjustACLAuthGroup(url string, authGroup []int) bool {
	acl, ok := i.aclAuthGroup[url]
	if !ok {
		return ok
	}

	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}
	acl.AuthGroup = authGroup
	ok = dal.UpateACL(dbhelper, acl)
	if ok {
		i.aclAuthGroup[url] = acl
	}

	return ok
}

func (i *aclManager) verifyAuthGroup(url string, authGroup []int) bool {
	acl, ok := i.aclAuthGroup[url]
	if !ok {
		// 如果找不到对应的acl，则说明不需要权限，直接判定有权限
		return true
	}

	found := false
	for _, v := range acl.AuthGroup {
		for _, g := range authGroup {
			if v == g {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	return found
}
