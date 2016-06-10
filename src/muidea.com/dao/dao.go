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

func (this *Dao) BeginTransaction() {
	if martini.Env != martini.Prod {
		log.Print("Begin Transaction")
	}
	
	tx, err := this.dbHandle.Begin()
	if err != nil {
		panic("begin transaction exception, err:" + err.Error())
	}
	
	this.dbTx = tx
}

func (this *Dao) Commit() {
	if martini.Env != martini.Prod {
		log.Print("Commit Transaction")
	}	
	
	if this.dbTx == nil {
		panic("dbTx is nil")
	}
	
	err := this.dbTx.Commit()
	if err != nil {		
		this.dbTx = nil
		
		panic("commit transaction exception, err:" + err.Error())
	}
	
	this.dbTx = nil
}

func (this *Dao) Rollback() {
	if martini.Env != martini.Prod {
		log.Print("Rollback Transaction")
	}
		
	if this.dbTx == nil {
		panic("dbTx is nil")
	}
	
	err := this.dbTx.Rollback()
	if err != nil {		
		this.dbTx = nil
		
		panic("rollback transaction exception, err:" + err.Error())
	}
	
	this.dbTx = nil
}

func (this *Dao) Query(sql string) {
		
	if this.dbHandle == nil {
		panic("dbHanlde is nil")
	}

	if martini.Env != martini.Prod {
		log.Print("query:" + sql)
	}
		
	rows, err := this.dbHandle.Query(sql)
	if err != nil {
		panic("query exception, err:" + err.Error())
	}
	
	this.rowsHandle = rows
}

func (this *Dao) Next() bool {
	if this.rowsHandle == nil {
		panic("rowsHandle is nil");
	}
	
	ret := this.rowsHandle.Next()
	if !ret {
		this.rowsHandle = nil
	}
	
	return ret
}

func (this *Dao) GetField(value ... interface{}){
	if this.rowsHandle == nil {
		panic("rowsHandle is nil");
	}
	
	err := this.rowsHandle.Scan(value...)
	if err != nil {
		panic("scan exception, err:" + err.Error())
	}
}

func (this *Dao) Execute(sql string) (int64,bool) {
	if this.dbHandle == nil {
		panic("dbHandle is nil");
	}
	
	if martini.Env != martini.Prod {
		log.Print("exec:" + sql)
	}
		
	result, err := this.dbHandle.Exec(sql)
	if err != nil {
		panic("exec exception, err:" + err.Error())
	}
	
	num, err := result.RowsAffected()
	if err != nil {
		panic("rows affected exception, err:" + err.Error())
	}
	
	return num,true
}



