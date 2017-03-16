package cache

import (
	"encoding/json"
	"net/http"

	"log"

	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/foundation/cache"
	"muidea.com/magicCenter/foundation/net"
)

// ID 模块ID
const ID = "0424492f-420a-42fb-9106-3882c07bf99e"

// Name 块名称
const Name = "Magic Cache"

// Description 模块描述信息
const Description = "Magic 缓存模块"

// URL 模块Url
const URL string = "/cache"

// LoadModule 加载Cache模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	instance := &cacheModule{cache: cache.NewCache()}

	modHub.RegisterModule(instance)
}

type cacheModule struct {
	cache cache.Cache
}

func (instance *cacheModule) ID() string {
	return ID
}

func (instance *cacheModule) Name() string {
	return Name
}

func (instance *cacheModule) Description() string {
	return Description
}

func (instance *cacheModule) Group() string {
	return "util"
}

func (instance *cacheModule) Type() int {
	return common.KERNEL
}

func (instance *cacheModule) URL() string {
	return URL
}

func (instance *cacheModule) Status() int {
	return 0
}

func (instance *cacheModule) EndPoint() interface{} {
	return instance.cache
}

func (instance *cacheModule) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route Cache 路由信息
func (instance *cacheModule) Routes() []common.Route {
	routes := []common.Route{
		common.NewRoute(common.GET, "[a-zA-Z0-9]*/", instance.getCacheActionHandler),
		common.NewRoute(common.POST, "", instance.postCacheActionHandler),
		common.NewRoute(common.DELETE, "[a-zA-Z0-9]*/", instance.deleteCacheActionHandler),
	}

	return routes
}

// Startup 启动Cache模块
func (instance *cacheModule) Startup() bool {
	return true
}

// Cleanup 清除Cache模块
func (instance *cacheModule) Cleanup() {
	instance.cache.Release()
}

type cacheResult struct {
	common.Result
	Cache interface{}
}

func (instance *cacheModule) getCacheActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getCacheActionHandler")

	result := cacheResult{}
	id, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	if ok {
		obj, ok := instance.cache.FetchOut(id)
		if ok {
			result.Cache = obj
			result.ErrCode = 0
			result.Reason = "查询成功"
		} else {
			result.ErrCode = 1
			result.Reason = "对象不存在"
		}
	} else {
		result.ErrCode = 1
		result.Reason = "参数非法"
	}
	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

func (instance *cacheModule) postCacheActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postCacheActionHandler")

	result := cacheResult{}
	for true {
		err := r.ParseForm()
		if err != nil {
			result.ErrCode = 1
			result.Reason = "非法参数"
			break
		}

		value := r.FormValue("cache-value")
		age, err := strconv.Atoi(r.FormValue("cache-age"))
		if err != nil {
			age = 10
		} else if age > 100.0 || age < 0 {
			age = 10
		}

		result.Cache = instance.cache.PutIn(value, float64(age))
		break
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}

func (instance *cacheModule) deleteCacheActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("deleteCacheActionHandler")

	result := common.Result{}
	id, _, ok := net.ParseRestAPIUrl(r.URL.Path)
	if ok {
		instance.cache.Remove(id)
		result.ErrCode = 0
		result.Reason = "清除成功"
	} else {
		result.ErrCode = 1
		result.Reason = "参数非法"
	}

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
