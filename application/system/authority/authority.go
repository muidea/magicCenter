package authority

/*
鉴权

实现service.Authority接口，提供管理员鉴权处理器和登陆鉴权处理器，返回martini的鉴权handler

应用端通过System获取接口对象
*/
import (
	"log"
	"net/http"

	"muidea.com/magicCenter/application/common/service"

	"github.com/go-martini/martini"
)

// Authority 权限校验处理器
// 用于在路由过程中进行权限校验
func Authority(authority service.Authority) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
		if authority.Verify(res, req) {
			// 拥有权限，继续执行
			c.Next()
		} else {
			// 没有权限
		}
	}
}

// CreateAuthority 创建Authority
func CreateAuthority() service.Authority {
	impl := &impl{}

	return impl
}

type impl struct {
}

func (i *impl) Verify(res http.ResponseWriter, req *http.Request) bool {
	return true
}
