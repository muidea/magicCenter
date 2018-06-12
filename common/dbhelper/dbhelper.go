package dbhelper

import (
	"errors"
	"log"

	"muidea.com/magicCommon/foundation/dao"
)

// DBHelper 数据访问助手
type DBHelper interface {
	BeginTransaction()

	Commit()

	Rollback()

	Query(string)

	Next() bool

	Finish()

	GetValue(...interface{})

	Execute(string) (int64, bool)

	Release()
}

type impl struct {
	dao dao.Dao
}

type databaseConfigInfo struct {
	Server   string
	Name     string
	Account  string
	Password string
}

var databaseInfo *databaseConfigInfo

// InitDB 初始化数据库
func InitDB(server, name, account, password string) {
	databaseInfo = &databaseConfigInfo{Server: server, Name: name, Account: account, Password: password}
}

// NewHelper 创建数据助手
func NewHelper() (DBHelper, error) {
	if databaseInfo == nil {
		return nil, errors.New("illegal database config info")
	}

	m := &impl{}
	dao, err := dao.Fetch(databaseInfo.Account, databaseInfo.Password, databaseInfo.Server, databaseInfo.Name)
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

func (db *impl) Finish() {
	db.dao.Finish()
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
