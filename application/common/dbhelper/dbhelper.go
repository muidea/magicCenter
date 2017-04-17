package dbhelper

import (
	"log"

	"muidea.com/magicCenter/foundation/dao"
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

type impl struct {
	dao dao.Dao
}

// NewHelper 创建数据助手
func NewHelper() (DBHelper, error) {
	m := &impl{}

	dao, err := dao.Fetch("magiccenter", "magiccenter", "localhost:3306", "magiccenter_db")
	if err != nil {
		log.Print("fetch database failed, err:" + err.Error())
		return nil, err
	}

	m.dao = dao
	return m, err
}

func (db *impl) BeginTransaction() {
	db.dao.BeginTransaction()
}

func (db *impl) Commit() {
	db.dao.Commit()
}

func (db *impl) Rollback() {
	db.dao.Rollback()
}

func (db *impl) Query(sql string) {
	db.dao.Query(sql)
}

func (db *impl) Next() bool {
	return db.dao.Next()
}

func (db *impl) GetValue(val ...interface{}) {
	db.dao.GetField(val...)
}

func (db *impl) Execute(sql string) (int64, bool) {
	return db.dao.Execute(sql)
}

func (db *impl) Release() {
	db.dao.Release()
}
