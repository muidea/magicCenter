package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/def"
	"muidea.com/magicCenter/foundation/net"
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
	common.Result
	Cache interface{}
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
	return common.UserAuthGroup.ID
}

func (i *queryCacheRoute) queryCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("queryCacheHandler")

	result := cacheQueryResult{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	obj, ok := i.cacheHandler.FetchOut(id)
	if ok {
		result.Cache = obj
		result.ErrCode = 0
	} else {
		result.ErrCode = 1
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

type cachePostResult struct {
	common.Result
	Token string
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
	return common.UserAuthGroup.ID
}

func (i *postCacheRoute) postCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postCacheHandler")

	result := cachePostResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		value := r.FormValue("value")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			age = 10
		} else if age > 100.0 || age < 0 {
			age = 10
		}

		result.Token = i.cacheHandler.PutIn(value, float64(age))
		result.ErrCode = 0
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
	common.Result
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
	return common.UserAuthGroup.ID
}

func (i *deleteCacheRoute) deleteCacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteCacheHandler")

	result := common.Result{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	i.cacheHandler.Remove(id)
	result.ErrCode = 0
	result.Reason = "清除成功"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
