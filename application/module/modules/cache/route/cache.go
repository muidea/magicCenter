package route

import (
	"encoding/json"
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/modules/cache/def"
	common_const "muidea.com/magicCommon/common"
	common_result "muidea.com/magicCommon/common"
	"muidea.com/magicCommon/foundation/net"
)

// AppendCacheRoute 追加cache 路由
func AppendCacheRoute(routes []common.Route, cacheHandler common.CacheHandler, sessionRegistry common.SessionRegistry) []common.Route {
	// 查询Cache
	rt := CreateQueryCacheRoute(cacheHandler, sessionRegistry)
	routes = append(routes, rt)

	// 提交Cache
	rt = CreatePostCacheRoute(cacheHandler, sessionRegistry)
	routes = append(routes, rt)

	// 删除Cache
	rt = CreateDeleteCacheRoute(cacheHandler, sessionRegistry)
	routes = append(routes, rt)

	return routes
}

// CreateQueryCacheRoute 新建queryCache 路由
func CreateQueryCacheRoute(cacheHandler common.CacheHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := queryCacheRoute{
		cacheHandler: cacheHandler}
	return &i
}

type queryCacheRoute struct {
	cacheHandler common.CacheHandler
}

type cacheQueryResult struct {
	common_result.Result
	Cache interface{} `json:"cache"`
}

func (i *queryCacheRoute) Method() string {
	return common.GET
}

func (i *queryCacheRoute) Pattern() string {
	return net.JoinURL(def.URL, def.GetItem)
}

func (i *queryCacheRoute) Handler() interface{} {
	return i.queryCacheHandler
}

func (i *queryCacheRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *queryCacheRoute) queryCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("queryCacheHandler")

	result := cacheQueryResult{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	obj, ok := i.cacheHandler.FetchOut(id)
	if ok {
		result.Cache = obj
		result.ErrorCode = common_result.Success
	} else {
		result.ErrorCode = common_result.Failed
		result.Reason = "对象不存在"
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// CreatePostCacheRoute 新建postCache 路由
func CreatePostCacheRoute(cacheHandler common.CacheHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := postCacheRoute{
		cacheHandler: cacheHandler}
	return &i
}

type postCacheRoute struct {
	cacheHandler common.CacheHandler
}

type cachePostParam struct {
	Value string `json:"value"`
	Age   int    `json:"age"`
}

type cachePostResult struct {
	common_result.Result
	Token string `json:"token"`
}

func (i *postCacheRoute) Method() string {
	return common.POST
}

func (i *postCacheRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostItem)
}

func (i *postCacheRoute) Handler() interface{} {
	return i.postCacheHandler
}

func (i *postCacheRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *postCacheRoute) postCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postCacheHandler")

	result := cachePostResult{}
	for true {
		param := &cachePostParam{}
		err := net.ParsePostJSON(r, param)
		if err != nil {
			result.ErrorCode = common_result.Failed
			result.Reason = "非法参数"
			break
		}

		if param.Age > 100.0 || param.Age < 0 {
			param.Age = 10
		}

		result.Token = i.cacheHandler.PutIn(param.Value, float64(param.Age))
		result.ErrorCode = common_result.Success
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

// CreateDeleteCacheRoute 新建postCache 路由
func CreateDeleteCacheRoute(cacheHandler common.CacheHandler, sessionRegistry common.SessionRegistry) common.Route {
	i := deleteCacheRoute{
		cacheHandler: cacheHandler}
	return &i
}

type deleteCacheRoute struct {
	cacheHandler common.CacheHandler
}

type cacheDeleteResult struct {
	common_result.Result
}

func (i *deleteCacheRoute) Method() string {
	return common.DELETE
}

func (i *deleteCacheRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteItem)
}

func (i *deleteCacheRoute) Handler() interface{} {
	return i.deleteCacheHandler
}

func (i *deleteCacheRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *deleteCacheRoute) deleteCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteCacheHandler")

	result := common_result.Result{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	i.cacheHandler.Remove(id)
	result.ErrorCode = common_result.Success
	result.Reason = "清除成功"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
