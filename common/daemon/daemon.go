package daemon

import (
	"log"
	"time"
)

// Handler handler
type Handler interface {
	Handle()
}

var runningFlag = true
var handlerList = []Handler{}

// RegisterTimerHandler 注册TimerHandler
func RegisterTimerHandler(handler Handler) {
	handlerList = append(handlerList, handler)
}

// Start 初始化
func Start() {
	go checkTimer()
}

// Stop 停止
func Stop() {
	runningFlag = false
}

func checkTimer() {
	timeOutTimer := time.NewTicker(1 * time.Minute)
	for runningFlag {
		select {
		case <-timeOutTimer.C:
			go invoker()
		}
	}
}

func invoker() {
	defer func() {
		if err := recover(); err != nil {
			stack := stack(3)
			log.Printf("PANIC: %s\n%s", err, stack)
		}
	}()

	for _, val := range handlerList {
		val.Handle()
	}
}
