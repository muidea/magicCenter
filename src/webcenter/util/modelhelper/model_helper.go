package modelhelper

import (
	"muidea.com/dao"	
)

type Model interface {
	BeginTransaction()
	
	Commit()
	
	Rollback()
	
	Query(string)
	
	Next() bool
	
	GetValue(... interface{})
	
	Execute(string) (int64,bool)
	
	Release()
}

type model struct {
	dao *dao.Dao
}

func NewHelper()(Model, error) {
	m := &model{}
	
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		return nil, err
	}
	
	m.dao = dao	
	return m, err
}

func (this *model) BeginTransaction() {
	this.dao.BeginTransaction()
}

func (this *model) Commit() {
	this.dao.Commit()
}

func (this *model) Rollback() {
	this.dao.Rollback()
}

func (this *model)Query(sql string) {
	this.dao.Query(sql)
}

func (this *model)Next() bool {
	return this.dao.Next()
}

func (this *model)GetValue(val ... interface{}) {
	this.dao.GetField(val...)
}

func (this *model)Execute(sql string) (int64, bool) {
	return this.dao.Execute(sql)
}

func (this *model)Release() {
	this.dao.Release()
} 

