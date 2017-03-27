package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/cas/def"
	"muidea.com/magicCenter/foundation/net"
	"muidea.com/magicCenter/foundation/util"
)

// AppendACLRoute 追加acl 路由
func AppendACLRoute(routes []common.Route, casHandler common.CASHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateQueryACLRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateAddACLRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDelACLRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateUpdateACLRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateQueryACLRoute 新建QueryACL 路由
func CreateQueryACLRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLQueryRoute{
		casHandler: casHandler}
	return &i
}

// CreateAddACLRoute 新建AddACL 路由
func CreateAddACLRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLAddRoute{
		casHandler: casHandler}
	return &i
}

// CreateDelACLRoute 新建DelACL 路由
func CreateDelACLRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLDelRoute{
		casHandler: casHandler}
	return &i
}

// CreateUpdateACLRoute 新建UpdateACL 路由
func CreateUpdateACLRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLUpdateRoute{
		casHandler: casHandler}
	return &i
}

type authorityACLQueryRoute struct {
	casHandler common.CASHandler
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

		acls, ok := i.casHandler.QueryACL(modules[0])
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

type authorityACLAddRoute struct {
	casHandler common.CASHandler
}

type authorityACLAddResult struct {
	common.Result
	ACL model.ACL
}

func (i *authorityACLAddRoute) Method() string {
	return common.POST
}

func (i *authorityACLAddRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/add/")
}

func (i *authorityACLAddRoute) Handler() interface{} {
	return i.addACLHandler
}

func (i *authorityACLAddRoute) addACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("addACLHandler")

	result := authorityACLAddResult{}
	for true {
		r.ParseForm()

		url := r.FormValue("acl-url")
		method := r.FormValue("acl-method")
		module := r.FormValue("acl-module")
		acl, ok := i.casHandler.AddACL(url, method, module)
		if !ok {
			result.ErrCode = 1
			result.Reason = "新增失败"
			break
		}

		result.ACL = acl
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityACLDelRoute struct {
	casHandler common.CASHandler
}

type authorityACLDelResult struct {
	common.Result
}

func (i *authorityACLDelRoute) Method() string {
	return common.POST
}

func (i *authorityACLDelRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/del/")
}

func (i *authorityACLDelRoute) Handler() interface{} {
	return i.delACLHandler
}

func (i *authorityACLDelRoute) delACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("delACLHandler")

	result := authorityACLDelResult{}
	for true {
		r.ParseForm()

		url := r.FormValue("acl-url")
		method := r.FormValue("acl-method")
		module := r.FormValue("acl-module")
		if !i.casHandler.DelACL(url, method, module) {
			result.ErrCode = 1
			result.Reason = "删除失败"
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

type authorityACLUpdateRoute struct {
	casHandler common.CASHandler
}

type authorityACLUpdateResult struct {
	common.Result
	ACL model.ACL
}

func (i *authorityACLUpdateRoute) Method() string {
	return common.POST
}

func (i *authorityACLUpdateRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/update/")
}

func (i *authorityACLUpdateRoute) Handler() interface{} {
	return i.updateACLHandler
}

func (i *authorityACLUpdateRoute) updateACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("updateACLHandler")

	result := authorityACLUpdateResult{}
	for true {
		r.ParseForm()

		url := r.FormValue("acl-url")
		method := r.FormValue("acl-method")
		module := r.FormValue("acl-module")
		authGroup, ok := util.Str2IntArray(r.FormValue("acl-authgroup"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}

		acl, ok := i.casHandler.AdjustACLAuthGroup(url, method, module, authGroup)
		if !ok {
			result.ErrCode = 1
			result.Reason = "更新失败"
			break
		}

		result.ACL = acl
		result.ErrCode = 0
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
