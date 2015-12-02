package modelhelper

import (
	"muidea.com/dao"	
)

type Model interface {
	
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

