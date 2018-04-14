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

func (s *impl) Release() {
	if s.dbTx != nil {
		panic("dbTx isn't nil")
	}

	if s.rowsHandle != nil {
		s.rowsHandle.Close()
	}
	s.rowsHandle = nil

	if s.dbHandle != nil {
		//log.Print("close database connection...")

		s.dbHandle.Close()
	}
	s.dbHandle = nil

}

func (s *impl) BeginTransaction() {
	if s.rowsHandle != nil {
		s.rowsHandle.Close()
	}
	s.rowsHandle = nil

	tx, err := s.dbHandle.Begin()
	if err != nil {
		panic("begin transaction exception, err:" + err.Error())
	}

	s.dbTx = tx
	//log.Print("BeginTransaction")
}

func (s *impl) Commit() {
	if s.dbTx == nil {
		panic("dbTx is nil")
	}

	err := s.dbTx.Commit()
	if err != nil {
		s.dbTx = nil

		panic("commit transaction exception, err:" + err.Error())
	}

	s.dbTx = nil
	//log.Print("Commit")
}

func (s *impl) Rollback() {
	if s.dbTx == nil {
		panic("dbTx is nil")
	}

	err := s.dbTx.Rollback()
	if err != nil {
		s.dbTx = nil

		panic("rollback transaction exception, err:" + err.Error())
	}

	s.dbTx = nil
	//log.Print("Rollback")
}

func (s *impl) Query(sql string) {
	//log.Printf("Query, sql:%s", sql)
	if s.dbTx == nil {
		if s.dbHandle == nil {
			panic("dbHanlde is nil")
		}
		if s.rowsHandle != nil {
			s.rowsHandle.Close()
			s.rowsHandle = nil
		}

		rows, err := s.dbHandle.Query(sql)
		if err != nil {
			panic("query exception, err:" + err.Error() + ", sql:" + sql)
		}
		s.rowsHandle = rows
	} else {

		if s.rowsHandle != nil {
			s.rowsHandle.Close()
			s.rowsHandle = nil
		}

		rows, err := s.dbTx.Query(sql)
		if err != nil {
			panic("query exception, err:" + err.Error() + ", sql:" + sql)
		}
		s.rowsHandle = rows
	}
}

func (s *impl) Next() bool {
	if s.rowsHandle == nil {
		panic("rowsHandle is nil")
	}

	ret := s.rowsHandle.Next()
	if !ret {
		//log.Print("Next, close rows")
		s.rowsHandle.Close()
		s.rowsHandle = nil
	}

	return ret
}

func (s *impl) GetField(value ...interface{}) {
	if s.rowsHandle == nil {
		panic("rowsHandle is nil")
	}

	err := s.rowsHandle.Scan(value...)
	if err != nil {
		panic("scan exception, err:" + err.Error())
	}
}

func (s *impl) Execute(sql string) (int64, bool) {
	if s.rowsHandle != nil {
		s.rowsHandle.Close()
	}
	s.rowsHandle = nil

	if s.dbTx == nil {
		if s.dbHandle == nil {
			panic("dbHandle is nil")
		}

		result, err := s.dbHandle.Exec(sql)
		if err != nil {
			panic("exec exception, err:" + err.Error() + ", sql:" + sql)
		}

		num, err := result.RowsAffected()
		if err != nil {
			panic("rows affected exception, err:" + err.Error())
		}

		return num, true
	}

	result, err := s.dbTx.Exec(sql)
	if err != nil {
		panic("exec exception, err:" + err.Error() + ", sql:" + sql)
	}

	num, err := result.RowsAffected()
	if err != nil {
		panic("rows affected exception, err:" + err.Error())
	}

	return num, true
}
