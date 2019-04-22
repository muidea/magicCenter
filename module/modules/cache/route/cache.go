package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muidea/magicCenter/common"
	"github.com/muidea/magicCenter/module/modules/cache/def"
	common_const "github.com/muidea/magicCommon/common"
	common_def "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/net"
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

	result := common_def.QueryCacheResult{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	obj, ok := i.cacheHandler.Fetch(id)
	if ok {
		result.Cache = obj
		result.ErrorCode = common_def.Success
	} else {
		result.ErrorCode = common_def.Failed
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
	i := cacheCreateRoute{
		cacheHandler: cacheHandler}
	return &i
}

type cacheCreateRoute struct {
	cacheHandler common.CacheHandler
}

func (i *cacheCreateRoute) Method() string {
	return common.POST
}

func (i *cacheCreateRoute) Pattern() string {
	return net.JoinURL(def.URL, def.PostItem)
}

func (i *cacheCreateRoute) Handler() interface{} {
	return i.postCacheHandler
}

func (i *cacheCreateRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *cacheCreateRoute) postCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postCacheHandler")

	result := common_def.CreateCacheResult{}
	for true {
		param := &common_def.CreateCacheParam{}
		err := net.ParseJSONBody(r, param)
		if err != nil {
			result.ErrorCode = common_def.Failed
			result.Reason = "非法参数"
			break
		}

		if param.Age > 100.0 || param.Age < 0 {
			param.Age = 10
		}

		result.Token = i.cacheHandler.Put(param.Value, float64(param.Age))
		result.ErrorCode = common_def.Success
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
	i := cacheDestroyRoute{
		cacheHandler: cacheHandler}
	return &i
}

type cacheDestroyRoute struct {
	cacheHandler common.CacheHandler
}

func (i *cacheDestroyRoute) Method() string {
	return common.DELETE
}

func (i *cacheDestroyRoute) Pattern() string {
	return net.JoinURL(def.URL, def.DeleteItem)
}

func (i *cacheDestroyRoute) Handler() interface{} {
	return i.deleteCacheHandler
}

func (i *cacheDestroyRoute) AuthGroup() int {
	return common_const.UserAuthGroup.ID
}

func (i *cacheDestroyRoute) deleteCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteCacheHandler")

	result := common_def.DestroyCacheResult{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	i.cacheHandler.Remove(id)
	result.ErrorCode = common_def.Success
	result.Reason = "清除成功"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
