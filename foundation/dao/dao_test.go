package dao_test

import (
	"fmt"
	"log"
	"testing"

	"muidea.com/magicCenter/foundation/dao"
)

type User struct {
	id       int
	name     string
	password string
	catalog  int
}

func TestInser(t *testing.T) {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		t.Errorf("Fetch dao failed, err:%s", err.Error())
	}
	defer dao.Release()

	insertSql := fmt.Sprintf("%s", "insert into magicid_db.user value(4,'test3','test3',1)")
	if !dao.Execute(insertSql) {
		t.Errorf("Insert data failed")
	}

}

func TestQuery(t *testing.T) {
	dao, err := dao.Fetch("root", "rootkit", "localhost:3306", "magicid_db")
	if err != nil {
		t.Errorf("Fetch dao failed, err:%s", err.Error())
	}
	defer dao.Release()

	selectSql := fmt.Sprint("select id,name,password,type from magicid_db.user")

	if !dao.Query(selectSql) {
		t.Errorf("Query all data failed")
	}

	for dao.Next() {
		user := User{}

		if !dao.GetField(&user.id, &user.name, &user.password, &user.catalog) {
			t.Errorf("Get Fileds failed")
		} else {
			log.Printf("id:%d, name:%s, pass:%s, type:%d", user.id, user.name, user.password, user.catalog)
		}
	}
}
