package dbhelper

import (
	"muidea.com/dao"
)

// DBHelper 数据访问助手
type DBHelper interface {
	BeginTransaction()

	Commit()

	Rollback()

	Query(string)

	Next() bool

	GetValue(...interface{})

	Execute(string) (int64, bool)

	Release()
}

type dbHelper struct {
	dao *dao.Dao
}

// NewHelper 创建数据助手
func NewHelper() (DBHelper, error) {
	m := &dbHelper{}

	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magiccenter_db")
	if err != nil {
		return nil, err
	}

	m.dao = dao
	return m, err
}

func (db *dbHelper) BeginTransaction() {
	db.dao.BeginTransaction()
}

func (db *dbHelper) Commit() {
	db.dao.Commit()
}

func (db *dbHelper) Rollback() {
	db.dao.Rollback()
}

func (db *dbHelper) Query(sql string) {
	db.dao.Query(sql)
}

func (db *dbHelper) Next() bool {
	return db.dao.Next()
}

func (db *dbHelper) GetValue(val ...interface{}) {
	db.dao.GetField(val...)
}

func (db *dbHelper) Execute(sql string) (int64, bool) {
	return db.dao.Execute(sql)
}

func (db *dbHelper) Release() {
	db.dao.Release()
}
