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
	"muidea.com/magicCenter/foundation/util"
)

// CreateGetACLByIDRoute GetAclByID
func CreateGetACLByIDRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclGetByIDRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateGetACLByModuleRoute GetAclByModule
func CreateGetACLByModuleRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclGetByModuleRoute{authorityHandler: authorityHandler}
	return &i
}

// CreatePostACLRoute CreateAcl
func CreatePostACLRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclPostRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateDeleteACLRoute DeleteAcl
func CreateDeleteACLRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclDeleteRoute{authorityHandler: authorityHandler}
	return &i
}

// CreatePutACLRoute UpdateAcl
func CreatePutACLRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclPutRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateGetACLAuthGroupRoute GetAclAuthGroup
func CreateGetACLAuthGroupRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclGetAclAuthGroupRoute{authorityHandler: authorityHandler}
	return &i
}

// CreatePutACLAuthGroupRoute UpdateAclAuthGroup
func CreatePutACLAuthGroupRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclPutACLAuthGroupRoute{authorityHandler: authorityHandler}
	return &i
}

type aclGetByIDRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclGetResult struct {
	common.Result
	ACL model.ACL
}

func (i *aclGetByIDRoute) Method() string {
	return common.GET
}

func (i *aclGetByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACLByID)
}

func (i *aclGetByIDRoute) Handler() interface{} {
	return i.getACLHandler
}

func (i *aclGetByIDRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclGetByIDRoute) getACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLByIDHandler")
	result := aclGetResult{}

	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
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

type aclGetByModuleRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclGetByModuleResult struct {
	common.Result
	Module string
	ACLs   []model.ACL
}

func (i *aclGetByModuleRoute) Method() string {
	return common.GET
}

func (i *aclGetByModuleRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QueryACLByModule)
}

func (i *aclGetByModuleRoute) Handler() interface{} {
	return i.getACLHandler
}

func (i *aclGetByModuleRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclGetByModuleRoute) getACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLByModuleHandler")
	result := aclGetByModuleResult{}

	for true {
		module := r.URL.Query().Get("module")

		result.Module = module
		result.ACLs = i.authorityHandler.QueryACLByModule(module)
		result.ErrCode = common.Success

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

		url := r.FormValue("url")
		method := r.FormValue("method")
		module := r.FormValue("module")
		authGroup, err := strconv.Atoi(r.FormValue("authgroup"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.InsertACL(url, method, module, 0, authGroup)
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

		enableList, _ := util.Str2IntArray(r.FormValue("enablelist"))
		disableList, _ := util.Str2IntArray(r.FormValue("disablelist"))

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

type aclGetAclAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclGetAclAuthGroupResult struct {
	common.Result
	ACL       int
	AuthGroup model.AuthGroup
}

func (i *aclGetAclAuthGroupRoute) Method() string {
	return common.GET
}

func (i *aclGetAclAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACLAuthGroup)
}

func (i *aclGetAclAuthGroupRoute) Handler() interface{} {
	return i.getACLAuthGroupHandler
}

func (i *aclGetAclAuthGroupRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclGetAclAuthGroupRoute) getACLAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLAuthGroupHandler")

	result := aclGetAclAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		result.ACL = id
		result.ErrCode = common.Success
		authGroup, ok := i.authorityHandler.QueryACLAuthGroup(id)
		if !ok {
			result.ErrCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		switch authGroup {
		case common.VisitorAuthGroup.ID:
			result.AuthGroup = common.VisitorAuthGroup
		case common.UserAuthGroup.ID:
			result.AuthGroup = common.UserAuthGroup
		case common.MaintainerAuthGroup.ID:
			result.AuthGroup = common.MaintainerAuthGroup
		default:
			result.ErrCode = common.Failed
			result.Reason = "非法授权组"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type aclPutACLAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclPutACLAuthGroupResult struct {
	common.Result
}

func (i *aclPutACLAuthGroupRoute) Method() string {
	return common.PUT
}

func (i *aclPutACLAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutACLAuthGroup)
}

func (i *aclPutACLAuthGroupRoute) Handler() interface{} {
	return i.putACLAuthGroupHandler
}

func (i *aclPutACLAuthGroupRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclPutACLAuthGroupRoute) putACLAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putACLAuthGroupHandler")

	result := aclPutACLAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		authGroup, err := strconv.Atoi(r.FormValue("authgroup"))
		if err != nil {
			result.ErrCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateACLAuthGroup(id, authGroup)
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
