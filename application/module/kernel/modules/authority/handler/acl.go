package handler

import (
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
)

type aclManager struct {
	dbhelper           dbhelper.DBHelper
	getACLAuthGroup    map[string]*model.ACL
	postACLAuthGroup   map[string]*model.ACL
	putACLAuthGroup    map[string]*model.ACL
	deleteACLAuthGroup map[string]*model.ACL
	allACLAuthGroup    map[int]*model.ACL
}

func createACLManager(dbhelper dbhelper.DBHelper) aclManager {
	aclManager := aclManager{
		dbhelper:           dbhelper,
		getACLAuthGroup:    make(map[string]*model.ACL),
		postACLAuthGroup:   make(map[string]*model.ACL),
		putACLAuthGroup:    make(map[string]*model.ACL),
		deleteACLAuthGroup: make(map[string]*model.ACL),
		allACLAuthGroup:    make(map[int]*model.ACL)}
	aclManager.loadAllACL()

	return aclManager
}

func (i *aclManager) loadAllACL() bool {
	/*
		acls := dal.LoadACL(i.dbhelper, common.GET)
		for _, acl := range acls {
			cur := acl
			i.getACLAuthGroup[acl.URL] = &cur
			i.allACLAuthGroup[acl.ID] = &cur
		}
		acls = dal.LoadACL(i.dbhelper, common.POST)
		for _, acl := range acls {
			cur := acl
			i.postACLAuthGroup[acl.URL] = &cur
			i.allACLAuthGroup[acl.ID] = &cur
		}
		acls = dal.LoadACL(i.dbhelper, common.PUT)
		for _, acl := range acls {
			cur := acl
			i.putACLAuthGroup[acl.URL] = &cur
			i.allACLAuthGroup[acl.ID] = &cur
		}
		acls = dal.LoadACL(i.dbhelper, common.DELETE)
		for _, acl := range acls {
			cur := acl
			i.deleteACLAuthGroup[acl.URL] = &cur
			i.allACLAuthGroup[acl.ID] = &cur
		}*/
	return true
}

func (i *aclManager) queryACL(module string, status int) ([]model.ACL, bool) {
	/*
		if strings.ToLower(module) == "all" {
			return dal.QueryAllACL(i.dbhelper, status), true
		}

		return dal.QueryACL(i.dbhelper, module, status), true
	*/
	return []model.ACL{}, true
}

func (i *aclManager) addACL(url, method, module string) (model.ACL, bool) {
	/*
		acl, ok := dal.InsertACL(i.dbhelper, url, method, module, 0)
		if ok {
			switch method {
			case common.GET:
				i.getACLAuthGroup[url] = &acl
			case common.POST:
				i.postACLAuthGroup[url] = &acl
			case common.PUT:
				i.putACLAuthGroup[url] = &acl
			case common.DELETE:
				i.deleteACLAuthGroup[url] = &acl
			default:
				log.Printf("illegal method ,value:%s", method)
			}
			i.allACLAuthGroup[acl.ID] = &acl
		}

		return acl, true
	*/
	return model.ACL{}, true
}

func (i *aclManager) delACL(url, method, module string) bool {
	/*
		aclID := -1
		ok := false
		switch method {
		case common.GET:
			acl, ok := i.getACLAuthGroup[url]
			if ok {
				aclID = acl.ID
				ok = dal.DeleteACL(i.dbhelper, aclID)
				if ok {
					delete(i.getACLAuthGroup, url)
				}
			}
		case common.POST:
			acl, ok := i.postACLAuthGroup[url]
			if ok {
				aclID = acl.ID
				ok = dal.DeleteACL(i.dbhelper, aclID)
				if ok {
					delete(i.postACLAuthGroup, url)
				}
			}
		case common.PUT:
			acl, ok := i.putACLAuthGroup[url]
			if ok {
				aclID = acl.ID
				ok = dal.DeleteACL(i.dbhelper, aclID)
				if ok {
					delete(i.putACLAuthGroup, url)
				}
			}
		case common.DELETE:
			acl, ok := i.deleteACLAuthGroup[url]
			if ok {
				aclID = acl.ID
				ok = dal.DeleteACL(i.dbhelper, aclID)
				if ok {
					delete(i.deleteACLAuthGroup, url)
				}
			}
		default:
			log.Printf("illegal method ,value:%s", method)
		}
		if ok {
			delete(i.allACLAuthGroup, aclID)
		}
	*/
	return true
}

func (i *aclManager) enableACL(acls []int) bool {
	/*
		ok := dal.EnableACL(i.dbhelper, acls)
		if ok {
			for _, v := range acls {
				acl, found := i.allACLAuthGroup[v]
				if found {
					acl.Status = 1
				}
			}
		}

		return ok
	*/
	return true
}

func (i *aclManager) disableACL(acls []int) bool {
	/*
		ok := dal.DisableACL(i.dbhelper, acls)
		if ok {
			for _, v := range acls {
				acl, found := i.allACLAuthGroup[v]
				if found {
					acl.Status = 1
				}
			}
		}

		return ok
	*/
	return true
}

func (i *aclManager) adjustACLAuthGroup(aclID int, authGroup []int) (model.ACL, bool) {
	/*
		acl, found := i.allACLAuthGroup[aclID]
		if !found {
			return model.ACL{}, false
		}

		ok := false

		acl.AuthGroup = authGroup
		ok = dal.UpateACL(i.dbhelper, *acl)
		switch acl.Method {
		case common.GET:
			if ok {
				i.getACLAuthGroup[acl.URL] = acl
			}
		case common.POST:
			if ok {
				i.postACLAuthGroup[acl.URL] = acl
			}
		case common.PUT:
			if ok {
				i.putACLAuthGroup[acl.URL] = acl
			}
		case common.DELETE:
			if ok {
				i.deleteACLAuthGroup[acl.URL] = acl
			}
		default:
			log.Printf("illegal method ,value:%s", acl.Method)
		}

		return *acl, ok
	*/
	return model.ACL{}, true
}

func (i *aclManager) verifyAuthGroup(url, method string, authGroup []int) bool {
	/*
		var acl *model.ACL
		ok := false

		switch method {
		case common.GET:
			acl, ok = i.getACLAuthGroup[url]
		case common.POST:
			acl, ok = i.postACLAuthGroup[url]
		case common.PUT:
			acl, ok = i.putACLAuthGroup[url]
		case common.DELETE:
			acl, ok = i.deleteACLAuthGroup[url]
		default:
			log.Printf("illegal method ,value:%s", method)
		}

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
	*/
	return true
}
