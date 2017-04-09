package handler

import (
	"log"

	"strings"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/dbhelper"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/dal"
)

type aclManager struct {
	getACLAuthGroup    map[string]model.ACL
	postACLAuthGroup   map[string]model.ACL
	putACLAuthGroup    map[string]model.ACL
	deleteACLAuthGroup map[string]model.ACL
	allACLAuthGroup    map[int]*model.ACL
}

func createACLManager() aclManager {
	aclManager := aclManager{
		getACLAuthGroup:    make(map[string]model.ACL),
		postACLAuthGroup:   make(map[string]model.ACL),
		putACLAuthGroup:    make(map[string]model.ACL),
		deleteACLAuthGroup: make(map[string]model.ACL),
		allACLAuthGroup:    make(map[int]*model.ACL)}
	aclManager.loadAllACL()

	return aclManager
}

func (i *aclManager) loadAllACL() bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	acls := dal.LoadACL(dbhelper, common.GET)
	for _, acl := range acls {
		i.getACLAuthGroup[acl.URL] = acl
		i.allACLAuthGroup[acl.ID] = &acl
	}
	acls = dal.LoadACL(dbhelper, common.POST)
	for _, acl := range acls {
		i.postACLAuthGroup[acl.URL] = acl
		i.allACLAuthGroup[acl.ID] = &acl
	}
	acls = dal.LoadACL(dbhelper, common.PUT)
	for _, acl := range acls {
		i.putACLAuthGroup[acl.URL] = acl
		i.allACLAuthGroup[acl.ID] = &acl
	}
	acls = dal.LoadACL(dbhelper, common.DELETE)
	for _, acl := range acls {
		i.deleteACLAuthGroup[acl.URL] = acl
		i.allACLAuthGroup[acl.ID] = &acl
	}
	return true
}

func (i *aclManager) queryACL(module string, status int) ([]model.ACL, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return []model.ACL{}, false
	}

	if strings.ToLower(module) == "all" {
		return dal.QueryAllACL(dbhelper, status), true
	}

	return dal.QueryACL(dbhelper, module, status), true
}

func (i *aclManager) addACL(url, method, module string) (model.ACL, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return model.ACL{}, false
	}

	acl, ok := dal.InsertACL(dbhelper, url, method, module, 0)
	if ok {
		switch method {
		case common.GET:
			i.getACLAuthGroup[url] = acl
		case common.POST:
			i.postACLAuthGroup[url] = acl
		case common.PUT:
			i.putACLAuthGroup[url] = acl
		case common.DELETE:
			i.deleteACLAuthGroup[url] = acl
		default:
			log.Printf("illegal method ,value:%s", method)
		}
		i.allACLAuthGroup[acl.ID] = &acl
	}

	return acl, true
}

func (i *aclManager) delACL(url, method, module string) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	acl := model.ACL{}
	ok := false
	switch method {
	case common.GET:
		acl, ok = i.getACLAuthGroup[url]
		if ok {
			ok = dal.DeleteACL(dbhelper, acl.ID)
			if ok {
				delete(i.getACLAuthGroup, url)
			}
		}
	case common.POST:
		acl, ok = i.postACLAuthGroup[url]
		if ok {
			ok = dal.DeleteACL(dbhelper, acl.ID)
			if ok {
				delete(i.postACLAuthGroup, url)
			}
		}
	case common.PUT:
		acl, ok = i.putACLAuthGroup[url]
		if ok {
			ok = dal.DeleteACL(dbhelper, acl.ID)
			if ok {
				delete(i.putACLAuthGroup, url)
			}
		}
	case common.DELETE:
		acl, ok = i.deleteACLAuthGroup[url]
		if ok {
			ok = dal.DeleteACL(dbhelper, acl.ID)
			if ok {
				delete(i.deleteACLAuthGroup, url)
			}
		}
	default:
		log.Printf("illegal method ,value:%s", method)
	}
	if ok {
		delete(i.allACLAuthGroup, acl.ID)
	}
	return ok
}

func (i *aclManager) enableACL(acls []int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	ok := dal.EnableACL(dbhelper, acls)
	if ok {
		for _, v := range acls {
			acl, found := i.allACLAuthGroup[v]
			if found {
				acl.Status = 1
			}
		}
	}

	return ok
}

func (i *aclManager) disableACL(acls []int) bool {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return false
	}

	ok := dal.DisableACL(dbhelper, acls)
	if ok {
		for _, v := range acls {
			acl, found := i.allACLAuthGroup[v]
			if found {
				acl.Status = 1
			}
		}
	}

	return ok
}

func (i *aclManager) adjustACLAuthGroup(aclID int, authGroup []int) (model.ACL, bool) {
	dbhelper, err := dbhelper.NewHelper()
	if err != nil {
		log.Println("create new dbhelper failed")
		return model.ACL{}, false
	}

	acl, found := i.allACLAuthGroup[aclID]
	if !found {
		return model.ACL{}, false
	}

	ok := false

	acl.AuthGroup = authGroup
	ok = dal.UpateACL(dbhelper, *acl)
	switch acl.Method {
	case common.GET:
		if ok {
			i.getACLAuthGroup[acl.URL] = *acl
		}
	case common.POST:
		if ok {
			i.postACLAuthGroup[acl.URL] = *acl
		}
	case common.PUT:
		if ok {
			i.putACLAuthGroup[acl.URL] = *acl
		}
	case common.DELETE:
		if ok {
			i.deleteACLAuthGroup[acl.URL] = *acl
		}
	default:
		log.Printf("illegal method ,value:%s", acl.Method)
	}

	return *acl, ok
}

func (i *aclManager) verifyAuthGroup(url, method string, authGroup []int) bool {
	acl := model.ACL{}
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
}
