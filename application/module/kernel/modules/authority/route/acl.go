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

// CreateACLGetRoute 新建ACLGetRoute
func CreateACLGetRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclGetRoute{authorityHandler: authorityHandler}
	return &i
}

// CreateACLGetByModuleRoute 新建ACLGetByModuleRoute
func CreateACLGetByModuleRoute(authorityHandler common.AuthorityHandler) common.Route {
	i := aclGetByModuleRoute{authorityHandler: authorityHandler}
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
	return net.JoinURL(def.URL, def.GetACL)
}

func (i *aclGetByModuleRoute) Handler() interface{} {
	return i.getACLHandler
}

func (i *aclGetByModuleRoute) AuthGroup() int {
	return common.MaintainerAuthGroup.ID
}

func (i *aclGetByModuleRoute) getACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getACLHandler")
	result := aclGetByModuleResult{}

	for true {
		module := r.URL.Query().Get("module")

		result.Module = module
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
		authGroup, err := strconv.Atoi(r.FormValue("acl-authgroup"))
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
	ACL       int
	AuthGroup model.AuthGroup
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
		authGroup := i.authorityHandler.QueryACLAuthGroup(id)
		switch authGroup {
		case common.VisitorAuthGroup.ID:
			result.AuthGroup = common.VisitorAuthGroup
		case common.UserAuthGroup.ID:
			result.AuthGroup = common.UserAuthGroup
		case common.MaintainerAuthGroup.ID:
			result.AuthGroup = common.MaintainerAuthGroup
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

		authGroup, err := strconv.Atoi(r.FormValue("acl-authgroup"))
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
