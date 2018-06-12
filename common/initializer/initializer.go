package initializer

// Handler handler
type Handler interface {
	Handle()
}

var handlerList = []Handler{}

// RegisterHandler 注册启动Handler
func RegisterHandler(handler Handler) {
	handlerList = append(handlerList, handler)
}

// InvokHandler 执行Handler
func InvokHandler() {
	for _, val := range handlerList {
		val.Handle()
	}
}
