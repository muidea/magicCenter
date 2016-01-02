package dao

import (
	"fmt"
	"log"	
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "martini"
)

type Dao struct {
	dbHandle *sql.DB
	dbTx *sql.Tx
	rowsHandle *sql.Rows
}

func Fetch(user string, password string, address string, dbName string) (*Dao, error) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user,password,address,dbName)
	
	if martini.Env != martini.Prod {
		log.Print(connectStr)
	}
	
	dao := Dao{dbHandle:nil, dbTx:nil, rowsHandle:nil}
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		panic("open database exception, err:" + err.Error())
	} else {
		dao.dbHandle = db		
	}
	
	return &dao,err
}

func (this *Dao) Release() {
	if martini.Env != martini.Prod {
		log.Print("close database connection")
	}
		
	if this.rowsHandle != nil {
		this.rowsHandle.Close()
	}
	this.rowsHandle = nil
	
	if this.dbHandle != nil {
		this.dbHandle.Close()
	}
	this.dbHandle = nil
		
}

func (this *Dao) BeginTransaction() bool {
	if martini.Env != martini.Prod {
		log.Print("Begin Transaction")
	}
	
	tx, err := this.dbHandle.Begin()
	if err != nil {
		panic("begin transaction exception, err:" + err.Error())
	}
	
	this.dbTx = tx
	return true
}

func (this *Dao) Commit() bool {
	if martini.Env != martini.Prod {
		log.Print("Commit Transaction")
	}	
	
	if this.dbTx == nil {
		panic("dbTx is nil")
	}
	
	err := this.dbTx.Commit()
	if err != nil {
		panic("commit transaction exception, err:" + err.Error())
		
		this.dbTx = nil
		return false;
	}
	
	this.dbTx = nil
	return true
}

func (this *Dao) Rollback() bool {
	if martini.Env != martini.Prod {
		log.Print("Rollback Transaction")
	}
		
	if this.dbTx == nil {
		panic("dbTx is nil")
	}
	
	err := this.dbTx.Rollback()
	if err != nil {		
		panic("rollback transaction exception, err:" + err.Error())
		
		this.dbTx = nil
		return false;
	}
	
	this.dbTx = nil
	return true
}

func (this *Dao) Query(sql string) bool {
		
	if this.dbHandle == nil {
		return false
	}

	if martini.Env != martini.Prod {
		log.Print("query:" + sql)
	}
		
	rows, err := this.dbHandle.Query(sql)
	if err != nil {
		panic("query exception, err:" + err.Error())
		return false
	}
	
	this.rowsHandle = rows
	
	return true		
}

func (this *Dao) Next() bool {
	if this.rowsHandle == nil {
		return false
	}
	ret := this.rowsHandle.Next()
	if !ret {
		this.rowsHandle = nil
	}
	
	return ret
}

func (this *Dao) GetField(value ... interface{}) bool {
	if this.rowsHandle == nil {
		panic("rowsHandle is nil");
	}
	
	err := this.rowsHandle.Scan(value...)
	if err != nil {
		panic("scan exception, err:" + err.Error())
	}
	
	return true
}

func (this *Dao) Execute(sql string) bool {
	if this.dbHandle == nil {
		return false
	}
	
	if martini.Env != martini.Prod {
		log.Print("exec:" + sql)
	}
		
	result, err := this.dbHandle.Exec(sql)
	if err != nil {
		panic("exec exception, err:" + err.Error())
	}
	
	_, err = result.RowsAffected()
	if err != nil {
		panic("rows affected exception, err:" + err.Error())
	}
	
	return true	
}



