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

// AppendAuthGropRoute 追加authgroup 路由
func AppendAuthGropRoute(routes []common.Route, authorityHandler common.CASHandler, sessionRegistry common.SessionRegistry) []common.Route {
	rt := CreateQueryAuthGroupRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateQueryUserAuthGroupRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)

	rt = CreateAdjustUserAuthGroupRoute(authorityHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateQueryAuthGroupRoute 新建QueryAuthGroup 路由
func CreateQueryAuthGroupRoute(authorityHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityAuthGroupQueryRoute{
		authorityHandler: authorityHandler}
	return &i
}

// CreateQueryUserAuthGroupRoute 新建QueryAuthGroup 路由
func CreateQueryUserAuthGroupRoute(authorityHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityUserAuthGroupQueryRoute{
		authorityHandler: authorityHandler}
	return &i
}

// CreateAdjustUserAuthGroupRoute 新建AdjustAuthGroup 路由
func CreateAdjustUserAuthGroupRoute(authorityHandler common.CASHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := authorityUserAuthGroupAdjustRoute{
		authorityHandler: authorityHandler}
	return &i
}

type authorityAuthGroupQueryRoute struct {
	authorityHandler common.CASHandler
}

type authorityAuthGroupQueryResult struct {
	common.Result
	AuthGroup []model.AuthGroup
}

func (i *authorityAuthGroupQueryRoute) Method() string {
	return common.GET
}

func (i *authorityAuthGroupQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, "/authgroup/")
}

func (i *authorityAuthGroupQueryRoute) Handler() interface{} {
	return i.queryAuthGroupHandler
}

func (i *authorityAuthGroupQueryRoute) queryAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryAuthGroupHandler")

	result := authorityAuthGroupQueryResult{}
	for true {
		modules := r.URL.Query()["module"]
		if len(modules) < 1 {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		authGroups, ok := i.authorityHandler.QueryAuthGroup(modules[0])
		if !ok {
			result.ErrCode = 1
			result.Reason = "查询失败"
			break
		}

		result.ErrCode = 0
		result.AuthGroup = authGroups
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityUserAuthGroupQueryRoute struct {
	authorityHandler common.CASHandler
}

type authorityUserAuthGroupQueryResult struct {
	common.Result
	AuthGroup []int
}

func (i *authorityUserAuthGroupQueryRoute) Method() string {
	return common.GET
}

func (i *authorityUserAuthGroupQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, "/authgroup/[0-9]+/")
}

func (i *authorityUserAuthGroupQueryRoute) Handler() interface{} {
	return i.queryUserAuthGroupHandler
}

func (i *authorityUserAuthGroupQueryRoute) queryUserAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("queryUserAuthGroupHandler")

	result := authorityUserAuthGroupQueryResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		authGroups, ok := i.authorityHandler.GetUserAuthGroup(id)
		if !ok {
			result.ErrCode = 1
			result.Reason = "查询失败"
			break
		}

		result.ErrCode = 0
		result.AuthGroup = authGroups
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type authorityUserAuthGroupAdjustRoute struct {
	authorityHandler common.CASHandler
}

type authorityUserAuthGroupAdjustResult struct {
	common.Result
	AuthGroup []int
}

func (i *authorityUserAuthGroupAdjustRoute) Method() string {
	return common.POST
}

func (i *authorityUserAuthGroupAdjustRoute) Pattern() string {
	return net.JoinURL(def.URL, "/authgroup/[0-9]+/")
}

func (i *authorityUserAuthGroupAdjustRoute) Handler() interface{} {
	return i.adjustUserAuthGroupHandler
}

func (i *authorityUserAuthGroupAdjustRoute) adjustUserAuthGroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("adjustUserAuthGroupHandler")

	result := authorityUserAuthGroupQueryResult{}
	_, value := net.SplitRESTAPI(r.URL.Path)
	for true {
		id, err := strconv.Atoi(value)
		if err != nil {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		r.ParseForm()
		authGroups, ok := util.Str2IntArray(r.FormValue("acl-authgroup"))
		if !ok {
			result.ErrCode = 1
			result.Reason = "无效参数"
			break
		}

		ok = i.authorityHandler.AdjustUserAuthGroup(id, authGroups)
		if !ok {
			result.ErrCode = 1
			result.Reason = "调整失败"
			break
		}

		result.ErrCode = 0
		result.AuthGroup = authGroups
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
