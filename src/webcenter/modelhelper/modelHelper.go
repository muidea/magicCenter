package modelhelper

import (
	"muidea.com/dao"	
)

type Model interface {
	BeginTransaction() bool
	
	Commit() bool
	
	Rollback() bool
	
	Query(string) bool
	
	Next() bool
	
	GetValue(... interface{}) bool
	
	Execute(string) bool
	
	Release()
}

type model struct {
	dao *dao.Dao
}

func NewModel()(Model, error) {
	m := &model{}
	
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		return nil, err
	}
	
	m.dao = dao	
	return m, err
}

func (this *model) BeginTransaction() bool {
	return this.dao.BeginTransaction()
}

func (this *model) Commit() bool {
	return this.dao.Commit()
}

func (this *model) Rollback() bool {
	return this.dao.Rollback()
}

func (this *model)Query(sql string) bool {
	return this.dao.Query(sql)
}

func (this *model)Next() bool {
	return this.dao.Next()
}

func (this *model)GetValue(val ... interface{}) bool {
	return this.dao.GetField(val...)
}

func (this *model)Execute(sql string) bool {
	return this.dao.Execute(sql)
}

func (this *model)Release() {
	this.dao.Release()
} 

