package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	"muidea.com/magicCenter/foundation/net"
	"muidea.com/magicCenter/foundation/util"
)

// CreateModuleACLGetRoute 新建ModuleACLGetRoute
func CreateModuleACLGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := moduleACLGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateModuleUserGetRoute 新建ModuleUserGetRoute
func CreateModuleUserGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := moduleUserGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateModuleUserPutRoute 新建ModuleUserPutRoute
func CreateModuleUserPutRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := moduleUserPutRoute{authorityHandler: authorityHandler}
	return &i
}

type moduleACLGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type moduleACLGetResult struct {
	common.Result
	module string
	ACLs   []model.ACL
}

func (i *moduleACLGetRoute) Method() string {
	return common.GET
}

func (i *moduleACLGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetModuleACL)
}

func (i *moduleACLGetRoute) Handler() interface{} {
	return i.getModuleACLHandler
}

func (i *moduleACLGetRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *moduleACLGetRoute) getModuleACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getModuleACLHandler")
	result := moduleACLGetResult{}

	for true {
		module := r.URL.Query().Get("module")
		result.module = module

		result.ACLs = i.authorityHandler.QueryModuleACL(module)
		result.ErrCode = common.Success

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type moduleUserGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type moduleUserGetResult struct {
	common.Result
	Module string
	Users  []int
}

func (i *moduleUserGetRoute) Method() string {
	return common.GET
}

func (i *moduleUserGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetModuleUser)
}

func (i *moduleUserGetRoute) Handler() interface{} {
	return i.getModuleUserHandler
}

func (i *moduleUserGetRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *moduleUserGetRoute) getModuleUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getModuleUserHandler")

	result := moduleUserGetResult{}
	for true {
		module := r.URL.Query().Get("module")
		result.Module = module
		result.Users = i.authorityHandler.QueryModuleUser(module)
		result.ErrCode = common.Success

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type moduleUserPutRoute struct {
	authorityHandler common.AuthorityHandler
}

type moduleUserPutResult struct {
	common.Result
}

func (i *moduleUserPutRoute) Method() string {
	return common.PUT
}

func (i *moduleUserPutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutModuleUser)
}

func (i *moduleUserPutRoute) Handler() interface{} {
	return i.putModuleUserHandler
}

func (i *moduleUserPutRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *moduleUserPutRoute) putModuleUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putModuleUserHandler")

	result := moduleUserPutResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		id := r.FormValue("module-id")
		users, ok := util.Str2IntArray(r.FormValue("module-user"))

		ok = i.authorityHandler.UpdateModuleUser(id, users)
		if ok {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
			result.Reason = "更新模块用户信息失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
