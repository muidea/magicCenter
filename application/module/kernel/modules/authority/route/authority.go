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
	"muidea.com/magicCenter/foundation/util"
)

// AppendAuthorityRoute append authority route
func AppendAuthorityRoute(routes []common.Route, authorityHandler common.AuthorityHandler) []common.Route {

	rt := CreateModuleACLGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLPostRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLDeleteRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLPutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLAuthGroupGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateACLAuthGroupPutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserModuleGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateUserModulePutRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateModuleUserGetRoute(authorityHandler)
	routes = append(routes, rt)

	rt = CreateModuleUserPutRoute(authorityHandler)
	routes = append(routes, rt)

	return routes
}

// CreateModuleACLGetRoute 新建ModuleACLGetRoute
func CreateModuleACLGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := moduleACLGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLGetRoute 新建ACLGetRoute
func CreateACLGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLPostRoute 新建ACLPostRoute
func CreateACLPostRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclPostRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLDeleteRoute 新建ACLDeleteRoute
func CreateACLDeleteRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclDeleteRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLPutRoute 新建ACLPutRoute
func CreateACLPutRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclPutRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLAuthGroupGetRoute 新建ACLAuthGroupGetRoute
func CreateACLAuthGroupGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclAuthGroupGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLAuthGroupPutRoute 新建ACLAuthGroupPutRoute
func CreateACLAuthGroupPutRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclAuthGroupPutRoute{authorityHandler: authorityHandler}
	return &i
}

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

type aclGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclGetResult struct {
	common.Result
	ACL model.ACL
}

func (i *aclGetRoute) Method() string {
	return common.GET
}

func (i *aclGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACL)
}

func (i *aclGetRoute) Handler() interface{} {
	return i.getACLHandler
}

func (i *aclGetRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclGetRoute) getACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLHandler")
	result := aclGetResult{}

	for true {
		url := r.URL.Query().Get("url")
		method := r.URL.Query().Get("method")

		acl, ok := i.authorityHandler.QueryACL(url, method)
		if ok {
			result.ACL = acl
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclPostRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclPostResult struct {
	common.Result
	ACL model.ACL
}

func (i *aclPostRoute) Method() string {
	return common.POST
}

func (i *aclPostRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostACL)
}

func (i *aclPostRoute) Handler() interface{} {
	return i.postACLHandler
}

func (i *aclPostRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclPostRoute) postACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("postACLHandler")

	result := aclPostResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		url := r.FormValue("acl-url")
		method := r.FormValue("acl-method")
		module := r.FormValue("acl-module")
		authGroups, ok := util.Str2IntArray(r.FormValue("acl-authgroup"))
		if !ok {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.InsertACL(url, method, module, 0, authGroups)
		if ok {
			result.ACL = acl
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
			result.Reason = "新建ACL失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclDeleteRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclDeleteResult struct {
	common.Result
}

func (i *aclDeleteRoute) Method() string {
	return common.POST
}

func (i *aclDeleteRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteACL)
}

func (i *aclDeleteRoute) Handler() interface{} {
	return i.deleteACLHandler
}

func (i *aclDeleteRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclDeleteRoute) deleteACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteACLHandler")

	result := aclDeleteResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.DeleteACL(id)
		if ok {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
			result.Reason = "删除ACL失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclPutRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclPutResult struct {
	common.Result
}

func (i *aclPutRoute) Method() string {
	return common.POST
}

func (i *aclPutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACL)
}

func (i *aclPutRoute) Handler() interface{} {
	return i.putACLHandler
}

func (i *aclPutRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclPutRoute) putACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putACLHandler")

	result := aclPutResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		enableList, _ := util.Str2IntArray(r.FormValue("acl-enablelist"))
		disableList, _ := util.Str2IntArray(r.FormValue("acl-disablelist"))

		ok := i.authorityHandler.UpdateACLStatus(enableList, disableList)
		if ok {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
			result.Reason = "更新ACL状态失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclAuthGroupGetRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclAuthGroupGetResult struct {
	common.Result
	ACL        int
	AuthGroups []model.AuthGroup
}

func (i *aclAuthGroupGetRoute) Method() string {
	return common.GET
}

func (i *aclAuthGroupGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACLAuthGroup)
}

func (i *aclAuthGroupGetRoute) Handler() interface{} {
	return i.getACLAuthGroupHandler
}

func (i *aclAuthGroupGetRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclAuthGroupGetRoute) getACLAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLAuthGroupHandler")

	result := aclAuthGroupGetResult{}
	for true {
		id, err := strconv.Atoi(r.URL.Query().Get("acl"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		result.ACL = id
		authGroups := i.authorityHandler.QueryACLAuthGroup(id)
		for _, v := range authGroups {
			switch v {
			case common.VisitorAuthGroup.ID:
				result.AuthGroups = append(result.AuthGroups, common.VisitorAuthGroup)
			case common.UserAuthGroup.ID:
				result.AuthGroups = append(result.AuthGroups, common.UserAuthGroup)
			case common.MaintainerAuthGroup.ID:
				result.AuthGroups = append(result.AuthGroups, common.MaintainerAuthGroup)
			}
		}

		result.ErrCode = common.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclAuthGroupPutRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclAuthGroupPutResult struct {
	common.Result
}

func (i *aclAuthGroupPutRoute) Method() string {
	return common.PUT
}

func (i *aclAuthGroupPutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACLAuthGroup)
}

func (i *aclAuthGroupPutRoute) Handler() interface{} {
	return i.putACLAuthGroupHandler
}

func (i *aclAuthGroupPutRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclAuthGroupPutRoute) putACLAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putACLAuthGroupHandler")

	result := aclAuthGroupPutResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}
		id, err := strconv.Atoi(r.FormValue("acl-id"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		authGroups, ok := util.Str2IntArray(r.FormValue("acl-authgroup"))
		if !ok {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok = i.authorityHandler.UpdateACLAuthGroup(id, authGroups)
		if ok {
			result.ErrCode = common.Success
		} else {
			result.ErrCode = common.Failed
			result.Reason = "更新ACL授权组失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
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
