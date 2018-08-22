package initializer

import (
	"log"
	"os"

	"muidea.com/magicCenter/common/dbhelper"
)

// Handler handler
type Handler interface {
	Handle()
}

var handlerList = []Handler{}

// RegisterHandler 注册启动Handler
func RegisterHandler(handler Handler) {
	handlerList = append(handlerList, handler)
}

// Initialize 初始化
func Initialize(bindPort, server, name, account, password string) error {
	os.Setenv("PORT", bindPort)

	dbhelper.InitDB(server, name, account, password)

	helper, err := dbhelper.NewHelper()
	defer helper.Release()
	if err != nil {
		errCode, _ := dbhelper.ParseError(err.Error())
		if errCode == 1049 {
			dbFile := "db.sql"
			_, err := os.Stat(dbFile)
			if err == nil {
				// 如果是数据库不存在，则导入数据库
				err = loadDatabase(server, name, account, password, dbFile)
				if err == nil {
					os.Remove(dbFile)
				}
			}
		} else {
			log.Printf("fetch dbhelper exception, errCode:%d, err:%s", errCode, err.Error())
		}
	}

	return err
}

// InvokHandler 执行Handler
func InvokHandler() {
	for _, val := range handlerList {
		val.Handle()
	}
}
