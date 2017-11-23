package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	"muidea.com/magicCenter/foundation/net"
)

// CreateUserModuleGetRoute 新建UserModuleGetRoute
func CreateUserModuleGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userModuleGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateUserModulePutRoute 新建UserModulePutRoute
func CreateUserModulePutRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userModulePutRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateUserAuthGroupGetRoute 新建UserAuthGroupGetRoute
func CreateUserAuthGroupGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userAuthGroupGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateUserAuthGroupPutRoute 新建UserAuthGroupPutRoute
func CreateUserAuthGroupPutRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userAuthGroupPutRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateUserACLGetRoute 新建UserACLGetRoute
func CreateUserACLGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := userACLGetRoute{authorityHandler: authorityHandler}
	return &i
}

type userModuleGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type userModuleGetResult struct {
	common.Result
	User    int
	Modules []string
}

func (i *userModuleGetRoute) Method() string {
	return common.GET
}

func (i *userModuleGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetUserModule)
}

func (i *userModuleGetRoute) Handler() interface{} {
	return i.getUserModuleHandler
}

func (i *userModuleGetRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *userModuleGetRoute) getUserModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserModuleHandler")

	result := userModuleGetResult{}
	for true {
		id, err := strconv.Atoi(r.URL.Query().Get("user"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		result.Modules = i.authorityHandler.QueryUserModule(id)

		result.ErrCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userModulePutRoute struct {
	authorityHandler common.AuthorityHandler
}

type userModulePutResult struct {
	common.Result
}

func (i *userModulePutRoute) Method() string {
	return common.PUT
}

func (i *userModulePutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutUserModule)
}

func (i *userModulePutRoute) Handler() interface{} {
	return i.putUserModuleHandler
}

func (i *userModulePutRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *userModulePutRoute) putUserModuleHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putUserModuleHandler")

	result := userModulePutResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		id, err := strconv.Atoi(r.FormValue("user-id"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}
		modules := strings.Split(r.FormValue("user-module"), ",")

		ok := i.authorityHandler.UpdateUserModule(id, modules)
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

type userAuthGroupGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type userAuthGroupGetResult struct {
	common.Result
	User      int
	AuthGroup model.AuthGroup
}

func (i *userAuthGroupGetRoute) Method() string {
	return common.GET
}

func (i *userAuthGroupGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetUserAuthGroup)
}

func (i *userAuthGroupGetRoute) Handler() interface{} {
	return i.getUserAuthGroupHandler
}

func (i *userAuthGroupGetRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *userAuthGroupGetRoute) getUserAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserAuthGroupHandler")

	result := userAuthGroupGetResult{}
	for true {
		id, err := strconv.Atoi(r.URL.Query().Get("user"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		result.User = id
		result.AuthGroup = i.authorityHandler.QueryUserAuthGroup(id)
		result.ErrCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userAuthGroupPutRoute struct {
	authorityHandler common.AuthorityHandler
}

type userAuthGroupPutResult struct {
	common.Result
}

func (i *userAuthGroupPutRoute) Method() string {
	return common.PUT
}

func (i *userAuthGroupPutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutUserAuthGroup)
}

func (i *userAuthGroupPutRoute) Handler() interface{} {
	return i.putUserAuthGroupHandler
}

func (i *userAuthGroupPutRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *userAuthGroupPutRoute) putUserAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putUserAuthGroupHandler")

	result := userModulePutResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		id, err := strconv.Atoi(r.FormValue("user-id"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}
		authGroup, err := strconv.Atoi(r.FormValue("user-authgroup"))
		if err != nil || (authGroup != common.VisitorAuthGroup.ID && authGroup != common.UserAuthGroup.ID && authGroup != common.MaintainerAuthGroup.ID) {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateUserAuthGroup(id, authGroup)
		if ok {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
			result.Reason = "更新用户授权组失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type userACLGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type userACLGetResult struct {
	common.Result
	User int
	ACLs []model.ACL
}

func (i *userACLGetRoute) Method() string {
	return common.GET
}

func (i *userACLGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetUserACL)
}

func (i *userACLGetRoute) Handler() interface{} {
	return i.getUserACLHandler
}

func (i *userACLGetRoute) AuthGroup() int {
	return common.UserAuthGroup.ID
}

func (i *userACLGetRoute) getUserACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getUserACLHandler")

	result := userACLGetResult{}
	for true {
		id, err := strconv.Atoi(r.URL.Query().Get("user"))
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
