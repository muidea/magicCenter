package authority

/*
鉴权

实现service.Authority接口，提供管理员鉴权处理器和登陆鉴权处理器，返回martini的鉴权handler

应用端通过System获取接口对象
*/
import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"muidea.com/magicCenter/application/common"
)

// Authority 权限控制服务
// 控制系统各个模块的访问权限
type Authority interface {
	Verify(authorityHandler common.AuthorityHandler, res http.ResponseWriter, req *http.Request) bool
}

// AuthorityVerifyHandler 权限校验处理器
// 用于在路由过程中进行权限校验
func AuthorityVerifyHandler(modHub common.ModuleHub, authority Authority) martini.Handler {
	var authorityHandler common.AuthorityHandler
	mod, ok := modHub.FindModule(common.AuthorityModuleID)
	if ok {
		authorityHandler = nil
		endPoint := mod.EndPoint()
		switch endPoint.(type) {
		case common.AuthorityHandler:
			authorityHandler = endPoint.(common.AuthorityHandler)
		}
	}

	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		if authority.Verify(authorityHandler, res, req) {
			// 拥有权限，继续执行
			c.Next()
		} else {
			// 没有权限
		}
	}
}

// CreateAuthority 创建Authority
func CreateAuthority() Authority {
	impl := &impl{}

	return impl
}

type impl struct {
}

func (i *impl) Verify(authorityHandler common.AuthorityHandler, res http.ResponseWriter, req *http.Request) bool {
	if authorityHandler == nil {
		return false
	}

	return authorityHandler.VerifyAuth(res, req)
}
