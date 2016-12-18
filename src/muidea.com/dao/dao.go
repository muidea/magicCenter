package dao

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql" //引入Mysql驱动
)

// Dao 数据库访问对象
type Dao interface {
	Release()
	BeginTransaction()
	Commit()
	Rollback()
	Query(sql string)
	Next() bool
	GetField(value ...interface{})
	Execute(sql string) (int64, bool)
}

type impl struct {
	dbHandle   *sql.DB
	dbTx       *sql.Tx
	rowsHandle *sql.Rows
}

// Fetch 获取一个数据访问对象
func Fetch(user string, password string, address string, dbName string) (Dao, error) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, address, dbName)

	if martini.Env != martini.Prod {
		log.Print(connectStr)
	}

	dao := &impl{dbHandle: nil, dbTx: nil, rowsHandle: nil}
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		panic("open database exception, err:" + err.Error())
	} else {
		dao.dbHandle = db
	}

	return dao, err
}

func (impl *impl) Release() {
	if martini.Env != martini.Prod {
		log.Print("close database connection")
	}

	if impl.rowsHandle != nil {
		impl.rowsHandle.Close()
	}
	impl.rowsHandle = nil

	if impl.dbHandle != nil {
		impl.dbHandle.Close()
	}
	impl.dbHandle = nil

}

func (impl *impl) BeginTransaction() {
	if martini.Env != martini.Prod {
		log.Print("Begin Transaction")
	}

	tx, err := impl.dbHandle.Begin()
	if err != nil {
		panic("begin transaction exception, err:" + err.Error())
	}

	impl.dbTx = tx
}

func (impl *impl) Commit() {
	if martini.Env != martini.Prod {
		log.Print("Commit Transaction")
	}

	if impl.dbTx == nil {
		panic("dbTx is nil")
	}

	err := impl.dbTx.Commit()
	if err != nil {
		impl.dbTx = nil

		panic("commit transaction exception, err:" + err.Error())
	}

	impl.dbTx = nil
}

func (impl *impl) Rollback() {
	if martini.Env != martini.Prod {
		log.Print("Rollback Transaction")
	}

	if impl.dbTx == nil {
		panic("dbTx is nil")
	}

	err := impl.dbTx.Rollback()
	if err != nil {
		impl.dbTx = nil

		panic("rollback transaction exception, err:" + err.Error())
	}

	impl.dbTx = nil
}

func (impl *impl) Query(sql string) {

	if impl.dbHandle == nil {
		panic("dbHanlde is nil")
	}

	if martini.Env != martini.Prod {
		log.Print("query:" + sql)
	}

	rows, err := impl.dbHandle.Query(sql)
	if err != nil {
		panic("query exception, err:" + err.Error())
	}

	impl.rowsHandle = rows
}

func (impl *impl) Next() bool {
	if impl.rowsHandle == nil {
		panic("rowsHandle is nil")
	}

	ret := impl.rowsHandle.Next()
	if !ret {
		impl.rowsHandle = nil
	}

	return ret
}

func (impl *impl) GetField(value ...interface{}) {
	if impl.rowsHandle == nil {
		panic("rowsHandle is nil")
	}

	err := impl.rowsHandle.Scan(value...)
	if err != nil {
		panic("scan exception, err:" + err.Error())
	}
}

func (impl *impl) Execute(sql string) (int64, bool) {
	if impl.dbHandle == nil {
		panic("dbHandle is nil")
	}

	if martini.Env != martini.Prod {
		log.Print("exec:" + sql)
	}

	result, err := impl.dbHandle.Exec(sql)
	if err != nil {
		panic("exec exception, err:" + err.Error())
	}

	num, err := result.RowsAffected()
	if err != nil {
		panic("rows affected exception, err:" + err.Error())
	}

	return num, true
}
