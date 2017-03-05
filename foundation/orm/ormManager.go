package orm

import "muidea.com/magicCenter/foundation/dao"

type commandData struct {
	action commandAction // 做什么
	dao    *dao.Dao
	value  DBObject // 操作对象
	result chan<- interface{}
}

type commandAction int

const (
	insert commandAction = iota
	remove
	update
	find
	end
)

type findResult struct {
	value interface{}
	found bool
}

type ormManager chan commandData

var instance ormManager = nil

func initialize() {

	instance = make(ormManager)

	go instance.run()
}

func (instance ormManager) query(dao *dao.Dao, obj DBObject) bool {

	reply := make(chan interface{})
	instance <- commandData{action: find, dao: dao, value: obj, result: reply}

	result := (<-reply).(bool)
	return result
}

func (instance ormManager) insert(dao *dao.Dao, obj DBObject) bool {

	reply := make(chan interface{})
	instance <- commandData{action: insert, dao: dao, value: obj, result: reply}

	result := (<-reply).(bool)
	return result
}

func (instance ormManager) update(dao *dao.Dao, obj DBObject) bool {

	reply := make(chan interface{})
	instance <- commandData{action: update, dao: dao, value: obj, result: reply}

	result := (<-reply).(bool)
	return result
}

func (instance ormManager) remove(dao *dao.Dao, obj DBObject) {

	instance <- commandData{action: remove, dao: dao, value: obj}
}

func (instance ormManager) run() {
	for command := range instance {
		switch command.action {
		case insert:
			obj := command.value.(DBObject)
			dao := command.dao
			command.result <- obj.insert(dao)
		case remove:
			obj := command.value.(DBObject)
			dao := command.dao
			obj.remove(dao)
		case update:
			obj := command.value.(DBObject)
			dao := command.dao
			command.result <- obj.update(dao)
		case find:
			obj := command.value.(DBObject)
			dao := command.dao
			command.result <- obj.query(dao)
		case end:
			close(instance)
		}
	}
}
