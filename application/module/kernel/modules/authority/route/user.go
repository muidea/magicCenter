package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	"muidea.com/magicCenter/foundation/net"
)

// CreateGetUserModuleAuthGroupRoute 新建GetUserModuleAuthGroupRoute
func CreateGetUserModuleAuthGroupRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {
	i := userGetModuleAuthGroupRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

// CreatePutUserModuleAuthGroupRoute 新建UserModulePutRoute
func CreatePutUserModuleAuthGroupRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {
	i := userPutModuleAuthGroupRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetUserACLRoute 新建UserACLGetRoute
func CreateGetUserACLRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {
	i := userGetACLRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

type userGetModuleAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type userGetModuleAuthGroupResult struct {
	common.Result
	model.UserModuleAuthGroupView
}

func (i *userGetModuleAuthGroupRoute) Method() string {
	return common.GET
}

func (i *userGetModuleAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetUserModuleAuthGroup)
}

func (i *userGetModuleAuthGroupRoute) Handler() interface{} {
	return i.getUserModuleAuthGroupHandler
}

func (i *userGetModuleAuthGroupRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *userGetModuleAuthGroupRoute) getUserModuleAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserModuleAuthGroupHandler")

	result := userGetModuleAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		val := i.authorityHandler.QueryUserModuleAuthGroup(id)
		for _, v := range val.ModuleAuthGroup {
			group := model.ModuleAuthGroupView{}
			group.ModuleAuthGroup = v

			if v.AuthGroup == common.UserAuthGroup.ID {
				group.AuthGroup = common.UserAuthGroup.Unit
			} else if v.AuthGroup == common.MaintainerAuthGroup.ID {
				group.AuthGroup = common.MaintainerAuthGroup.Unit
			} else {
				group.AuthGroup = common.VisitorAuthGroup.Unit
			}

			result.ModuleAuthGroup = append(result.ModuleAuthGroup, group)
		}

		user, _ := i.accountHandler.FindUserByID(val.User)
		result.User.ID = user.ID
		result.User.Name = user.Name

		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userPutModuleAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type userPutModuleAuthGroupParam struct {
	ModuleAuthGroup []model.ModuleAuthGroup
}

type userPutModuleAuthGroupResult struct {
	common.Result
}

func (i *userPutModuleAuthGroupRoute) Method() string {
	return common.PUT
}

func (i *userPutModuleAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutUserModuleAuthGroup)
}

func (i *userPutModuleAuthGroupRoute) Handler() interface{} {
	return i.putUserModuleAuthGroupHandler
}

func (i *userPutModuleAuthGroupRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *userPutModuleAuthGroupRoute) putUserModuleAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putUserModuleAuthGroupHandler")

	result := userPutModuleAuthGroupResult{}
	for true {
		param := &userPutModuleAuthGroupParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "非法参数"
			break
		}
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateUserModuleAuthGroup(id, param.ModuleAuthGroup)
		if ok {
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
			result.Reason = "更新用户管理模块信息失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userGetACLRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type userGetACLResult struct {
	common.Result
	User model.User            `json:"user"`
	ACL  []model.ACLDetailView `json:"acl"`
}

func (i *userGetACLRoute) Method() string {
	return common.GET
}

func (i *userGetACLRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetUserACL)
}

func (i *userGetACLRoute) Handler() interface{} {
	return i.getUserACLHandler
}

func (i *userGetACLRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *userGetACLRoute) getUserACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserACLHandler")

	result := userGetACLResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		acls := i.authorityHandler.QueryUserACL(id)
		for _, val := range acls {
			acl := model.ACLDetailView{}
			acl.ACLDetail = val

			result.ACL = append(result.ACL, acl)
		}

		user, _ := i.accountHandler.FindUserByID(id)
		result.User.ID = user.ID
		result.User.Name = user.Name

		result.ErrorCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
