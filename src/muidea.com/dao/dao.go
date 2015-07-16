package dao

import (
	"fmt"
	"log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	dbHandle *sql.DB
	rowsHandle *sql.Rows
}

func Fetch(user string, password string, address string, dbName string) (*Dao, error) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user,password,address,dbName)
	
	dao := Dao{}
	db, err := sql.Open("mysql", connectStr)
	if err != nil {
		log.Printf("open database failed, err:%s", err.Error())
	} else {
		dao.dbHandle = db		
	}
	
	return &dao,err
}

func (this *Dao) Release() {
	if this.rowsHandle != nil {
		this.rowsHandle.Close()
	}
	this.rowsHandle = nil
	
	if this.dbHandle != nil {
		this.dbHandle.Close()
	}
	this.dbHandle = nil
		
}

func (this *Dao) Query(sql string) bool {
	if this.dbHandle == nil {
		return false
	}
	
	rows, err := this.dbHandle.Query(sql)
	if err != nil {
		log.Printf("query failed, err:%s", err.Error())
		return false
	}
	
	this.rowsHandle = rows
	
	return true		
}

func (this *Dao) Next() bool {
	if this.rowsHandle == nil {
		return false
	}
	
	return this.rowsHandle.Next()
}

func (this *Dao) GetField(value ... interface{}) bool {
	err := this.rowsHandle.Scan(value...)
	if err != nil {
		log.Printf("get field failed, err:%s", err.Error())
		return false
	}
	
	return true
}

func (this *Dao) Execute(sql string) bool {
	if this.dbHandle == nil {
		return false
	}
	
	result, err := this.dbHandle.Exec(sql)
	if err != nil {
		log.Printf("exec failed, err:%s", err.Error())
		return false
	}
	
	_, err = result.RowsAffected()
	if err != nil {
		log.Printf("exec failed, err:%s", err.Error())
		return false
	}
	
	return true	
}



