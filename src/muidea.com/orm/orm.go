package orm

import (
	"muidea.com/dao"
)

func init() {
	initialize()
}


type DBObject interface {
	query(dao *dao.Dao) bool
	insert(dao *dao.Dao) bool
	update(dao *dao.Dao) bool
	remove(dao *dao.Dao)
}


type Orm struct {
	dao *dao.Dao
}

func New(dao *dao.Dao) *Orm {
	return &Orm{dao:dao}
}

func (this *Orm)Insert(obj DBObject) bool {
	return instance.insert(this.dao,obj)
}

func (this *Orm)Update(obj DBObject) bool {
	return instance.insert(this.dao,obj)
}

func (this *Orm)Query(obj DBObject) bool {
	return instance.query(this.dao,obj)
}

func (this *Orm)Remove(obj DBObject) {
	instance.remove(this.dao,obj)
}

