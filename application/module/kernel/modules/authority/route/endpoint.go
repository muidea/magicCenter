package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/authority/def"
	common_const "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/net"
	"muidea.com/magicCommon/foundation/util"
	"muidea.com/magicCommon/model"
)

// CreateQueryEndpointRoute QueryEndpoint
func CreateQueryEndpointRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointQueryRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

// CreatePostEndpointRoute CreateEndpoint
func CreatePostEndpointRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointPostRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

// CreateDeleteEndpointRoute DeleteEndpoint
func CreateDeleteEndpointRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {
	i := endpointDeleteRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

// CreatePutEndpointRoute UpdateEndpoint
func CreatePutEndpointRoute(authorityHandler common.AuthorityHandler, accountHandler common.AccountHandler) common.Route {

	i := endpointPutRoute{authorityHandler: authorityHandler, accountHandler: accountHandler}
	return &i
}

type endpointQueryRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type endpointQueryResult struct {
	common_result.Result
	Endpoint []model.EndpointView `json:"endpoint"`
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
	result := endpointQueryResult{}

	for {
		endpoints := i.authorityHandler.QueryAllEndpoint()
		for _, val := range endpoints {
			endpoint := model.EndpointView{}
			endpoint.Endpoint = val
			endpoint.User = i.accountHandler.GetUsers(val.User)

			result.Endpoint = append(result.Endpoint, endpoint)
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

type endpointPostRoute struct {
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type endpointPostParam struct {
	ID          string `json:"url"`
	Name        string `json:"method"`
	Description string `json:"module"`
	User        []int  `json:"user"`
	Status      int    `json:"status"`
}

type endpointPostResult struct {
	common_result.Result
	Endpoint model.EndpointView `json:"endpoint"`
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

	result := endpointPostResult{}
	for true {
		param := &endpointPostParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "参数非法"
			break
		}

		accessToken := util.RandomAlphanumeric(32)
		endpoint, ok := i.authorityHandler.InsertEndpoint(param.ID, param.Name, param.Description, param.User, param.Status, accessToken)
		if ok {
			result.Endpoint.Endpoint = endpoint
			result.Endpoint.User = i.accountHandler.GetUsers(param.User)
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type endpointDeleteResult struct {
	common_result.Result
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

	result := endpointDeleteResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		ok := i.authorityHandler.DeleteEndpoint(strID)
		if ok {
			result.ErrorCode = common_result.Success
		} else {
			result.ErrorCode = common_result.Failed
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
	authorityHandler common.AuthorityHandler
	accountHandler   common.AccountHandler
}

type endpointPutParam struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        []int  `json:"user"`
	Status      int    `json:"status"`
}

type endpointPutResult struct {
	common_result.Result
	Endpoint model.EndpointView `json:"endpoint"`
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

	result := endpointPutResult{}
	for true {
		_, strID := net.SplitRESTAPI(r.URL.Path)
		endpoint, ok := i.authorityHandler.QueryEndpointByID(strID)
		if !ok {
			result.ErrorCode = common_result.Failed
			result.Reason = "对象不存在"
			break
		}

		param := &endpointPutParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		endpoint.Name = param.Name
		endpoint.Description = param.Description
		endpoint.User = param.User
		endpoint.Status = param.Status

		endpoint, ok = i.authorityHandler.UpdateEndpoint(endpoint)
		if ok {
			result.ErrorCode = common_result.Success
			result.Endpoint.Endpoint = endpoint
			result.Endpoint.User = i.accountHandler.GetUsers(param.User)
		} else {
			result.ErrorCode = common_result.Failed
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
