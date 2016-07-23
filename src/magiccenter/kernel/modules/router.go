package modules

import (
	"magiccenter/kernel/modules/account"
	"magiccenter/kernel/modules/content"
)

// RegisterRouter 注册路由
func RegisterRouter() {
	account.RegisterRouter()

	content.RegisterRouter()
}
