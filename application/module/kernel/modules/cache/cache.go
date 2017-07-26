package cache

import (
	"encoding/json"
	"net/http"

	"log"

	"strconv"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/def"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/handler"
	"muidea.com/magicCenter/foundation/net"
)

// LoadModule 加载Cache模块
func LoadModule(cfg common.Configuration, modHub common.ModuleHub) {
	instance := &cacheModule{cacheHandler: handler.CreateCacheHandler(cfg, modHub)}

	modHub.RegisterModule(instance)
}

type cacheModule struct {
	cacheHandler common.CacheHandler
}

func (instance *cacheModule) ID() string {
	return def.ID
}

func (instance *cacheModule) Name() string {
	return def.Name
}

func (instance *cacheModule) Description() string {
	return def.Description
}

func (instance *cacheModule) Group() string {
	return "util"
}

func (instance *cacheModule) Type() int {
	return common.KERNEL
}

func (instance *cacheModule) Status() int {
	return 0
}

func (instance *cacheModule) EntryPoint() interface{} {
	return instance.cacheHandler
}

func (instance *cacheModule) AuthGroups() []model.AuthGroup {
	groups := []model.AuthGroup{}

	return groups
}

// Route Cache 路由信息
func (instance *cacheModule) Routes() []common.Route {
	routes := []common.Route{
		common.NewRoute(common.GET, net.JoinURL(def.URL, "[a-zA-Z0-9]+/"), instance.getCacheActionHandler),
		common.NewRoute(common.POST, net.JoinURL(def.URL, ""), instance.postCacheActionHandler),
		common.NewRoute(common.DELETE, net.JoinURL(def.URL, "[a-zA-Z0-9]+/"), instance.deleteCacheActionHandler),
	}

	return routes
}

// Startup 启动Cache模块
func (instance *cacheModule) Startup() bool {
	return true
}

// Cleanup 清除Cache模块
func (instance *cacheModule) Cleanup() {
}

type cacheResult struct {
	common.Result
	Cache interface{}
}

func (instance *cacheModule) getCacheActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getCacheActionHandler")

	result := cacheResult{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	obj, ok := instance.cacheHandler.FetchOut(id)
	if ok {
		result.Cache = obj
		result.ErrCode = 0
		result.Reason = "查询成功"
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

		result.Cache = instance.cacheHandler.PutIn(value, float64(age))
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
	log.Print(r.URL.Path)
	result := common.Result{}
	_, id := net.SplitRESTAPI(r.URL.Path)
	log.Print(id)
	instance.cacheHandler.Remove(id)
	result.ErrCode = 0
	result.Reason = "清除成功"

	b, err := json.Marshal(result)
	if err != nil {
		panic("json.Marshal, failed, err:" + err.Error())
	}

	w.Write(b)
}
