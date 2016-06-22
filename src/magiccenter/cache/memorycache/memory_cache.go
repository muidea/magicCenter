package memorycache

import (
	"strings"
	"time"

	"muidea.com/util"
)

type commandAction int

const (
	putIn        commandAction = iota // 存放数据
	fetchOut                          // 获取数据
	remove                            // 删除指定数据
	clearAll                          // 清除全部数据
	checkTimeOut                      // 检查超过生命周期的数据
	end                               // 停止Cache
)

type putInData struct {
	data   interface{}
	maxAge float64
}

type putInResult struct {
	value string
}

type fetchOutData struct {
	id string
}

type fetchOutResult struct {
	value interface{}
	found bool
}

type removeData struct {
	id string
}

type cacheData struct {
	putInData
	cacheTime time.Time
}

type commandData struct {
	action commandAction
	value  interface{}
	result chan<- interface{} //单向Channel
}

type memoryCache chan commandData

func NewCache() *memoryCache {
	cache := make(memoryCache)

	go cache.run()

	go cache.checkTimeOut()

	return &cache
}

func DestroyCache(cache interface{}) {
	close(*cache.(*memoryCache))
}

func (right memoryCache) PutIn(data interface{}, maxAge float64) string {

	reply := make(chan interface{})

	putInData := &putInData{}
	putInData.data = data
	putInData.maxAge = maxAge

	right <- commandData{action: putIn, value: putInData, result: reply}

	result := (<-reply).(*putInResult).value
	return result
}

func (right memoryCache) FetchOut(id string) (interface{}, bool) {

	reply := make(chan interface{})

	fetchOutData := &fetchOutData{}
	fetchOutData.id = id

	right <- commandData{action: fetchOut, value: fetchOutData, result: reply}

	result := (<-reply).(*fetchOutResult)
	return result.value, result.found
}

func (right memoryCache) Remove(id string) {
	removeData := &removeData{}
	removeData.id = id

	right <- commandData{action: remove, value: removeData}
}

func (right memoryCache) ClearAll() {

	right <- commandData{action: clearAll}
}

func (right memoryCache) Release() {
	right <- commandData{action: end}
}

func (right memoryCache) run() {
	_cacheData := make(map[string]cacheData)

	for command := range right {
		switch command.action {
		case putIn:
			id := strings.ToLower(util.RandomAlphanumeric(32))

			cacheData := cacheData{}
			cacheData.putInData = *(command.value.(*putInData))
			cacheData.cacheTime = time.Now()

			_cacheData[id] = cacheData

			result := &putInResult{}
			result.value = id

			command.result <- result
		case fetchOut:
			id := command.value.(*fetchOutData).id

			cacheData, found := _cacheData[id]

			result := &fetchOutResult{}
			result.found = found
			if found {
				result.value = cacheData.data
			}

			command.result <- result
		case remove:
			id := command.value.(*removeData).id

			delete(_cacheData, id)
		case clearAll:
			_cacheData = make(map[string]cacheData)

		case checkTimeOut:
			// 检查每项数据是否超时，超时数据需要主动清除掉
			for k, v := range _cacheData {
				current := time.Now()
				elapse := current.Sub(v.cacheTime).Minutes()
				if elapse > v.maxAge {
					delete(_cacheData, k)
				}
			}
		case end:
			_cacheData = nil
		}
	}
}

func (right memoryCache) checkTimeOut() {
	timeOutTimer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timeOutTimer.C:
			right <- commandData{action: checkTimeOut}
		}
	}
}
