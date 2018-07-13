package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/authority/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
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
	i := aclCreateRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreateDeleteACLRoute DeleteAcl
func CreateDeleteACLRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclDeleteRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
	return &i
}

// CreatePutACLRoute UpdateAcl
func CreatePutACLRoute(authorityHandler common.AuthorityHandler, moduleHub common.ModuleHub) common.Route {
	i := aclUpdateRoute{authorityHandler: authorityHandler, moduleHub: moduleHub}
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
	result := common_def.GetACLListResult{}

	for true {
		module := r.URL.Query().Get("module")
		if module != "" {
			acls := i.authorityHandler.QueryACLByModule(module)
			for _, val := range acls {
				mod, _ := i.moduleHub.FindModule(module)
				acl := model.ACLView{}
				acl.ACL = val
				acl.Status = common_const.GetStatus(val.Status)
				acl.AuthGroup = common_const.GetAuthGroup(val.AuthGroup)

				acl.Module.ID = mod.ID()
				acl.Module.Name = mod.Name()
				result.ACL = append(result.ACL, acl)
			}
		} else {
			acls := i.authorityHandler.QueryAllACL()
			for _, val := range acls {
				mod, _ := i.moduleHub.FindModule(val.Module)

				acl := model.ACLView{}
				acl.ACL = val
				acl.Status = common_const.GetStatus(val.Status)
				acl.AuthGroup = common_const.GetAuthGroup(val.AuthGroup)

				acl.Module.ID = mod.ID()
				acl.Module.Name = mod.Name()

				result.ACL = append(result.ACL, acl)
			}
		}
		result.ErrorCode = common_def.Success

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
	result := common_def.GetACLResult{}

	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if ok {
			mod, _ := i.moduleHub.FindModule(acl.Module)
			result.ACL.ACL = acl
			result.ACL.Status = common_const.GetStatus(acl.Status)
			result.ACL.AuthGroup = common_const.GetAuthGroup(acl.AuthGroup)

			result.ACL.Module.ID = mod.ID()
			result.ACL.Module.Name = mod.Name()

			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclCreateRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

func (i *aclCreateRoute) Method() string {
	return common.POST
}

func (i *aclCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostACL)
}

func (i *aclCreateRoute) Handler() interface{} {
	return i.postHandler
}

func (i *aclCreateRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclCreateRoute) postHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("postHandler")

	result := common_def.CreateACLResult{}
	for true {
		param := &common_def.CreateACLParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.InsertACL(param.URL, param.Method, param.Module, 0, param.AuthGroup)
		if ok {
			mod, _ := i.moduleHub.FindModule(acl.Module)
			result.ACL.ACL = acl
			result.ACL.Status = common_const.GetStatus(acl.Status)
			result.ACL.AuthGroup = common_const.GetAuthGroup(acl.AuthGroup)

			result.ACL.Module.ID = mod.ID()
			result.ACL.Module.Name = mod.Name()

			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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

	result := common_def.DestroyACLResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.DeleteACL(id)
		if ok {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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

type aclUpdateRoute struct {
	authorityHandler common.AuthorityHandler
	moduleHub        common.ModuleHub
}

func (i *aclUpdateRoute) Method() string {
	return common.PUT
}

func (i *aclUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACL)
}

func (i *aclUpdateRoute) Handler() interface{} {
	return i.putHandler
}

func (i *aclUpdateRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *aclUpdateRoute) putHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putHandler")

	result := common_def.UpdateACLResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法参数"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "对象不存在"
			break
		}

		param := &common_def.UpdateACLParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法参数"
			break
		}
		acl.AuthGroup = param.AuthGroup
		acl.Status = param.Status

		ok = i.authorityHandler.UpdateACL(acl)
		if ok {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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

	result := common_def.UpdateACLStatusResult{}
	for true {
		param := &common_def.UpdateACLStatusParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateACLStatus(param.EnableList, param.DisableList)
		if ok {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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

	result := common_def.GetAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		authGroup, ok := i.authorityHandler.QueryACLAuthGroup(id)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "无效参数"
			break
		}

		result.AuthGroup = authGroup
		result.ErrorCode = common_def.Success

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

	result := common_def.UpdateAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		param := &common_def.UpdateAuthGroupParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.UpdateACLAuthGroup(id, param.AuthGroup)
		if ok {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
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
