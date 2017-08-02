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

// AppendACLRoute 追加acl 路由
func AppendACLRoute(routes []common.Route, authorityHandler common.AuthorityHandler, sessionRegistry common.SessionRegistry) []common.Route {
	// 查询全部ACL或指定Module的ACL
	rt := CreateQueryACLRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)
	// Enable或Disable ACL
	rt = CreateStatusACLRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)
	// 调整ACL对应的AuthGroup
	rt = CreateUpdateACLRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateQueryACLRoute 新建QueryACL 路由
func CreateQueryACLRoute(authorityHandler common.AuthorityHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLQueryRoute{
		authorityHandler: authorityHandler}
	return &i
}

// CreateStatusACLRoute 新建StatusACL 路由
func CreateStatusACLRoute(authorityHandler common.AuthorityHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLStatusRoute{
		authorityHandler: authorityHandler}
	return &i
}

// CreateUpdateACLRoute 新建UpdateACL 路由
func CreateUpdateACLRoute(authorityHandler common.AuthorityHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLAuthGroupRoute{
		authorityHandler: authorityHandler}
	return &i
}

type authorityACLQueryRoute struct {
	authorityHandler common.AuthorityHandler
}

type authorityACLQueryResult struct {
	common.Result
	ACLs []model.ACL
}

func (i *authorityACLQueryRoute) Method() string {
	return common.GET
}

func (i *authorityACLQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/")
}

func (i *authorityACLQueryRoute) Handler() interface{} {
	return i.queryACLHandler
}

func (i *authorityACLQueryRoute) queryACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryACLHandler")

	result := authorityACLQueryResult{}
	for true {
		modules := r.URL.Query()["module"]
		if len(modules) < 1 {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}
		status := r.URL.Query()["status"]
		if len(status) < 1 {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}
		val, err := strconv.Atoi(status[0])
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		acls, ok := i.authorityHandler.QueryACL(modules[0], val)
		if !ok {
			result.ErrCode = 1
			result.Reason = "查询失败"
			break
		}

		result.ErrCode = 0
		result.ACLs = acls
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityACLStatusRoute struct {
	authorityHandler common.AuthorityHandler
}

type authorityACLStatusResult struct {
	common.Result
}

func (i *authorityACLStatusRoute) Method() string {
	return common.POST
}

func (i *authorityACLStatusRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/status/")
}

func (i *authorityACLStatusRoute) Handler() interface{} {
	return i.statusACLHandler
}

func (i *authorityACLStatusRoute) statusACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("statusACLHandler")

	result := authorityACLStatusResult{}
	for true {
		r.ParseForm()
		enableAcls, ok := util.Str2IntArray(r.FormValue("enable-list"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}
		disableAcls, ok := util.Str2IntArray(r.FormValue("disable-list"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}
		ok = i.authorityHandler.EnableACL(enableAcls)
		if !ok {
			result.ErrCode = 1
			result.Reason = "启用失败"
			break
		}
		ok = i.authorityHandler.DisableACL(disableAcls)
		if !ok {
			result.ErrCode = 1
			result.Reason = "禁用失败"
			break
		}

		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityACLAuthGroupRoute struct {
	authorityHandler common.AuthorityHandler
}

type authorityACLAuthGroupResult struct {
	common.Result
}

func (i *authorityACLAuthGroupRoute) Method() string {
	return common.POST
}

func (i *authorityACLAuthGroupRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/authgroup/")
}

func (i *authorityACLAuthGroupRoute) Handler() interface{} {
	return i.updateACLAuthGroupHandler
}

func (i *authorityACLAuthGroupRoute) updateACLAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateACLAuthGroupHandler")

	result := authorityACLAuthGroupResult{}
	for true {
		r.ParseForm()

		aclID, err := strconv.Atoi(r.FormValue("acl-id"))
		if err != nil {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}
		authGroup, ok := util.Str2IntArray(r.FormValue("acl-authgroup"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}

		_, ok = i.authorityHandler.AdjustACLAuthGroup(aclID, authGroup)
		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}

		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
