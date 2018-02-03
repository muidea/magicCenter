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
func CreateGetUserModuleAuthGroupRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userGetModuleAuthGroupRoute{authorityHandler: authorityHandler}
	return &i
}

// CreatePutUserModuleAuthGroupRoute 新建UserModulePutRoute
func CreatePutUserModuleAuthGroupRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userPutModuleAuthGroupRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateGetUserACLRoute 新建UserACLGetRoute
func CreateGetUserACLRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userGetACLRoute{authorityHandler: authorityHandler}
	return &i
}

type userGetModuleAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
}

type userGetModuleAuthGroupResult struct {
	common.Result
	model.UserModuleAuthGroupInfo
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
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		val := i.authorityHandler.QueryUserModuleAuthGroup(id)
		result.User = id
		result.ModuleAuthGroups = val.ModuleAuthGroups

		result.ErrCode = common.Success
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
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		var moduleAuthGroups []model.ModuleAuthGroup
		content := r.FormValue("moduleAuthGroup")
		err = json.Unmarshal([]byte(content), &moduleAuthGroups)
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateUserModuleAuthGroup(id, moduleAuthGroups)
		if ok {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
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
}

type userGetACLResult struct {
	common.Result
	User int
	ACLs []model.ACL
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
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		result.User = id
		result.ACLs = i.authorityHandler.QueryUserACL(id)
		result.ErrCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}