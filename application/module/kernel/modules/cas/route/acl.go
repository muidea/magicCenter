package route

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

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

	rt = CreateEnableACLRoute(casHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateDisableACLRoute(casHandler, sessionRegistry)
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

// CreateEnableACLRoute 新建AddACL 路由
func CreateEnableACLRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLEnableRoute{
		casHandler: casHandler}
	return &i
}

// CreateDisableACLRoute 新建DelACL 路由
func CreateDisableACLRoute(casHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityACLDisableRoute{
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

		acls, ok := i.casHandler.QueryACL(modules[0], val)
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

type authorityACLEnableRoute struct {
	casHandler common.CASHandler
}

type authorityACLEnableResult struct {
	common.Result
}

func (i *authorityACLEnableRoute) Method() string {
	return common.POST
}

func (i *authorityACLEnableRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/enable/")
}

func (i *authorityACLEnableRoute) Handler() interface{} {
	return i.enableACLHandler
}

func (i *authorityACLEnableRoute) enableACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("enableACLHandler")

	result := authorityACLEnableResult{}
	for true {
		r.ParseForm()
		acls, ok := util.Str2IntArray(r.FormValue("acl-list"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}
		ok = i.casHandler.EnableACL(acls)
		if !ok {
			result.ErrCode = 1
			result.Reason = "新增失败"
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

type authorityACLDisableRoute struct {
	casHandler common.CASHandler
}

type authorityACLDisableResult struct {
	common.Result
}

func (i *authorityACLDisableRoute) Method() string {
	return common.POST
}

func (i *authorityACLDisableRoute) Pattern() string {
	return net.JoinURL(def.URL, "/acl/disable/")
}

func (i *authorityACLDisableRoute) Handler() interface{} {
	return i.disableACLHandler
}

func (i *authorityACLDisableRoute) disableACLHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("disableACLHandler")

	result := authorityACLDisableResult{}
	for true {
		r.ParseForm()
		acls, ok := util.Str2IntArray(r.FormValue("acl-list"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "参数非法"
			break
		}
		if !i.casHandler.DisableACL(acls) {
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
