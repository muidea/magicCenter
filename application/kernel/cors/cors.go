package cors

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"muidea.com/magicCenter/application/common"
)

// Cors cors控制服务
type Cors interface {
	CheckCors(corsHandler common.CorsHandler, res http.ResponseWriter, req *http.Request) bool
}

// CheckHandler Cors处理器
func CheckHandler(moduleHub common.ModuleHub, cors Cors) martini.Handler {
	var corsHandler common.CorsHandler
	corsModule, ok := moduleHub.FindModule(common.CORSModuleID)
	if ok {
		corsHandler = nil
		entryPoint := corsModule.EntryPoint()
		switch entryPoint.(type) {
		case common.CorsHandler:
			corsHandler = entryPoint.(common.CorsHandler)
		}
	}
	if corsHandler == nil {
		panic("can\\'t find CorsHandler")
	}

	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		if cors.CheckCors(corsHandler, res, req) {
			// 拥有权限，继续执行
			c.Next()
		}
	}
}

// CreateCors 创建Cors
func CreateCors() Cors {
	impl := &impl{}

	return impl
}

type impl struct {
}

func (i *impl) CheckCors(corsHandler common.CorsHandler, res http.ResponseWriter, req *http.Request) bool {
	if corsHandler == nil {
		return false
	}

	return corsHandler.CheckCors(res, req)
}
