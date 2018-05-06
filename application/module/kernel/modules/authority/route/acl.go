package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	"muidea.com/magicCommon/foundation/net"
	common_const "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/model"
)

// CreateQueryACLRoute GetAclByModule
func CreateQueryACLRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclGetRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreateGetACLByIDRoute GetAclByID
func CreateGetACLByIDRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclGetByIDRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreatePostACLRoute CreateAcl
func CreatePostACLRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclPostRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreateDeleteACLRoute DeleteAcl
func CreateDeleteACLRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclDeleteRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreatePutACLRoute UpdateAcl
func CreatePutACLRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclPutRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreatePutACLsRoute UpdateAcls
func CreatePutACLsRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclPutsRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreateGetACLAuthGroupRoute GetAclAuthGroup
func CreateGetACLAuthGroupRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclGetAuthGroupRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreatePutACLAuthGroupRoute UpdateAclAuthGroup
func CreatePutACLAuthGroupRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclPutAuthGroupRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

type aclGetRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

type aclGetResult struct {
	common_result.Result
	ACL []model.ACLView `json:"acl"`
}

func (i *aclGetRoute) Method() string {
	return common.GET
}

func (i *aclGetRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QueryACL)
}

func (i *aclGetRoute) Handler() interface{} {
	return i.getHandler
}

func (i *aclGetRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclGetRoute) getHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getHandler")
	result := aclGetResult{}

	for true {
		module := r.URL.Query().Get("module")
		if module != "" {
			acls := i.authorityHandler.QueryACLByModule(module)
			for _, val := range acls {
				acl := model.ACLView{}
				acl.ACL = val

				result.ACL = append(result.ACL, acl)
			}
		} else {
			acls := i.authorityHandler.QueryAllACL()
			for _, val := range acls {
				acl := model.ACLView{}
				acl.ACL = val
				acl.Status = common_const.GetStatus(val.Status)

				result.ACL = append(result.ACL, acl)
			}
		}
		result.ErrorCode = common_result.Success

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclGetByIDRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

type aclGetByIDResult struct {
	common_result.Result
	ACLDetail model.ACLDetailView `json:"acl"`
}

func (i *aclGetByIDRoute) Method() string {
	return common.GET
}

func (i *aclGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACLByID)
}

func (i *aclGetByIDRoute) Handler() interface{} {
	return i.getByIDHandler
}

func (i *aclGetByIDRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclGetByIDRoute) getByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getByIDHandler")
	result := aclGetByIDResult{}

	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if ok {
			mod, _ := i.moduleHub.FindModule(acl.Module)
			result.ACLDetail.ACLDetail = acl
			result.ACLDetail.Status = common_const.GetStatus(acl.Status)
			result.ACLDetail.AuthGroup = common_const.GetAuthGroup(acl.AuthGroup)

			result.ACLDetail.Module.ID = mod.ID()
			result.ACLDetail.Module.Name = mod.Name()

			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
	moduleHub        common.ModuleHub
}

type aclPostParam struct {
	URL       string `json:"url"`
	Method    string `json:"method"`
	Module    string `json:"module"`
	AuthGroup int    `json:"authGroup"`
}

type aclPostResult struct {
	common_result.Result
	ACLDetail model.ACLDetailView `json:"acl"`
}

func (i *aclPostRoute) Method() string {
	return common.POST
}

func (i *aclPostRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostACL)
}

func (i *aclPostRoute) Handler() interface{} {
	return i.postHandler
}

func (i *aclPostRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclPostRoute) postHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("postHandler")

	result := aclPostResult{}
	for true {
		param := &aclPostParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.InsertACL(param.URL, param.Method, param.Module, 0, param.AuthGroup)
		if ok {
			result.ACLDetail.ACLDetail = acl
			result.ACLDetail.Status = common_const.GetStatus(acl.Status)
			result.ACLDetail.AuthGroup = common_const.GetAuthGroup(acl.AuthGroup)
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
	moduleHub        common.ModuleHub
}

type aclDeleteResult struct {
	common_result.Result
}

func (i *aclDeleteRoute) Method() string {
	return common.DELETE
}

func (i *aclDeleteRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteACL)
}

func (i *aclDeleteRoute) Handler() interface{} {
	return i.deleteHandler
}

func (i *aclDeleteRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclDeleteRoute) deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteHandler")

	result := aclDeleteResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.DeleteACL(id)
		if ok {
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
	moduleHub        common.ModuleHub
}

type aclPutParam struct {
	URL       string `json:"url"`
	Method    string `json:"method"`
	Module    string `json:"module"`
	AuthGroup int    `json:"authGroup"`
	Status    int    `json:"status"`
}

type aclPutResult struct {
	common_result.Result
}

func (i *aclPutRoute) Method() string {
	return common.PUT
}

func (i *aclPutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACL)
}

func (i *aclPutRoute) Handler() interface{} {
	return i.putHandler
}

func (i *aclPutRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclPutRoute) putHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putHandler")

	result := aclPutResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "对象不存在"
			break
		}

		param := &aclPutParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}
		acl.AuthGroup = param.AuthGroup
		acl.Status = param.Status

		ok = i.authorityHandler.UpdateACL(acl)
		if ok {
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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

type aclPutsRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

type aclPutsParam struct {
	EnableList  []int `json:"enablelist"`
	DisableList []int `json:"disablelist"`
}

type aclPutsResult struct {
	common_result.Result
}

func (i *aclPutsRoute) Method() string {
	return common.PUT
}

func (i *aclPutsRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACLs)
}

func (i *aclPutsRoute) Handler() interface{} {
	return i.putsHandler
}

func (i *aclPutsRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclPutsRoute) putsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putsHandler")

	result := aclPutsResult{}
	for true {
		param := &aclPutsParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateACLStatus(param.EnableList, param.DisableList)
		if ok {
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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

type aclGetAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

type aclGetAuthGroupResult struct {
	common_result.Result
	ACL       model.ACL       `json:"acl"`
	AuthGroup model.AuthGroup `json:"authGroup"`
}

func (i *aclGetAuthGroupRoute) Method() string {
	return common.GET
}

func (i *aclGetAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACLAuthGroup)
}

func (i *aclGetAuthGroupRoute) Handler() interface{} {
	return i.getAuthGroupHandler
}

func (i *aclGetAuthGroupRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclGetAuthGroupRoute) getAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getAuthGroupHandler")

	result := aclGetAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		authGroup, ok := i.authorityHandler.QueryACLAuthGroup(id)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "无效参数"
			break
		}

		result.ACL = acl.ACL
		result.AuthGroup = authGroup
		result.ErrorCode = common_result.Success

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclPutAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

type aclPutAuthGroupParam struct {
	AuthGroup int `json:"authGroup"`
}

type aclPutACLAuthGroupResult struct {
	common_result.Result
}

func (i *aclPutAuthGroupRoute) Method() string {
	return common.PUT
}

func (i *aclPutAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACLAuthGroup)
}

func (i *aclPutAuthGroupRoute) Handler() interface{} {
	return i.putAuthGroupHandler
}

func (i *aclPutAuthGroupRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclPutAuthGroupRoute) putAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putAuthGroupHandler")

	result := aclPutACLAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		param := &aclPutAuthGroupParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.UpdateACLAuthGroup(id, param.AuthGroup)
		if ok {
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
