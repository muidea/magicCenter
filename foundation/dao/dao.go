package dao

import (
	"database/sql"
	"fmt"

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

	i := impl{dbHandle: nil, dbTx: nil, rowsHandle: nil}
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		panic("open database exception, err:" + err.Error())
	} else {
		//log.Print("open database connection...")
		i.dbHandle = db
	}

	return &i, err
}

func (impl *impl) Release() {

	if impl.rowsHandle != nil {
		impl.rowsHandle.Close()
	}
	impl.rowsHandle = nil

	if impl.dbHandle != nil {
		//log.Print("close database connection...")

		impl.dbHandle.Close()
	}
	impl.dbHandle = nil

}

func (impl *impl) BeginTransaction() {
	tx, err := impl.dbHandle.Begin()
	if err != nil {
		panic("begin transaction exception, err:" + err.Error())
	}

	impl.dbTx = tx
	// log.Print("BeginTransaction")
}

func (impl *impl) Commit() {
	if impl.dbTx == nil {
		panic("dbTx is nil")
	}

	err := impl.dbTx.Commit()
	if err != nil {
		impl.dbTx = nil

		panic("commit transaction exception, err:" + err.Error())
	}

	impl.dbTx = nil
	// log.Print("Commit")
}

func (impl *impl) Rollback() {
	if impl.dbTx == nil {
		panic("dbTx is nil")
	}

	err := impl.dbTx.Rollback()
	if err != nil {
		impl.dbTx = nil

		panic("rollback transaction exception, err:" + err.Error())
	}

	impl.dbTx = nil
	// log.Print("Rollback")
}

func (impl *impl) Query(sql string) {

	if impl.dbHandle == nil {
		panic("dbHanlde is nil")
	}
	if impl.rowsHandle != nil {
		impl.rowsHandle.Close()
		impl.rowsHandle = nil
	}

	rows, err := impl.dbHandle.Query(sql)
	if err != nil {
		panic("query exception, err:" + err.Error() + ", sql:" + sql)
	}
	impl.rowsHandle = rows
}

func (impl *impl) Next() bool {
	if impl.rowsHandle == nil {
		panic("rowsHandle is nil")
	}

	ret := impl.rowsHandle.Next()
	if !ret {
		impl.rowsHandle.Close()
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

	result, err := impl.dbHandle.Exec(sql)
	if err != nil {
		panic("exec exception, err:" + err.Error() + ", sql:" + sql)
	}

	num, err := result.RowsAffected()
	if err != nil {
		panic("rows affected exception, err:" + err.Error())
	}

	return num, true
}
