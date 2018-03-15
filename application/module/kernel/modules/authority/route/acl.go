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
	i := aclGetACLAuthGroupRoute{authorityHandler: authorityHandler}
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
	ACLDetail model.ACLDetailView `json:"acl"`
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
			result.ErrorCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if ok {
			result.ACLDetail.ACLDetail = acl
			if acl.AuthGroup == common.UserAuthGroup.ID {
				result.ACLDetail.AuthGroup = common.UserAuthGroup.Unit
			} else if acl.AuthGroup == common.MaintainerAuthGroup.ID {
				result.ACLDetail.AuthGroup = common.MaintainerAuthGroup.Unit
			} else {
				result.ACLDetail.AuthGroup = common.VisitorAuthGroup.Unit
			}

			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
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
	ACL []model.ACLView `json:"acl"`
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

				result.ACL = append(result.ACL, acl)
			}
		}
		result.ErrorCode = common.Success

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

type aclPostParam struct {
	URL       string `json:"url"`
	Method    string `json:"method"`
	Module    string `json:"module"`
	AuthGroup int    `json:"authGroup"`
}

type aclPostResult struct {
	common.Result
	ACLDetail model.ACLDetailView `json:"acl"`
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
		param := &aclPostParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.InsertACL(param.URL, param.Method, param.Module, 0, param.AuthGroup)
		if ok {
			result.ACLDetail.ACLDetail = acl
			if acl.AuthGroup == common.UserAuthGroup.ID {
				result.ACLDetail.AuthGroup = common.UserAuthGroup.Unit
			} else if acl.AuthGroup == common.MaintainerAuthGroup.ID {
				result.ACLDetail.AuthGroup = common.MaintainerAuthGroup.Unit
			} else {
				result.ACLDetail.AuthGroup = common.VisitorAuthGroup.Unit
			}
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
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
			result.ErrorCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.DeleteACL(id)
		if ok {
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
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

type aclPutParam struct {
	EnableList  []int `json:"enablelist"`
	DisableList []int `json:"disablelist"`
}

type aclPutResult struct {
	common.Result
}

func (i *aclPutRoute) Method() string {
	return common.PUT
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
		param := &aclPutParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "非法参数"
			break
		}

		ok := i.authorityHandler.UpdateACLStatus(param.EnableList, param.DisableList)
		if ok {
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
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

type aclGetACLAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
}

type aclGetACLAuthGroupResult struct {
	common.Result
	ACL       model.ACL       `json:"acl"`
	AuthGroup model.AuthGroup `json:"authGroup"`
}

func (i *aclGetACLAuthGroupRoute) Method() string {
	return common.GET
}

func (i *aclGetACLAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetACLAuthGroup)
}

func (i *aclGetACLAuthGroupRoute) Handler() interface{} {
	return i.getACLAuthGroupHandler
}

func (i *aclGetACLAuthGroupRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclGetACLAuthGroupRoute) getACLAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLAuthGroupHandler")

	result := aclGetACLAuthGroupResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		id, err := strconv.Atoi(strID)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.authorityHandler.QueryACLByID(id)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		authGroup, ok := i.authorityHandler.QueryACLAuthGroup(id)
		if !ok {
			result.ErrorCode = common.Failed
			result.Reason = "无效参数"
			break
		}

		result.ACL = acl.ACL
		result.AuthGroup = authGroup
		result.ErrorCode = common.Success

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

type aclPutACLAuthGroupParam struct {
	AuthGroup int `json:"authGroup"`
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
			result.ErrorCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		param := &aclPutACLAuthGroupParam{}
		err = net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common.Failed
			result.Reason = "参数非法"
			break
		}

		ok := i.authorityHandler.UpdateACLAuthGroup(id, param.AuthGroup)
		if ok {
			result.ErrorCode = common.Success
		} else {
			result.ErrorCode = common.Failed
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
