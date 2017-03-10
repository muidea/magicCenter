package cache

import (
	"net/http"

	"log"

	"muidea.com/magicCenter/application/common"
	"muidea.com/magicCenter/application/common/configuration"
	"muidea.com/magicCenter/application/common/model"
	"muidea.com/magicCenter/application/kernel/modulehub"
	"muidea.com/magicCenter/application/module/kernel/modules/cache/memorycache"
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
func LoadModule(cfg configuration.Configuration, modHub modulehub.ModuleHub) {
	instance := &cacheModule{cache: memorycache.NewCache()}

	modHub.RegisterModule(instance)
}

// PutIn 存放数据
func PutIn(modHub modulehub.ModuleHub, data interface{}, maxAge float64) string {
	cacheModule, found := modHub.FindModule(ID)
	if !found {
		panic("can't find mail module")
	}

	endPoint := cacheModule.EndPoint().(Cache)

	return endPoint.PutIn(data, maxAge)
}

// FetchOut 取缓存数据
func FetchOut(modHub modulehub.ModuleHub, id string) (interface{}, bool) {
	cacheModule, found := modHub.FindModule(ID)
	if !found {
		panic("can't find mail module")
	}

	endPoint := cacheModule.EndPoint().(Cache)
	return endPoint.FetchOut(id)
}

// Remove 清除指定的缓存数据
func Remove(modHub modulehub.ModuleHub, id string) {
	cacheModule, found := modHub.FindModule(ID)
	if !found {
		panic("can't find mail module")
	}

	endPoint := cacheModule.EndPoint().(Cache)
	endPoint.Remove(id)
}

// ClearAll 清空全部缓存数据
func ClearAll(modHub modulehub.ModuleHub) {
	cacheModule, found := modHub.FindModule(ID)
	if !found {
		panic("can't find mail module")
	}

	endPoint := cacheModule.EndPoint().(Cache)
	endPoint.ClearAll()
}

// Cache 对象，由于系统临时保存信息
// Cache会返回一个string用于应用来获取临时保存的对象
//  存放的对象是有生命周期的，超过设定的存放时间会被系统清除掉
// maxAge 单位为minute
type Cache interface {
	PutIn(data interface{}, maxAge float64) string
	FetchOut(id string) (interface{}, bool)
	Remove(id string)
	ClearAll()
}

type cacheModule struct {
	cache Cache
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
		common.NewRoute(common.GET, "[a-zA-Z0-9]*/", getCacheActionHandler),
		common.NewRoute(common.POST, "", postCacheActionHandler),
		common.NewRoute(common.DELETE, "[a-zA-Z0-9]*/", deleteCacheActionHandler),
	}

	return routes
}

// Startup 启动Cache模块
func (instance *cacheModule) Startup() bool {
	return true
}

// Cleanup 清除Cache模块
func (instance *cacheModule) Cleanup() {
	memorycache.DestroyCache(instance.cache)
}

func getCacheActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("getCacheActionHandler")

}

func postCacheActionHandler(w http.ResponseWriter, r *http.Request) {
}

func deleteCacheActionHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.RawPath)
}
