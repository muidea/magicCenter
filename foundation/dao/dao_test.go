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

func testFun(t *testing.T) {
	dao, err := dao.Fetch("magiccenter", "magiccenter", "localhost:3306", "magiccenter_db")
	if err != nil {
		t.Errorf("Fetch dao failed, err:%s", err.Error())
	}
	defer dao.Release()

	insertSql := fmt.Sprintf("%s", "insert into magiccenter_db.user value(1,'test3','test3','test11','aaa',1)")
	num, ok := dao.Execute(insertSql)
	if num == 1 && ok {
		t.Errorf("Insert data failed")
	}

	querySql := "select * from magiccenter_db.user where id=1"
	dao.Query(querySql)
}

func TestInser(t *testing.T) {
	testFun(t)

	forever := make(chan bool)
	log.Printf(" [*] To exit press CTRL+C")
	<-forever
}

func TestQuery(t *testing.T) {
	dao, err := dao.Fetch("magiccenter", "magiccenter", "localhost:3306", "magiccenter_db")
	if err != nil {
		t.Errorf("Fetch dao failed, err:%s", err.Error())
	}
	defer dao.Release()

	selectSql := fmt.Sprint("select id,name,password,type from magiccenter_db.user")

	dao.Query(selectSql)

	for dao.Next() {
		user := User{}

		dao.GetField(&user.id, &user.name, &user.password, &user.catalog)
	}
}
