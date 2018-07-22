package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/common"
	"muidea.com/magicCenter/module/modules/endpoint/def"
	common_const "muidea.com/magicCommon/common"
	common_def "muidea.com/magicCommon/def"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// CreateQueryEndpointRoute QueryEndpoint
func CreateQueryEndpointRoute(endpointHandler common.EndpointHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointQueryRoute{endpointHandler: endpointHandler, accountHandler: accountHandler}
	return &i
}

// CreateQueryByIDEndpointRoute QueryByIDEndpoint
func CreateQueryByIDEndpointRoute(endpointHandler common.EndpointHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointQueryByIDRoute{endpointHandler: endpointHandler, accountHandler: accountHandler}
	return &i
}

// CreatePostEndpointRoute CreateEndpoint
func CreatePostEndpointRoute(endpointHandler common.EndpointHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointPostRoute{endpointHandler: endpointHandler, accountHandler: accountHandler}
	return &i
}

// CreateDeleteEndpointRoute DeleteEndpoint
func CreateDeleteEndpointRoute(endpointHandler common.EndpointHandler, accountHandler common.AccountHandler) common.Route {
	i := endpointDeleteRoute{endpointHandler: endpointHandler, accountHandler: accountHandler}
	return &i
}

// CreatePutEndpointRoute UpdateEndpoint
func CreatePutEndpointRoute(endpointHandler common.EndpointHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointPutRoute{endpointHandler: endpointHandler, accountHandler: accountHandler}
	return &i
}

// CreateGetEndpointAuthRoute VerifyEndpointAuth
func CreateGetEndpointAuthRoute(endpointHandler common.EndpointHandler, sessionRegistry common.SessionRegistry) common.Route {

	i := endpointAuthRoute{endpointHandler: endpointHandler, sessionRegistry: sessionRegistry}
	return &i
}

type endpointQueryRoute struct {
	endpointHandler common.EndpointHandler
	accountHandler  common.AccountHandler
}

func (i *endpointQueryRoute) Method() string {
	return common.GET
}

func (i *endpointQueryRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QueryEndpoint)
}

func (i *endpointQueryRoute) Handler() interface{} {
	return i.getHandler
}

func (i *endpointQueryRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *endpointQueryRoute) getHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getHandler")
	result := common_def.QueryEndpointListResult{}

	for {
		endpoints := i.endpointHandler.QueryAllEndpoint()
		for _, val := range endpoints {
			endpoint := model.EndpointView{}
			endpoint.Endpoint = val
			endpoint.User = i.accountHandler.GetUsers(val.User)

			result.Endpoint = append(result.Endpoint, endpoint)
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

type endpointQueryByIDRoute struct {
	endpointHandler common.EndpointHandler
	accountHandler  common.AccountHandler
}

func (i *endpointQueryByIDRoute) Method() string {
	return common.GET
}

func (i *endpointQueryByIDRoute) Pattern() string {
	return net.JoinURL(def.URL, def.QueryByIDEndpoint)
}

func (i *endpointQueryByIDRoute) Handler() interface{} {
	return i.getHandler
}

func (i *endpointQueryByIDRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *endpointQueryByIDRoute) getHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("getHandler")
	result := common_def.QueryEndpointResult{}

	for {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		endpoint, ok := i.endpointHandler.QueryEndpointByID(strID)
		if ok {
			result.Endpoint.Endpoint = endpoint
			result.Endpoint.User = i.accountHandler.GetUsers(endpoint.User)
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.NoExist
			result.Reason = "对象不存在"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type endpointPostRoute struct {
	endpointHandler common.EndpointHandler
	accountHandler  common.AccountHandler
}

func (i *endpointPostRoute) Method() string {
	return common.POST
}

func (i *endpointPostRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostEndpoint)
}

func (i *endpointPostRoute) Handler() interface{} {
	return i.postHandler
}

func (i *endpointPostRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *endpointPostRoute) postHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("postHandler")

	result := common_def.CreateEndpointResult{}
	for true {
		param := &common_def.CreateEndpointParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "参数非法"
			break
		}

		authToken := util.RandomAlphanumeric(32)
		endpoint, ok := i.endpointHandler.InsertEndpoint(param.ID, param.Name, param.Description, param.User, param.Status, authToken)
		if ok {
			result.Endpoint.Endpoint = endpoint
			result.Endpoint.User = i.accountHandler.GetUsers(param.User)
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
			result.Reason = "新建Endpoint失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type endpointDeleteRoute struct {
	endpointHandler common.EndpointHandler
	accountHandler  common.AccountHandler
}

func (i *endpointDeleteRoute) Method() string {
	return common.DELETE
}

func (i *endpointDeleteRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteEndpoint)
}

func (i *endpointDeleteRoute) Handler() interface{} {
	return i.deleteHandler
}

func (i *endpointDeleteRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *endpointDeleteRoute) deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteHandler")

	result := common_def.DestroyEndpointResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		ok := i.endpointHandler.DeleteEndpoint(strID)
		if ok {
			result.ErrorCode = common_def.Success
		} else {
			result.ErrorCode = common_def.Failed
			result.Reason = "删除Endpoint失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type endpointPutRoute struct {
	endpointHandler common.EndpointHandler
	accountHandler  common.AccountHandler
}

func (i *endpointPutRoute) Method() string {
	return common.PUT
}

func (i *endpointPutRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PutEndpoint)
}

func (i *endpointPutRoute) Handler() interface{} {
	return i.putHandler
}

func (i *endpointPutRoute) AuthGroup() int {
	return common_const.MaintainerAuthGroup.ID
}

func (i *endpointPutRoute) putHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("putHandler")

	result := common_def.UpdateEndpointResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		endpoint, ok := i.endpointHandler.QueryEndpointByID(strID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "对象不存在"
			break
		}

		param := &common_def.UpdateEndpointParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法参数"
			break
		}

		endpoint.Name = param.Name
		endpoint.Description = param.Description
		endpoint.User = param.User
		endpoint.Status = param.Status

		endpoint, ok = i.endpointHandler.UpdateEndpoint(endpoint)
		if ok {
			result.ErrorCode = common_def.Success
			result.Endpoint.Endpoint = endpoint
			result.Endpoint.User = i.accountHandler.GetUsers(param.User)
		} else {
			result.ErrorCode = common_def.Failed
			result.Reason = "更新Endpoint状态失败"
		}

		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

type endpointAuthRoute struct {
	endpointHandler common.EndpointHandler
	sessionRegistry common.SessionRegistry
}

func (i *endpointAuthRoute) Method() string {
	return common.GET
}

func (i *endpointAuthRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetEndpointAuth)
}

func (i *endpointAuthRoute) Handler() interface{} {
	return i.verifyHandler
}

func (i *endpointAuthRoute) AuthGroup() int {
	return common_const.VisitorAuthGroup.ID
}

func (i *endpointAuthRoute) verifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("verifyHandler")

	session := i.sessionRegistry.GetSession(w, r)
	result := common_def.VerifyEndpointResult{}
	for {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		endpoint, ok := i.endpointHandler.QueryEndpointByID(strID)
		if !ok {
			result.ErrorCode = common_def.Failed
			result.Reason = "对象不存在"
			break
		}

		authToken := r.URL.Query().Get(common_const.AuthToken)
		if endpoint.AuthToken != authToken {
			result.ErrorCode = common_def.InvalidAuthority
			result.Reason = "无效授权"
			break
		}

		session.SetOption(common_const.AuthToken, authToken)
		session.SetOption(common_const.ExpiryDate, -1)
		session.Flush()

		result.ErrorCode = common_def.Success
		result.SessionID = session.ID()
		break
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
